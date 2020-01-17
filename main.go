package main

import (
	"github.com/davidv171/release-subscriber/cmd"
	"github.com/davidv171/release-subscriber/flags"
	"os"
)

/* Read CLI arguments */
/* Exmaple:*/
/* ./release-subscriber subscribe github.com/bruh */

func main() {

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		cmd.Help()
		os.Exit(0)
	}

	//Subcommand download
	if os.Args[1] == "download" || os.Args[1] == "d" {

		flags.Parse()

	}

}
