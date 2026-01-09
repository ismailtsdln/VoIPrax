package ui

import (
	"fmt"

	"github.com/fatih/color"
)

const Banner = `
V  o  I  P  r  a  x
VoIP Penetration & Analysis eXtreme
`

// PrintBanner prints the stylized ASCII banner
func PrintBanner() {
	c := color.New(color.FgCyan, color.Bold)
	c.Println(Banner)
	fmt.Println(color.HiBlackString("--------------------------------------------------"))
}

// Info prints an informational message
func Info(msg string, args ...interface{}) {
	fmt.Printf("%s %s\n", color.BlueString("[*]"), fmt.Sprintf(msg, args...))
}

// Success prints a success message
func Success(msg string, args ...interface{}) {
	fmt.Printf("%s %s\n", color.GreenString("[+]"), fmt.Sprintf(msg, args...))
}

// Error prints an error message
func Error(msg string, args ...interface{}) {
	fmt.Printf("%s %s\n", color.RedString("[!]"), fmt.Sprintf(msg, args...))
}

// Warning prints a warning message
func Warning(msg string, args ...interface{}) {
	fmt.Printf("%s %s\n", color.YellowString("[?]"), fmt.Sprintf(msg, args...))
}
