package main

import (
	"github.com/marco-lancini/goscan/core/cli"
	"github.com/marco-lancini/goscan/core/utils"
)

var (
	author  string
	version string
)

// ---------------------------------------------------------------------------------------
// INIT
// ---------------------------------------------------------------------------------------
func initCore() {
	// Check sudo
	utils.CheckSudo()
	// Initialize global config (db, logger, etc.)
	// From now on it will be accessible as utils.Config
	utils.InitConfig()
}

// ---------------------------------------------------------------------------------------
// MAIN
// ---------------------------------------------------------------------------------------
func main() {
	// Setup core
	initCore()

	// Show banner
	cli.PrintBanner()

	// Show interactive menu system in loop
	for {
		cli.ShowMenu()
	}
}
