package main

import (
	"github.com/hawx/hadfield"
	"github.com/hawx/img-filtr/recipes"
	"github.com/hawx/img/utils"
	"fmt"
	"os"
	"image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
)


func runBrdl(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	img = recipes.Brdl(img)
	utils.WriteStdout(img, data)
}

func runDazd(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	img = recipes.Dazd(img)
	utils.WriteStdout(img, data)
}

func runDthr(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	img = recipes.Dthr(img)
	utils.WriteStdout(img, data)
}

func runEdwn(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	img = recipes.Edwn(img)
	utils.WriteStdout(img, data)
}

func runHeathr(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	if len(args) < 1 {
		utils.Warn("Need an<other> image to compose with!")
		os.Exit(2)
	}

	file, err := os.Open(args[0])
	if err != nil {
		utils.Warn("Problem opening", args[0])
		os.Exit(2)
	}

	other, _, err := image.Decode(file)
	if err != nil {
		utils.Warn("Could not decode image, must be jpeg/png/gif.")
		os.Exit(2)
	}

	img = recipes.Heathr(img, other)
	utils.WriteStdout(img, data)
}

func runPostcrd(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	img = recipes.Postcrd(img)
	utils.WriteStdout(img, data)
}

func runPostr(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	img = recipes.Postr(img)
	utils.WriteStdout(img, data)
}

func runRockstr(cmd *hadfield.Command, args []string) {
	img, data := utils.ReadStdin()
	img = recipes.Rockstr(img)
	utils.WriteStdout(img, data)
}

var commands = hadfield.Commands{
	&hadfield.Command{
		Run:   runBrdl,
		Usage: "brdl",
		Short: "the smallest possible transmittable unit",
	  Long: `
  Repaints the image with a random dominant colour.
`,
	},
	&hadfield.Command{
		Run:   runDazd,
		Usage: "dazd",
		Short: "blown out",
	  Long: `
  Lightens, saturates and increases the contrast.
`,
	},
	&hadfield.Command{
		Run:   runDthr,
		Usage: "dthr",
		Short: "dithered",
	  Long: `
  Turns the image into a dithered pattern of white and black pixels.
`,
	},
	&hadfield.Command{
		Run:  runEdwn,
		Usage: "edwn",
		Short: "polaroid",
		Long: `
  Cuts the photo into a square and sticks it onto a simplistic polaroid background.
`,
	},
	&hadfield.Command{
		Run:   runHeathr,
		Usage: "heathr <other>",
		Short: "polaroids, side-by-side",
	  Long: `
  Cuts the photos into squares and sticks them onto simple polaroid-esque backings.
`,
	},
	&hadfield.Command{
		Run:   runPostcrd,
		Usage: "postcrd",
		Short: "saturated",
	  Long: `
  Saturates, lightens and blurs the image.
`,
	},
	&hadfield.Command{
		Run:   runPostr,
		Usage: "postr",
		Short: "dark and saturated",
	  Long: `
  Like a darker, subtler version of postcrd.
`,
	},
	&hadfield.Command{
		Run:   runRockstr,
		Usage: "rockstr",
		Short: "etched",
	  Long: `
  Almost etches the image onto the screen in pencil.
`,
	},
}

var templates = hadfield.Templates{
Usage: `Usage: img filtr [command] [arguments]

  An implementation of straup/filtr as a single executable.

  Commands: {{range .}}
    {{.Name | printf "%-15s"}} # {{.Short}}{{end}}

`,
Help: `Usage: img filtr {{.Usage}}
{{.Long}}
`,
}

const (
	USAGE = "filtr [command] [options]"
	SHORT = "reimplementation of straup/filtr"
)

var longTemplate hadfield.Template = `  This is a description.

  Commands: {{range .}}
    {{.Name | printf "%-15s"}} # {{.Short}}{{end}}
`

func main() {
	os.Args = utils.GetOutput(os.Args)
	args := os.Args

	if len(args) == 2 && args[1] == "--long" {
		longTemplate.Render(os.Stdout, commands.Data())
	} else if len(args) == 2 && args[1] == "--short" {
		fmt.Println(SHORT)
	} else if len(args) == 2 && args[1] == "--usage" {
		fmt.Println(USAGE)
	} else {
		hadfield.Run(commands, templates)
	}

	os.Exit(0)
}
