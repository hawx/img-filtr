package main

import (
	"github.com/hawx/img-filtr/recipes"
	"github.com/hawx/img/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Run    func(cmd *Command, args []string)
	Usage  string
	Short  string
	Long   string
}

func runBrdl(cmd *Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Brdl(img)
	utils.WriteStdout(img)
}

func runDthr(cmd *Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Dthr(img)
	utils.WriteStdout(img)
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

var commands = []*Command{
	&Command{
		Run:   runBrdl,
		Usage: "brdl [options]",
		Short: "the smallest possible transmittable unit",
	  Long: `
  Repaints the image with a random dominant colour.
`,
	},
	&Command{
		Run:   runDthr,
		Usage: "dthr [options]",
		Short: "",
	  Long: `
...
`,
	},
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
