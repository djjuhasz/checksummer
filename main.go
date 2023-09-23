package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println(`
Usage:
------
checksummer PATH1 [PATH2 [...]]`)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing PATH.")
		usage()
		os.Exit(1)
	}

	if err := Run(os.Args[1:]); err != nil {
		fmt.Printf("Fatal error: %v\n", err.Error())
		os.Exit(1)
	}
}
