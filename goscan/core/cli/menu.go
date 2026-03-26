package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/marco-lancini/goscan/core/model"
	"github.com/marco-lancini/goscan/core/utils"
)

// Color definitions for better UI
var (
	colorRed     = color.New(color.FgRed).SprintFunc()
	colorGreen   = color.New(color.FgGreen).SprintFunc()
	colorYellow  = color.New(color.FgYellow).SprintFunc()
	colorBlue    = color.New(color.FgBlue).SprintFunc()
	colorCyan    = color.New(color.FgCyan).SprintFunc()
	colorMagenta = color.New(color.FgMagenta).SprintFunc()
	colorWhite   = color.New(color.FgWhite).SprintFunc()

	boldGreen  = color.New(color.FgGreen, color.Bold).SprintFunc()
	boldRed    = color.New(color.FgRed, color.Bold).SprintFunc()
	boldCyan   = color.New(color.FgCyan, color.Bold).SprintFunc()
	boldYellow = color.New(color.FgYellow, color.Bold).SprintFunc()
)

// Menu structure
type MenuItem struct {
	ID          int
	Title       string
	Description string
	Handler     func()
}

// ScanType represents different scanning categories
type ScanType struct {
	ID          int
	Name        string
	Description string
	Handler     func()
}

// Full catalog of scans (used for See More)
var fullMenuItems = []MenuItem{
	{1, "Basic Scans", "Fundamental network scanning operations", handleBasicScans},
	{2, "Security Scans", "Advanced security-focused scanning", handleSecurityScans},
	{3, "Network Discovery", "Discover devices and hosts on network", handleNetworkDiscovery},
	{4, "TCP Scans", "Detailed TCP port and service scanning", handleTCPScans},
	{5, "ARP Scans", "ARP-based device discovery", handleARPScans},
	{6, "Authentication Scans", "Test authentication services", handleAuthScans},
	{7, "Web Application Scans", "Web server and application scanning", handleWebAppScans},
	{8, "Mobile Device Scans", "Scan mobile and IoT devices", handleMobileScans},
	{9, "Gaming Console Scans", "Gaming device discovery and scanning", handleGamingScans},
	{10, "IoT Device Scans", "Internet of Things device scanning", handleIoTScans},
	{11, "CCTV & Drone Scans", "Surveillance and drone device scanning", handleCCTVScans},
	{12, "MAC Address Analysis", "MAC address and vendor identification", handleMACAnalysis},
	{13, "View Scan Logs", "View previous scan results and logs", handleVirusScanLogs},
	{14, "View Scan Cache", "Access cached scan results", handleScanCache},
}

// Primary menu shows top scans plus utilities and navigation
var primaryMenuItems = []MenuItem{
	{1, "Basic Scans", "Fundamental network scanning operations", handleBasicScans},
	{2, "Security Scans", "Advanced security-focused scanning", handleSecurityScans},
	{3, "Network Discovery", "Discover devices and hosts on network", handleNetworkDiscovery},
	{4, "TCP Scans", "Detailed TCP port and service scanning", handleTCPScans},
	{5, "ARP Scans", "ARP-based device discovery", handleARPScans},
	{6, "Authentication Scans", "Test authentication services", handleAuthScans},
	{7, "Web Application Scans", "Web server and application scanning", handleWebAppScans},
	{8, "See More", "View all scan categories", handleSeeMore},
	{9, "View Scan Logs", "View previous scan results and logs", handleVirusScanLogs},
	{10, "View Scan Cache", "Access cached scan results", handleScanCache},
	{11, "Exit", "Exit the program", func() { os.Exit(0) }},
}

// PrintBanner displays the formatted banner with animation
func PrintBanner() {
	// Animated banner border
	fmt.Print(boldGreen("╔"))
	for i := 0; i < 70; i++ {
		fmt.Print(boldGreen("═"))
		time.Sleep(2 * time.Millisecond)
	}
	fmt.Println(boldGreen("╗"))

	// Title with animation
	fmt.Print(boldGreen("║"))
	fmt.Print(colorCyan("                              "))
	fmt.Print(boldGreen("ADVANCESCANNER"))
	fmt.Println(colorCyan("                           " + boldGreen("║")))

	// Subtitle
	fmt.Print(boldGreen("╚"))
	for i := 0; i < 70; i++ {
		fmt.Print(boldGreen("═"))
		time.Sleep(2 * time.Millisecond)
	}
	fmt.Println(boldGreen("╝"))

	// Animated subtitle text
	fmt.Println()
	utils.ScrollText("⚡ Cross-Platform Network Toolkit ⚡")
	fmt.Println(colorYellow("⚠  WARNING: Unauthorized scanning may be illegal. Use responsibly and with proper authorization."))
	fmt.Println()

	// Creator attribution
	fmt.Println(colorCyan("Created by: ") + colorMagenta("Suvaditya Roy"))
	fmt.Println()
}

