package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

// AdrConfig ADR configuration, loaded and used by each sub-command
type AdrConfig struct {
	Sequence int    `json:"sequence"`
	Language string `json:"language"`
	Template string `json:"template"`
	Prefix   string `json:"prefix"`
	Author   string `json:"author"`
	path     string
}

const (
	adr_config_folder = ".adr"
	docs_folder       = "docs"
	adrs_folder       = "adrs"
	template          = "template"
	config_file       = "config.json"
)

func DefaultConfig() AdrConfig {
	return AdrConfig{
		Language: "en",
		Sequence: 0,
		Template: template,
		Prefix:   "ADR",
		Author:   "No name",
		path:     path(),
	}
}

func path() string {
	curDir, _ := os.Getwd()
	return curDir
}

func (c *AdrConfig) ProjectPath() string {
	return c.path
}

func (c *AdrConfig) DocsPath() string {
	return filepath.Join(c.path, docs_folder)
}

func (c *AdrConfig) DocsAdrPath() string {
	return filepath.Join(c.path, docs_folder, adrs_folder)
}

func (c *AdrConfig) AdrConfigPath() string {
	return filepath.Join(c.path, adr_config_folder)
}

func (c *AdrConfig) ConfigPath() string {
	return filepath.Join(c.path, adr_config_folder, config_file)
}

func (c *AdrConfig) Get() AdrConfig {

	dir := filepath.Join(c.path, adr_config_folder, config_file)

	bytes, err := os.ReadFile(dir)
	if err != nil {
		color.Red("No ADR configuration is found!")
		color.HiGreen("Start by initializing ADR configuration, check 'adr init --help' for more help")
		os.Exit(1)
	}

	var cfg = DefaultConfig()
	json.Unmarshal(bytes, &cfg)

	return cfg
}

func (c *AdrConfig) Update(cfg AdrConfig) {

	dir := filepath.Join(c.path, adr_config_folder, config_file)

	bytes, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		panic(err)
	}
	os.WriteFile(dir, bytes, 0644)
}
