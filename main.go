package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version = "0.1.0"
)

var (
	flagVersion = flag.Bool("version", false, "Prints version and exists")
	flagHelp    = flag.Bool("help", false, "Prints this help and exits")
)

func main() {
	flag.Parse()
	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}
	if *flagHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("Ninja Virtual Machine started")

	fmt.Println("Ninja Virtual Machine stopped")
}
