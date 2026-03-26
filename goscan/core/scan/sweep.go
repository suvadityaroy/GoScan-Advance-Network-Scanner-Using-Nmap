package scan

import (
	"fmt"

	"github.com/marco-lancini/goscan/core/model"
	"github.com/marco-lancini/goscan/core/utils"
)

// ---------------------------------------------------------------------------------------
// DISPATCHER
// ---------------------------------------------------------------------------------------
func ScanSweep(kind string, target string) {
	// Dispatch scan
	switch kind {
	case "PING":
		utils.Config.Log.LogInfo("Starting Ping Sweep")
		folder, file, nmapArgs := "sweep", "ping", utils.Const_NMAP_SWEEP
		execSweep(file, target, folder, file, nmapArgs)

	default:
		utils.Config.Log.LogError("Invalid type of scan")
		return
	}
}

func execSweep(name, target, folder, file, nmapArgs string) {
	// If no database is available (Windows build), run directly against the provided target
	if !utils.IsDBAvailable() {
		temp := model.Target{Address: target, Step: model.IMPORTED.String()}
		fname := fmt.Sprintf("%s_%s", file, target)
		go workerSweep(name, &temp, folder, fname, nmapArgs)
		return
	}

	targets := model.GetAllTargets(utils.Config.DB)
	for _, h := range targets {
		// Scan only if:
		//   - target is ALL
		//   - or if target is TO_ANALYZE and host still need to be analyzed
		//   - or if host is the selected one
		if target == "ALL" ||
			(target == "TO_ANALYZE" && h.Step == model.IMPORTED.String()) ||
			target == h.Address {
			temp := h
			fname := fmt.Sprintf("%s_%s", file, h.Address)
			go workerSweep(name, &temp, folder, fname, nmapArgs)
		}
	}
}

// ---------------------------------------------------------------------------------------
// WORKER
// ---------------------------------------------------------------------------------------
func workerSweep(name string, h *model.Target, folder string, file string, nmapArgs string) {
	// Instantiate new NmapScan
	s := NewScan(name, h.Address, folder, file, nmapArgs)
	ScansList = append(ScansList, s)

	// Run the scan
	s.RunNmap()

	// Parse nmap's output
	res := s.ParseOutput()
	if res != nil {
		// Identify live hosts
		for _, host := range res.Hosts {
			status := host.Status.State
			if status == "up" {
				addr := host.Addresses[0].Addr
				if utils.IsDBAvailable() {
					model.AddHost(utils.Config.DB, addr, "up", model.NEW.String())
				} else {
					utils.Config.Log.LogInfo(fmt.Sprintf("Host discovered: %s (not persisted - DB disabled)", addr))
				}
			}
		}
	}

	// Update status of target
	if utils.IsDBAvailable() {
		model.Mutex.Lock()
		h.Step = model.SWEEPED.String()
		utils.Config.DB.Save(&h)
		model.Mutex.Unlock()
	}
}
