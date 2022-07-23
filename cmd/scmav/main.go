package main

import (
	"flag"
	"fmt"
	"scmav/pkg/compute"
	"scmav/pkg/readfs"
)

func main() {
	r := flag.Float64("r", 1.96, "bias of the error")
	flag.Parse()

	path := flag.Args()[0]
	data, err := readfs.Read(path)
	if err != nil {
		fmt.Printf("Could not open file: %s.\n", path)
		return
	}

	avg := compute.Average(data)
	error := compute.Error(data, *r)
	fmt.Printf("%f±%f\n", avg, error)
}