// PrintMainMenu displays the main menu with animation
func PrintMainMenu() {
	fmt.Print("\n")

	// Animated header
	fmt.Print(boldCyan("╔"))
	for i := 0; i < 66; i++ {
		fmt.Print(boldCyan("═"))
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(boldCyan("╗"))

	fmt.Println(boldCyan("║") + colorWhite("                       MAIN MENU                                  ") + boldCyan("║"))

	fmt.Print(boldCyan("╚"))
	for i := 0; i < 66; i++ {
		fmt.Print(boldCyan("═"))
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(boldCyan("╝"))

	fmt.Print("\n")

	// Animated menu items
	for i, item := range primaryMenuItems {
		colorDot := colorGreen("●")
		if item.Title == "See More" {
			colorDot = colorBlue("●")
		} else if item.Title == "Exit" {
			colorDot = colorRed("●")
		}

		// Fade in effect for each menu item
		fmt.Printf("%s [%2d] %s - %s\n", colorDot, item.ID, boldGreen(item.Title), colorWhite(item.Description))

		// Add slight delay for smooth animation
		if i < len(primaryMenuItems)-1 {
			time.Sleep(30 * time.Millisecond)
		}
	}
	fmt.Print("\n")
}

// PrintFullMenu displays all scan categories with animation
func PrintFullMenu() {
	fmt.Print("\n")

	// Animated header
	fmt.Print(boldCyan("╔"))
	for i := 0; i < 66; i++ {
		fmt.Print(boldCyan("═"))
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(boldCyan("╗"))

	fmt.Println(boldCyan("║") + colorWhite("                     ALL SCAN CATEGORIES                           ") + boldCyan("║"))

	fmt.Print(boldCyan("╚"))
	for i := 0; i < 66; i++ {
		fmt.Print(boldCyan("═"))
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(boldCyan("╝"))

	fmt.Print("\n")

	// Animated menu items
	for idx, item := range fullMenuItems {
		fmt.Printf("%s [%2d] %s - %s\n", colorGreen("●"), idx+1, boldGreen(item.Title), colorWhite(item.Description))
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Printf("%s [%2d] %s - %s\n", colorRed("●"), 0, boldRed("Back"), colorRed("Return to main menu"))
	fmt.Print("\n")
}

// GetUserChoice gets menu choice from user
func GetUserChoice(maxChoice int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s Enter your choice (0-%d): ", colorCyan("▶"), maxChoice)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 0 || choice > maxChoice {
			fmt.Println(colorRed("✗ Invalid choice. Please try again."))
			continue
		}
		return choice
	}
}

// ShowMenu displays and handles menu interaction
func ShowMenu() {
	for {
		PrintMainMenu()
		choice := GetUserChoice(len(primaryMenuItems))

		if choice == 0 {
			fmt.Println(colorYellow("Please select a valid option."))
			continue
		}

		executeMenuSelection(primaryMenuItems, choice)
	}
}

// executeMenuSelection finds item by ID and runs handler
func executeMenuSelection(items []MenuItem, choice int) {
	for _, item := range items {
		if item.ID == choice {
			fmt.Print("\n")
			fmt.Println(boldYellow("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
			item.Handler()
			fmt.Println(boldYellow("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
			fmt.Print("\n")
			return
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════════════
// HANDLER FUNCTIONS
// ═══════════════════════════════════════════════════════════════════════════════════════

func handleBasicScans() {
	fmt.Println(boldCyan("\n[1] BASIC SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ Network Ping Sweep") + " - Discover live hosts")
	fmt.Println(colorGreen("▸ Host Discovery") + " - Identify devices on network")
	fmt.Println(colorGreen("▸ Quick Port Scan") + " - Scan top 1000 ports")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter target IP or range: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting ping sweep on %s\n", colorGreen("✓"), colorYellow(target))
		// Execute sweep command
		ExecuteSweep(target)
	}
}

func handleSecurityScans() {
	fmt.Println(boldCyan("\n[2] SECURITY SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ Vulnerability Scan") + " - Scan for common vulnerabilities")
	fmt.Println(colorGreen("▸ SSL/TLS Analysis") + " - Analyze SSL/TLS certificates and protocols")
	fmt.Println(colorGreen("▸ Service Version Detection") + " - Identify service versions")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter target IP: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting security scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteSecurityScan(target)
	}
}

func handleNetworkDiscovery() {
	fmt.Println(boldCyan("\n[3] NETWORK DISCOVERY"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ ARP Discovery") + " - Find devices using ARP")
	fmt.Println(colorGreen("▸ DNS Enumeration") + " - Enumerate DNS records")
	fmt.Println(colorGreen("▸ Network Mapping") + " - Create network topology map")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter network range (e.g., 192.168.1.0/24): ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting network discovery on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteNetworkDiscovery(target)
	}
}

func handleTCPScans() {
	fmt.Println(boldCyan("\n[4] TCP SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ TCP Connect Scan") + " - Full TCP 3-way handshake")
	fmt.Println(colorGreen("▸ TCP SYN Scan") + " - Stealth SYN scan (requires root)")
	fmt.Println(colorGreen("▸ TCP Full Port Scan") + " - Scan all 65535 ports")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter target IP: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting TCP scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteTCPScan(target)
	}
}

func handleARPScans() {
	fmt.Println(boldCyan("\n[5] ARP SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ ARP Sweep") + " - Discover live hosts using ARP")
	fmt.Println(colorGreen("▸ MAC Vendor Lookup") + " - Identify device manufacturers")
	fmt.Println(colorGreen("▸ ARP Spoofing Detection") + " - Detect ARP spoofing attempts")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter network range: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting ARP scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteARPScan(target)
	}
}

func handleAuthScans() {
	fmt.Println(boldCyan("\n[6] AUTHENTICATION SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ SSH Brute Force") + " - Test SSH credentials")
	fmt.Println(colorGreen("▸ FTP Enumeration") + " - Enumerate FTP services")
	fmt.Println(colorGreen("▸ HTTP Basic Auth Test") + " - Test HTTP authentication")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter target IP: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting authentication scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteAuthScan(target)
	}
}

func handleWebAppScans() {
	fmt.Println(boldCyan("\n[7] WEB APPLICATION SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ Web Server Detection") + " - Identify web servers")
	fmt.Println(colorGreen("▸ Web Vulnerability Scan") + " - Scan for web vulnerabilities")
	fmt.Println(colorGreen("▸ Screenshot Capture (EyeWitness)") + " - Visual web service analysis")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter target URL or IP: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting web app scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteWebAppScan(target)
	}
}

func handleSeeMore() {
	for {
		PrintFullMenu()
		choice := GetUserChoice(len(fullMenuItems))
		if choice == 0 {
			return
		}
		if choice > 0 && choice <= len(fullMenuItems) {
			// In See More we index by position (choice-1)
			fmt.Print("\n")
			fmt.Println(boldYellow("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
			fullMenuItems[choice-1].Handler()
			fmt.Println(boldYellow("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
			fmt.Print("\n")
		}
	}
}

func handleMobileScans() {
	fmt.Println(boldCyan("\n[8] MOBILE DEVICE SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ Android Device Discovery") + " - Find Android devices")
	fmt.Println(colorGreen("▸ iOS Device Detection") + " - Detect iOS devices")
	fmt.Println(colorGreen("▸ Mobile App Scanning") + " - Scan mobile apps for vulnerabilities")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter network range: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting mobile device scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteMobileScan(target)
	}
}

func handleGamingScans() {
	fmt.Println(boldCyan("\n[9] GAMING CONSOLE SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ PlayStation Detection") + " - Find PS4/PS5 consoles")
	fmt.Println(colorGreen("▸ Xbox Detection") + " - Find Xbox consoles")
	fmt.Println(colorGreen("▸ Nintendo Detection") + " - Find Nintendo devices")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter network range: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting gaming console scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteGamingScan(target)
	}
}

func handleIoTScans() {
	fmt.Println(boldCyan("\n[10] IOT DEVICE SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ Smart Home Device Detection") + " - Find IoT devices")
	fmt.Println(colorGreen("▸ Raspberry Pi Discovery") + " - Identify Raspberry Pi devices")
	fmt.Println(colorGreen("▸ MQTT Service Scan") + " - Detect MQTT brokers")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter network range: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting IoT device scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteIoTScan(target)
	}
}

func handleCCTVScans() {
	fmt.Println(boldCyan("\n[11] CCTV & DRONE SCANS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ CCTV Camera Detection") + " - Find surveillance cameras")
	fmt.Println(colorGreen("▸ Drone Discovery") + " - Identify drone devices")
	fmt.Println(colorGreen("▸ Video Stream Detection") + " - Locate video streams")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter network range: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting CCTV/Drone scan on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteCCTVScan(target)
	}
}

func handleMACAnalysis() {
	fmt.Println(boldCyan("\n[12] MAC ADDRESS ANALYSIS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ Vendor Lookup") + " - Identify MAC address vendor")
	fmt.Println(colorGreen("▸ Local Network Analysis") + " - Analyze all MAC addresses on network")
	fmt.Println(colorGreen("▸ Device Fingerprinting") + " - Fingerprint devices by MAC")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s Enter network range or MAC address: ", colorCyan("▶"))
	target, _ := reader.ReadString('\n')
	target = strings.TrimSpace(target)

	if target != "" {
		fmt.Printf("\n%s Starting MAC analysis on %s\n", colorGreen("✓"), colorYellow(target))
		ExecuteMACAnalysis(target)
	}
}

func handleVirusScanLogs() {
	fmt.Println(boldCyan("\n[13] VIRUS SCAN LOGS"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ View Scan History") + " - Display previous scans")
	fmt.Println(colorGreen("▸ Threat Analysis") + " - Analyze detected threats")
	fmt.Println(colorGreen("▸ Export Results") + " - Save results to file")

	fmt.Printf("\n%s Retrieving scan logs...\n", colorGreen("✓"))
	DisplayScanLogs()
}

func handleScanCache() {
	fmt.Println(boldCyan("\n[14] VIEW SCAN CACHE"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────\n"))
	fmt.Println(colorGreen("▸ Database Cache") + " - View cached results")
	fmt.Println(colorGreen("▸ Import Cache") + " - Import cached data")
	fmt.Println(colorGreen("▸ Clear Cache") + " - Delete cached results")

	fmt.Printf("\n%s Accessing scan cache...\n", colorGreen("✓"))
	DisplayScanCache()
}

// ═══════════════════════════════════════════════════════════════════════════════════════
// EXECUTION FUNCTIONS
// ═══════════════════════════════════════════════════════════════════════════════════════

func ExecuteSweep(target string) {
	fmt.Printf("%s [SWEEP] Performing ping sweep on %s...\n", colorGreen("►"), colorYellow(target))
	// Ensure target is tracked when a database is available
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	// Call the actual nmap sweep function using the expected keyword
	cmdSweep([]string{"PING", target})
}

func ExecuteSecurityScan(target string) {
	fmt.Printf("%s [SECURITY] Performing security scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddHost(utils.Config.DB, target, "up", model.NEW.String())
	}
	cmdPortscan([]string{"TCP-VULN-SCAN", target})
}

func ExecuteNetworkDiscovery(target string) {
	fmt.Printf("%s [DISCOVERY] Discovering network topology on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	cmdSweep([]string{"PING", target})
}

func ExecuteTCPScan(target string) {
	fmt.Printf("%s [TCP] Performing TCP scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddHost(utils.Config.DB, target, "up", model.NEW.String())
	}
	cmdPortscan([]string{"TCP-FULL", target})
}

func ExecuteARPScan(target string) {
	fmt.Printf("%s [ARP] Performing ARP scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	cmdSweep([]string{"PING", target})
}

func ExecuteAuthScan(target string) {
	fmt.Printf("%s [AUTH] Performing authentication enumeration on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddHost(utils.Config.DB, target, "up", model.NEW.String())
	}
	cmdEnumerate([]string{"SSH", "POLITE", target})
}

func ExecuteWebAppScan(target string) {
	fmt.Printf("%s [WEB] Performing web application scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddHost(utils.Config.DB, target, "up", model.NEW.String())
	}
	cmdEnumerate([]string{"HTTP", "POLITE", target})
}

func ExecuteMobileScan(target string) {
	fmt.Printf("%s [MOBILE] Performing mobile device scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	cmdSweep([]string{"PING", target})
}

func ExecuteGamingScan(target string) {
	fmt.Printf("%s [GAMING] Performing gaming console scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	cmdSweep([]string{"PING", target})
}

func ExecuteIoTScan(target string) {
	fmt.Printf("%s [IOT] Performing IoT device scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	cmdSweep([]string{"PING", target})
}

func ExecuteCCTVScan(target string) {
	fmt.Printf("%s [CCTV] Performing CCTV/Drone scan on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	cmdSweep([]string{"PING", target})
}

func ExecuteMACAnalysis(target string) {
	fmt.Printf("%s [MAC] Performing MAC address analysis on %s...\n", colorGreen("►"), colorYellow(target))
	if utils.IsDBAvailable() {
		model.AddTarget(utils.Config.DB, target, model.IMPORTED.String())
	}
	cmdSweep([]string{"PING", target})
}

func DisplayScanLogs() {
	fmt.Println("\n" + boldCyan("Recent Scan History:"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))
	// Retrieve from database
	cmdShow([]string{"targets"})
}

func DisplayScanCache() {
	fmt.Println("\n" + boldCyan("Scan Cache:"))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))
	// Retrieve from database
	cmdShow([]string{"ports"})
}
