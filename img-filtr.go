package main

import (
	"github.com/hawx/img-filtr/recipes"
	"github.com/hawx/img/utils"
	"flag"
	"fmt"
	"os"
	"strings"
	"image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
)

type Command struct {
	Run    func(cmd *Command, args []string)
	Usage  string
	Short  string
	Long   string
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}


func runBrdl(cmd *Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Brdl(img)
	utils.WriteStdout(img)
}

// func runDazd(cmd *Command, args []string) {
// 	img := utils.ReadStdin()
// 	img  = recipes.Dazd(img)
// 	utils.WriteStdout(img)
// }

func runDthr(cmd *Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Dthr(img)
	utils.WriteStdout(img)
}

func runEdwn(cmd *Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Edwn(img)
	utils.WriteStdout(img)
}

func runHeathr(cmd *Command, args []string) {
	img := utils.ReadStdin()
	if len(args) < 2 {
		utils.Warn("Need an<other> image to compose with!")
		os.Exit(2)
	}

	file, err := os.Open(args[1])
	if err != nil {
		utils.Warn("Problem opening", args[1])
		os.Exit(2)
	}

	other, _, err := image.Decode(file)
	if err != nil {
		utils.Warn("Could not decode image, must be jpeg/png/gif.")
		os.Exit(2)
	}

	img = recipes.Heathr(img, other)
	utils.WriteStdout(img)
}

// func runPostcrd(cmd *Command, args []string) {
// 	img := utils.ReadStdin()
// 	img  = recipes.Postcrd(img)
// 	utils.WriteStdout(img)
// }

// func runPostr(cmd *Command, args []string) {
// 	img := utils.ReadStdin()
// 	img  = recipes.Postr(img)
// 	utils.WriteStdout(img)
// }

var commands = []*Command{
	&Command{
		Run:   runBrdl,
		Usage: "brdl [options]",
		Short: "the smallest possible transmittable unit",
	  Long: `
  Repaints the image with a random dominant colour.
`,
	},
// 	&Command{
// 		Run:   runDazd,
// 		Usage: "dazd [options]",
// 		Short: "",
// 	  Long: `
// ...
// `,
// 	},
	&Command{
		Run:   runDthr,
		Usage: "dthr [options]",
		Short: "",
	  Long: `
...
`,
	},
	&Command{
		Run:  runEdwn,
		Usage: "edwn [options]",
		Short: "",
		Long: `
...
`,
	},
	&Command{
		Run:   runHeathr,
		Usage: "heathr <right> [options]",
		Short: "",
	  Long: `
...
`,
	},
// 	&Command{
// 		Run:   runPostcrd,
// 		Usage: "postcrd [options]",
// 		Short: "",
// 	  Long: `
// ...
// `,
// 	},
// 	&Command{
// 		Run:   runPostr,
// 		Usage: "postr [options]",
// 		Short: "",
// 	  Long: `
// ...
// `,
// 	},
}

func main() {
	// var long  = flag.Bool("long",  false, "")
	// var short = flag.Bool("short", false, "")
	// var usage = flag.Bool("usage", false, "")

	flag.Parse()
	args := flag.Args()

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			if len(args) > 1 && args[1] == "--long" {
				fmt.Println(cmd.Long)
			} else if len(args) > 1 && args[1] == "--short" {
				fmt.Println(cmd.Short)
			} else if len(args) > 1 && args[1] == "--usage" {
				fmt.Println(cmd.Usage)
			} else {
				cmd.Run(cmd, args)
			}
			os.Exit(0)
			return
		}
	}

	os.Exit(2)
}
