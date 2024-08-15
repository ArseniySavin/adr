package commands

import (
	"github.com/urfave/cli/v2"
)

const (
	title  = "title"
	status = "status"
	author = "author"
)

func newFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     title,
			Aliases:  []string{"t"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    status,
			Aliases: []string{"s"},
			Value:   string(PROPOSED),
		},
	}
}
