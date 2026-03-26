package utils

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Animation effects and utilities
var (
	colorGreen   = color.New(color.FgGreen).SprintFunc()
	colorCyan    = color.New(color.FgCyan).SprintFunc()
	colorYellow  = color.New(color.FgYellow).SprintFunc()
	colorBlue    = color.New(color.FgBlue).SprintFunc()
	colorMagenta = color.New(color.FgMagenta).SprintFunc()
	boldGreen    = color.New(color.FgGreen, color.Bold).SprintFunc()
	boldCyan     = color.New(color.FgCyan, color.Bold).SprintFunc()
)

// LoadingSpinner shows an animated loading spinner
func LoadingSpinner(message string, duration time.Duration) {
	spinnerFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range spinnerFrames {
			if time.Now().After(endTime) {
				break
			}
			fmt.Printf("\r%s %s", colorCyan(frame), message)
			time.Sleep(80 * time.Millisecond)
		}
	}
	fmt.Printf("\r%s %s\n", colorGreen("✓"), message)
}

// TypewriterEffect prints text with a typewriter effect
func TypewriterEffect(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

// ProgressBar shows an animated progress bar
func ProgressBar(label string, current int, total int, width int) {
	if total == 0 {
		return
	}

	percent := float64(current) / float64(total)
	filled := int(float64(width) * percent)

	bar := "["
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += "]"

	fmt.Printf("\r%s %s %3d%%", label, colorGreen(bar), int(percent*100))
	if current == total {
		fmt.Println()
	}
}

// ScanStartAnimation shows animation when scan starts
func ScanStartAnimation(scanType string, target string) {
	fmt.Println()
	fmt.Println(colorCyan("╔════════════════════════════════════════════════════════════════╗"))

	// Animated scan type
	fmt.Printf(colorCyan("║ "))
	TypewriterEffect("Starting "+scanType+" on "+target, 15*time.Millisecond)

	frames := []string{"⠏", "⠛", "⠖", "⠒", "⠐", "⠠", "⠄"}
	for i := 0; i < 3; i++ {
		for _, frame := range frames {
			if i == 2 && frames[len(frames)-1] == frame {
				fmt.Printf("%s\n", colorCyan("║ ✓ Scan initializing..."))
			} else {
				fmt.Printf("\r%s", colorCyan("║ "+frame+" Scan initializing..."))
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	fmt.Println(colorCyan("╚════════════════════════════════════════════════════════════════╝"))
	fmt.Println()
}

// ScanCompleteAnimation shows animation when scan completes
func ScanCompleteAnimation(scanType string, target string, resultsCount int) {
	fmt.Println()
	fmt.Println(colorGreen("╔════════════════════════════════════════════════════════════════╗"))
	fmt.Println(colorGreen("║                    ✓ SCAN COMPLETED SUCCESSFULLY ✓             ║"))
	fmt.Println(colorGreen("╚════════════════════════════════════════════════════════════════╝"))
	fmt.Println()

	// Confetti-like animation
	confetti := []string{"✨", "🎉", "⭐", "💫", "🌟"}
	for i := 0; i < 3; i++ {
		fmt.Printf("%s ", colorYellow(confetti[i%len(confetti)]))
	}
	fmt.Println()

	// Scan summary
	fmt.Printf("%s Scan Type: %s\n", colorGreen("►"), scanType)
	fmt.Printf("%s Target: %s\n", colorGreen("►"), target)
	fmt.Printf("%s Results Found: %s\n", colorGreen("►"), colorGreen(fmt.Sprintf("%d", resultsCount)))

	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Printf("%s ", colorYellow(confetti[i%len(confetti)]))
	}
	fmt.Println()
	fmt.Println()
}

// ScanFailedAnimation shows animation when scan fails
func ScanFailedAnimation(scanType string, target string, errorMsg string) {
	fmt.Println()
	fmt.Println(color.New(color.FgRed).Sprint("╔════════════════════════════════════════════════════════════════╗"))
	fmt.Println(color.New(color.FgRed).Sprint("║                    ✗ SCAN FAILED ✗                           ║"))
	fmt.Println(color.New(color.FgRed).Sprint("╚════════════════════════════════════════════════════════════════╝"))
	fmt.Println()

	fmt.Printf("%s Scan Type: %s\n", color.New(color.FgRed).Sprint("►"), scanType)
	fmt.Printf("%s Target: %s\n", color.New(color.FgRed).Sprint("►"), target)
	fmt.Printf("%s Error: %s\n", color.New(color.FgRed).Sprint("►"), errorMsg)
	fmt.Println()
}

// MenuTransition shows smooth transition between menus
func MenuTransition(fromTitle string, toTitle string) {
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Printf("\r%s", colorCyan("|"))
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\r%s", colorCyan("/"))
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\r%s", colorCyan("-"))
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\r%s", colorCyan("\\"))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r")
	fmt.Println()
}

// PulseEffect creates a pulsing animation effect
func PulseEffect(message string, pulses int) {
	for i := 0; i < pulses; i++ {
		fmt.Printf("\r%s", colorMagenta(message))
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("\r%s", message)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("\r%s\n", colorMagenta(message))
}

// WaveEffect shows a wave animation
func WaveEffect() {
	waves := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"}
	for i := 0; i < 3; i++ {
		for _, wave := range waves {
			fmt.Printf("\r%s", colorCyan(wave))
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Printf("\r")
	fmt.Println()
}

// FadeInText displays text with fade-in effect
func FadeInText(text string, lines []string) {
	fmt.Println(colorCyan(text))
	for _, line := range lines {
		fmt.Println(colorYellow("  " + line))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

// ScrollText displays text that scrolls in from left to right
func ScrollText(text string) {
	padding := ""
	for i := 0; i < len(text); i++ {
		padding += text[i : i+1]
		fmt.Printf("\r%s", colorCyan(padding))
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Printf("\r%s\n", colorCyan(text))
}

// AnimatedBox displays a box with animation
func AnimatedBox(title string, content []string) {
	boxWidth := 60

	// Top border
	fmt.Print(colorCyan("╔"))
	for i := 0; i < boxWidth-2; i++ {
		fmt.Print(colorCyan("═"))
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println(colorCyan("╗"))

	// Title
	padding := (boxWidth - 4 - len(title)) / 2
	fmt.Print(colorCyan("║ "))
	fmt.Print(colorYellow(title))
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(colorCyan(" ║"))

	// Divider
	fmt.Print(colorCyan("╠"))
	for i := 0; i < boxWidth-2; i++ {
		fmt.Print(colorCyan("═"))
	}
	fmt.Println(colorCyan("╣"))

	// Content
	for _, line := range content {
		fmt.Print(colorCyan("║ "))
		fmt.Print(line)
		fmt.Println(colorCyan(" ║"))
		time.Sleep(100 * time.Millisecond)
	}

	// Bottom border
	fmt.Print(colorCyan("╚"))
	for i := 0; i < boxWidth-2; i++ {
		fmt.Print(colorCyan("═"))
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println(colorCyan("╝"))
	fmt.Println()
}

// CountdownTimer shows an animated countdown
func CountdownTimer(seconds int) {
	for i := seconds; i > 0; i-- {
		fmt.Printf("\r%s %d seconds...", colorYellow("⏱"), i)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\r%s\n", colorGreen("✓ Ready!"))
}

// BlinkText makes text blink
func BlinkText(text string, times int) {
	for i := 0; i < times*2; i++ {
		if i%2 == 0 {
			fmt.Printf("\r%s", colorMagenta(text))
		} else {
			fmt.Printf("\r%s", "                                                                                  ")
		}
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Printf("\r%s\n", colorMagenta(text))
}
