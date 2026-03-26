package scan

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var (
	colorGreen  = color.New(color.FgGreen).SprintFunc()
	colorCyan   = color.New(color.FgCyan).SprintFunc()
	colorYellow = color.New(color.FgYellow).SprintFunc()
	boldGreen   = color.New(color.FgGreen, color.Bold).SprintFunc()
	boldCyan    = color.New(color.FgCyan, color.Bold).SprintFunc()
)

// SecurityScan performs vulnerability and security assessment
func SecurityScan(target string) {
	fmt.Printf("%s Starting Security Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Vulnerability scan with NSE scripts
	cmd := exec.Command("nmap", "-sV", "--script", "vuln", "-p", "1-10000", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during security scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s Security scan completed successfully\n", colorGreen("✓"))
	}
}

// MobileScan scans for mobile and IoT devices
func MobileScan(target string) {
	fmt.Printf("%s Starting Mobile Device Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Scan for common mobile device ports
	commonMobilePorts := []string{
		"5555",  // ADB
		"8080",  // HTTP alt
		"8443",  // HTTPS alt
		"9100",  // Print service
		"10000", // Webmin
		"5900",  // VNC
	}

	cmd := exec.Command("nmap", "-p", strings.Join(commonMobilePorts, ","), "-sV", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during mobile scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s Mobile device scan completed\n", colorGreen("✓"))
	}
}

// IoTScan scans for IoT devices and services
func IoTScan(target string) {
	fmt.Printf("%s Starting IoT Device Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Common IoT ports
	iotPorts := []string{
		"1883",  // MQTT
		"8883",  // MQTT SSL
		"5353",  // mDNS
		"5900",  // VNC
		"8080",  // HTTP
		"8443",  // HTTPS
		"9100",  // Print
		"9200",  // Elasticsearch
		"27017", // MongoDB
		"5432",  // PostgreSQL
	}

	cmd := exec.Command("nmap", "-p", strings.Join(iotPorts, ","), "-sV", "-O", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during IoT scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s IoT device scan completed\n", colorGreen("✓"))
	}
}

// WebAppScan scans web applications for vulnerabilities
func WebAppScan(target string) {
	fmt.Printf("%s Starting Web Application Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Web-specific ports and services
	webPorts := []string{
		"80",   // HTTP
		"8080", // HTTP alt
		"8000", // Python http
		"443",  // HTTPS
		"8443", // HTTPS alt
		"3000", // Node.js
		"5000", // Flask
		"9000", // Elasticsearch
	}

	cmd := exec.Command("nmap", "-p", strings.Join(webPorts, ","), "-sV", "--script", "http-*", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during web app scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s Web application scan completed\n", colorGreen("✓"))
	}
}

// GamingConsoleScan scans for gaming devices
func GamingConsoleScan(target string) {
	fmt.Printf("%s Starting Gaming Console Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Gaming console ports
	gamingPorts := []string{
		"3074", // Xbox Live
		"5223", // PlayStation Network
		"8080", // Web services
		"443",  // HTTPS
		"80",   // HTTP
		"53",   // DNS
	}

	cmd := exec.Command("nmap", "-p", strings.Join(gamingPorts, ","), "-sV", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during gaming console scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s Gaming console scan completed\n", colorGreen("✓"))
	}
}

// CCTVScan scans for surveillance and drone devices
func CCTVScan(target string) {
	fmt.Printf("%s Starting CCTV/Drone Device Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// CCTV and drone ports
	cctvPorts := []string{
		"80",   // HTTP
		"8080", // HTTP alt
		"443",  // HTTPS
		"8443", // HTTPS alt
		"554",  // RTSP (streaming)
		"5000", // UPnP
		"5900", // VNC
		"9000", // Axis cameras
	}

	cmd := exec.Command("nmap", "-p", strings.Join(cctvPorts, ","), "-sV", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during CCTV scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s CCTV/Drone scan completed\n", colorGreen("✓"))
	}
}

// MACAddressAnalysis performs MAC address lookup and analysis
func MACAddressAnalysis(target string) {
	fmt.Printf("%s Starting MAC Address Analysis on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Perform ARP scan to get MAC addresses
	cmd := exec.Command("nmap", "-sn", "-PE", "-PA", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during MAC analysis: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s MAC address analysis completed\n", colorGreen("✓"))
	}
}

// ARPScan performs ARP-based network discovery
func ARPScan(target string) {
	fmt.Printf("%s Starting ARP Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// ARP scan with OS detection
	cmd := exec.Command("nmap", "-PR", "-O", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during ARP scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s ARP scan completed\n", colorGreen("✓"))
	}
}

// NetworkDiscoveryScan performs comprehensive network discovery
func NetworkDiscoveryScan(target string) {
	fmt.Printf("%s Starting Network Discovery Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Comprehensive network discovery
	cmd := exec.Command("nmap", "-sn", "-PE", "-PP", "-PM", "-PU", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during network discovery: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s Network discovery completed\n", colorGreen("✓"))
	}
}

// AuthenticationScan tests authentication services
func AuthenticationScan(target string) {
	fmt.Printf("%s Starting Authentication Scan on %s\n", boldCyan("►"), colorYellow(target))
	fmt.Println(colorCyan("─────────────────────────────────────────────────────────────────"))

	// Scan for authentication services
	authPorts := []string{
		"22",   // SSH
		"21",   // FTP
		"3389", // RDP
		"445",  // SMB
		"389",  // LDAP
		"636",  // LDAP SSL
		"5432", // PostgreSQL
		"3306", // MySQL
		"1433", // MSSQL
	}

	cmd := exec.Command("nmap", "-p", strings.Join(authPorts, ","), "-sV", "--script", "auth-*", target)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s Error during authentication scan: %v\n", color.RedString("✗"), err)
	} else {
		fmt.Printf("%s Authentication scan completed\n", colorGreen("✓"))
	}
}
