package commands

import (
	"adr/config"
	"adr/internal/git"
	"adr/templates"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var InitCommand = cli.Command{
	Name:        "init",
	Aliases:     []string{"i"},
	Usage:       "Initializes the ADR configurations",
	UsageText:   "adr init",
	Description: "Initializes the ADR configuration in the project directory\n This is a a prerequisite to running any other adr sub-command",
	Flags:       initFlags(),
	Action: func(c *cli.Context) error {
		cfg := config.DefaultConfig()

		cfg.Language = c.String(language)
		cfg.Prefix = c.String(prefix)
		cfg.Template = c.String(template_name)

		initAdrFolder(cfg.AdrConfigPath())

		cfg.ProjectPath()

		cfg.Author = git.GitUser(cfg.ProjectPath())

		initConfig(cfg.ConfigPath(), cfg)

		template := fmt.Sprintf("%s.%s.md", cfg.Language, cfg.Template)
		templateGen := fmt.Sprintf("%s.%s.md", cfg.Template, cfg.Language)

		templateGen = filepath.Join(cfg.AdrConfigPath(), templateGen)
		initTemplate(template, templateGen)
		return nil
	},
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func initAdrFolder(dir string) {
	color.Green("Initializing adr at " + dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0744)
	} else {
		color.Red(dir + " already exists, skipping folder creation")
	}
}

func initConfig(dir string, cfg config.AdrConfig) {
	bytes, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		panic(err)
	}
	os.WriteFile(dir, bytes, 0644)
}

func initTemplate(template, templateGen string) {
	data, err := templates.TemplatesFs.ReadFile(template)
	if err != nil {
		color.Red("Embedded file does not exist " + err.Error())
	}

	if err := os.WriteFile(templateGen, data, 0644); err != nil {
		color.Red(templateGen + err.Error())
	}
}
