package main

import (
	"flag"
	"fmt"

	"github.com/goodbaikin/scci/pkg/compute"
	"github.com/goodbaikin/scci/pkg/readfs"
)

func main() {
	r := flag.Float64("r", 1.96, "bias of the error")
	n := flag.Int("n", 1, "number of the significant digits.\nIf this value is smaller than 1, you'll get an non-formatted value.")
	flag.Parse()

	path := flag.Args()[0]
	data, err := readfs.Read(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	avg := compute.Average(data)
	error := compute.Error(data, *r)
	output := compute.Format(avg, error, *n)
	fmt.Println(output)
}
