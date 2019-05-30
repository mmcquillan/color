package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func main() {

	// pull in cli values
	Color := flag.String("color", "", "Color to highlight")
	Background := flag.Bool("background", false, "Background highlighting [false]")
	Pattern := flag.String("pattern", "", "Pattern to highlight")
	Strip := flag.Bool("strip", false, "Strip all colors [false]")
	flag.Parse()
	strip := *Strip

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
			highlight = color.BgCyan
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
			highlight = color.FgCyan
		}
	}
	if highlight == color.FgWhite && !strip {
		fmt.Fprintln(os.Stderr, "ERROR: Not a valid color")
	}
	reColor := color.New(highlight).SprintFunc()

	// check for pattern
	pattern := *Pattern
	if pattern == "" && !strip {
		fmt.Fprintln(os.Stderr, "ERROR: No provided pattern")
	}

	// strip patterns
	stripreg := regexp.MustCompile("(\\033|\\027)\\[[0-9]*m")

	// setup standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strip {
			fmt.Println(stripreg.ReplaceAllString(scanner.Text(), ""))
		} else {
			fmt.Println(strings.Replace(scanner.Text(), pattern, reColor(pattern), -1))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: ", err)
	}
}
