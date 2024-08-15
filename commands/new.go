package commands

import (
	"adr/config"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var NewCommand = cli.Command{
	Name:    "new",
	Aliases: []string{"c"},
	Usage:   "Create a new ADR",
	Flags:   newFlags(),
	Action: func(c *cli.Context) error {

		cfg := config.DefaultConfig()

		cfg = cfg.Get()

		cfg.Sequence++

		cfg.Update(cfg)
		err := newAdr(cfg, c.String(title), c.String(status), cfg.Author)
		if err != nil {
			return err
		}
		return nil
	},
}

type AdrStatus string

// TODO need create commands for a changing status
const (
	PROPOSED   AdrStatus = "Proposed"
	ACCEPTED   AdrStatus = "Accepted"
	DEPRECATED AdrStatus = "Deprecated"
	SUPERSEDED AdrStatus = "Superseded"
)

type Adr struct {
	Author   string
	Title    string
	Date     string
	Status   AdrStatus
	Sequence int
}

func newAdr(cfg config.AdrConfig, title, status, author string) error {
	adr := Adr{
		Title:    strings.Trim(title, "\n \t"),
		Date:     time.Now().Format("02-01-2006 15:04:05"),
		Sequence: cfg.Sequence,
		Status:   AdrStatus(status),
		Author:   author,
	}

	adrTemplate := fmt.Sprintf("%s/%s.%s.md", cfg.AdrConfigPath(), cfg.Template, cfg.Language)
	template, err := template.ParseFiles(adrTemplate)
	if err != nil {
		return err
	}
	adrFileName := fmt.Sprintf("%d.%s-%s.md", adr.Sequence, cfg.Prefix, strings.Join(strings.Split(adr.Title, " "), "-"))

	adrFile := filepath.Join(cfg.DocsAdrPath(), adrFileName)

	if _, err := os.Stat(cfg.DocsPath()); os.IsNotExist(err) {
		os.Mkdir(cfg.DocsPath(), 0744)
	}
	if _, err := os.Stat(cfg.DocsAdrPath()); os.IsNotExist(err) {
		os.Mkdir(cfg.DocsAdrPath(), 0744)
	}

	f, err := os.Create(adrFile)
	if err != nil {
		f.Close()
		os.Remove(adrFile)
		return err
	}
	err = template.Execute(f, adr)
	if err != nil {
		color.Red(fmt.Sprintf("Template error %s", err.Error()))
		f.Close()
		os.Remove(adrFile)
		return err
	}

	color.Green(fmt.Sprintf("ADR number %d was successfully written to : %s ", adr.Sequence, adrFile))
	return nil
}
