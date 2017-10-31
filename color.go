package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {

	// pull in cli values
	Color := flag.String("color", "", "Color to highlight")
	Background := flag.Bool("background", false, "Background highlighting [false]")
	Pattern := flag.String("pattern", "", "Pattern to highlight")
	flag.Parse()

	// check color
	highlight := color.FgWhite
	if *Background {
		switch strings.ToLower(*Color) {
		case "red":
			highlight = color.BgRed
		case "green":
			highlight = color.BgGreen
		case "yellow":
			highlight = color.BgYellow
		case "blue":
			highlight = color.BgBlue
		case "magenta":
			highlight = color.BgMagenta
		case "cyan":
			highlight = color.BgMagenta
		}
	} else {
		switch strings.ToLower(*Color) {
		case "red":
			highlight = color.FgRed
		case "green":
			highlight = color.FgGreen
		case "yellow":
			highlight = color.FgYellow
		case "blue":
			highlight = color.FgBlue
		case "magenta":
			highlight = color.FgMagenta
		case "cyan":
			highlight = color.FgMagenta
		}
	}
	if highlight == color.FgWhite {
		fmt.Fprintln(os.Stderr, "ERROR: Not a valid color")
	}
	reColor := color.New(highlight).SprintFunc()

	// check for pattern
	pattern := *Pattern
	if pattern == "" {
		fmt.Fprintln(os.Stderr, "ERROR: No provided pattern")
	}

	// setup standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(strings.Replace(scanner.Text(), pattern, reColor(pattern), -1))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: ", err)
	}
}
