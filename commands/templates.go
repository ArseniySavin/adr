package commands

import (
	"adr/templates"

	_ "embed"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var TemplatesCommand = cli.Command{
	Name:        "templates",
	Aliases:     []string{"t"},
	Usage:       "Show embedded templates",
	UsageText:   "adr templates",
	Description: "You can add your template. You should make the request in the repository.",
	Action: func(c *cli.Context) error {

		files, err := templates.TemplatesFs.ReadDir(".")
		if err != nil {
			return err
		}

		color.Green("Embedded templates:")
		for _, v := range files {
			color.Cyan(v.Name())
		}

		return nil
	},
}
