package main

import (
	"fmt"
	"code.google.com/p/goncurses"
)

func main() {

	// You should always test to make sure ncurses has initialized properly.
	// In order for your error messages to be visible on the terminal you will
	// need to either log error messages or output them to to stderr.

	stdscr, err := goncurses.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer goncurses.End()
	stdscr.Print("Press enter to continue...")
	stdscr.Refresh()
}
