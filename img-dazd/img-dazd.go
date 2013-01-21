package main

import (
	"github.com/hawx/img/utils"
	"github.com/hawx/img/unsharp" // doesn't exist!
	"github.com/hawx/img/brightness"
	"github.com/hawx/img/hsla"
	"github.com/hawx/img/contrast"
	"github.com/hawx/img/blur" // doesn't exist!
	"fmt"
	"flag"
)

func main() {
	var long  = flag.Bool("long", false, "")
	var short = flag.Bool("short", false, "")
	var usage = flag.Bool("usage", false, "")

	flag.Parse()

	if *long {
		fmt.Println("  ")

	} else if *short {
		fmt.Println("")

	} else if *usage {
		fmt.Println("dazd [options]")

	} else {
		img := utils.ReadStdin()

		// sharpen the image (does not exist yet!)
		img = unsharp.Apply(img, 1.5, 1.5)

		img = brightness.Adjust(img, utils.Multiplier(1.75))
		img = hsla.Saturation(img, utils.Multiplier(1.5))

		// guessing amount here!
		img = contrast.Adjust(img, utils.Multiplier(1.3))

		// gaussian with radius=1, sigma=2
		img = blur.Gaussian(img, 1, 2)

		utils.WriteStdout(img)
	}
}
