package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	parseFlags()
	if lFlag {
		fmt.Println("gonna list all")
		return
	}
}

func printHeader() {
	fmt.Print(
		`
Tada - a CLI todo application
=============================

`)
}

func showHelp() {
	printHeader()
	flag.PrintDefaults()
	os.Exit(0)
}
