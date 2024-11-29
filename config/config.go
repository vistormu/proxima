package config

import (
    "os"

    // "proxima/errors"

    "github.com/BurntSushi/toml"
)

const DefaultConfig = `[parser]
line_break_value = "\n"
double_line_break_value = "\n\n"

[evaluator]
python = "python3"
begin_with = ""
end_with = ""
# text_replacements = [
#     { from = "", to = "" },
# ]

[components]
path = "components/"
use_modules = false
exclude = []
`

// ======
// PARSER
// ======
type ParserConfig struct {
    LineBreakValue string `toml:"line_break_value"`
    DoubleLineBreakValue string `toml:"double_line_break_value"`
}

// =========
// EVALUATOR
// =========
type TextReplacement struct {
    From string `toml:"from"`
    To string `toml:"to"`
}

type EvaluatorConfig struct {
    Python string `toml:"python"`
    BeginWith string `toml:"begin_with"`
    EndWith string `toml:"end_with"`
    TextReplacements []TextReplacement `toml:"text_replacements"`
}

// ==========
// COMPONENTS
// ==========
type ComponentsConfig struct {
    Path string `toml:"path"`
    UseModules bool `toml:"use_modules"`
    Exclude []string `toml:"exclude"`
}

type Config struct {
    Parser ParserConfig `toml:"parser"`
    Evaluator EvaluatorConfig `toml:"evaluator"`
    Components ComponentsConfig `toml:"components"`
}

func GetDefaultConfig() *Config {
    return &Config{
        Parser: ParserConfig{
            LineBreakValue: "\n",
            DoubleLineBreakValue: "\n\n",
        },
        Evaluator: EvaluatorConfig{
            Python: "python3",
            BeginWith: "",
            EndWith: "",
            TextReplacements: []TextReplacement{},
        },
        Components: ComponentsConfig{
            Path: "components/",
            UseModules: false,
            Exclude: []string{},
        },
    }
}

func LoadConfig() (*Config, error) {
    config := GetDefaultConfig()
    
    if _, err := os.Stat("proxima.toml"); os.IsNotExist(err) {
        return config, nil
    }

    _, err := toml.DecodeFile("proxima.toml", config)

    // return config, errors.New(errors.CONFIG, err.Error())
    return config, err
}
