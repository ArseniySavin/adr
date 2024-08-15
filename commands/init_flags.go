package commands

import (
	"github.com/urfave/cli/v2"
)

const (
	language      = "language"
	prefix        = "prefix"
	template_name = "template"
)

func initFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    language,
			Aliases: []string{"l"},
			Value:   "en",
		},
		&cli.StringFlag{
			Name:    prefix,
			Aliases: []string{"pf"},
			Value:   "ADR",
		},
		&cli.StringFlag{
			Name:    template_name,
			Aliases: []string{"tm"},
			Value:   "short",
		},
	}
}
