package config

import (
	"io/fs"
	"os"

	"proxima/internal/assets"

	"github.com/BurntSushi/toml"
)

type ParserConfig struct {
	LineBreakValue       string `toml:"line_break_value"`
	DoubleLineBreakValue string `toml:"double_line_break_value"`
}

type TextReplacement struct {
	From string `toml:"from"`
	To   string `toml:"to"`
}

type EvaluatorConfig struct {
	Python           string            `toml:"python"`
	BeginWith        string            `toml:"begin_with"`
	EndWith          string            `toml:"end_with"`
	TextReplacements []TextReplacement `toml:"text_replacements"`
}

type ComponentsConfig struct {
	Path       string   `toml:"path"`
	UseModules bool     `toml:"use_modules"`
	Exclude    []string `toml:"exclude"`
}

type Config struct {
	Parser     ParserConfig     `toml:"parser"`
	Evaluator  EvaluatorConfig  `toml:"evaluator"`
	Components ComponentsConfig `toml:"components"`
}

func LoadConfig(path string) (*Config, error) {
	var f fs.File
	var err error

	f, err = os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = assets.DefaultConfig.Open("proxima.toml")
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	defer f.Close()

	cfg := new(Config)
	_, err = toml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
