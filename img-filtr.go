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
	img := utils.ReadStdin()
	img  = recipes.Brdl(img)
	utils.WriteStdout(img)
}

func runDazd(cmd *hadfield.Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Dazd(img)
	utils.WriteStdout(img)
}

func runDthr(cmd *hadfield.Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Dthr(img)
	utils.WriteStdout(img)
}

func runEdwn(cmd *hadfield.Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Edwn(img)
	utils.WriteStdout(img)
}

func runHeathr(cmd *hadfield.Command, args []string) {
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

func runPostcrd(cmd *hadfield.Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Postcrd(img)
	utils.WriteStdout(img)
}

func runPostr(cmd *hadfield.Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Postr(img)
	utils.WriteStdout(img)
}

func runRockstr(cmd *hadfield.Command, args []string) {
	img := utils.ReadStdin()
	img  = recipes.Rockstr(img)
	utils.WriteStdout(img)
}

var commands = hadfield.Commands{
	&hadfield.Command{
		Run:   runBrdl,
		Usage: "brdl [options]",
		Short: "the smallest possible transmittable unit",
	  Long: `
  Repaints the image with a random dominant colour.
`,
	},
	&hadfield.Command{
		Run:   runDazd,
		Usage: "dazd [options]",
		Short: "",
	  Long: `
...
`,
	},
	&hadfield.Command{
		Run:   runDthr,
		Usage: "dthr [options]",
		Short: "",
	  Long: `
...
`,
	},
	&hadfield.Command{
		Run:  runEdwn,
		Usage: "edwn [options]",
		Short: "",
		Long: `
...
`,
	},
	&hadfield.Command{
		Run:   runHeathr,
		Usage: "heathr <right> [options]",
		Short: "",
	  Long: `
...
`,
	},
	&hadfield.Command{
		Run:   runPostcrd,
		Usage: "postcrd [options]",
		Short: "",
	  Long: `
...
`,
	},
	&hadfield.Command{
		Run:   runPostr,
		Usage: "postr [options]",
		Short: "",
	  Long: `
...
`,
	},
	&hadfield.Command{
		Run:   runRockstr,
		Usage: "rockstr [options]",
		Short: "",
	  Long: `
...
`,
	},
}

var templates = hadfield.Templates{
Usage: `Usage: img filtr [command] [arguments]

  This is a description.

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
	args := os.Args

	if len(args) > 1 && args[1] == "--long" {
		longTemplate.Render(os.Stdout, commands.Data())
	} else if len(args) > 1 && args[1] == "--short" {
		fmt.Println(SHORT)
	} else if len(args) > 1 && args[1] == "--usage" {
		fmt.Println(USAGE)
	} else {
		hadfield.Run(commands, templates)
	}

	os.Exit(0)
}
