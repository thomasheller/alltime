package main

import (
	"fmt"
	"os"
	"time"

	"github.com/thomasheller/alltime"
)

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		usage()
	}

	ds := alltime.Durations(d)

	for _, line := range ds {
		fmt.Println(line)
	}
}

func usage() {
	fmt.Println("usage: alltime [duration] (e.g. 120s, 40m, 360h)")
	os.Exit(1)
}
