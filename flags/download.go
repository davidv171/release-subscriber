package flags

import (
	"flag"
	"fmt"
	"github.com/davidv171/release-subscriber/cmd"
	"os"
)

type DownloadCommand struct {
	repo        string
	owner       string
	newest      bool
	destination string
}

func Parse() {

	dw := flag.NewFlagSet("download", flag.ExitOnError)

	repo := dw.String("repo", "",
		"The repo name, https://github.com/google/go-github/ 's repo name is go-github")

	owner := dw.String("owner", "",
		"The repo owner,  https://github.com/google/go-github/ 's repo name is google")

	newest := dw.Bool("newest", true,
		"If you want to download the newest release of a repo, requires -owner AND(require both) "+
			"-repo flags to be populated, defaults to true if no value is given")

	destination := dw.String("destination", ".",
		"The location of the download, optional arg, defaults to current directory(.)")

	dw.Parse(os.Args[2:])

	command := DownloadCommand{
		repo:        *repo,
		owner:       *owner,
		newest:      *newest,
		destination: *destination,
	}

	fmt.Println("Getting newest realease for : ", command.owner, command.repo)
	//Trigger the download
	download(&command)
}

func download(command *DownloadCommand) {

	if command.newest {
		if command.repo == "" || command.owner == "" {
			fmt.Println("Missing owner or repo flag, example: gpm -repo=go-github -owner=google -newest=true")
			os.Exit(1)
		}
		cmd.GetNewestRelease(command.owner, command.repo, command.destination)
	}
}
