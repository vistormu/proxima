package config

import (
    "os"
    "github.com/BurntSushi/toml"
)

const DefaultConfig = `[parser]
line_break_value = "\n"
double_line_break_value = "\n\n"

[runtimes]
python = "python3"
javascript = "node"
lua = "lua"
ruby = "ruby"

[evaluator]
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

// ========
// RUNTIMES
// ========
type RuntimesConfig struct {
    Python string `toml:"python"`
    JavaScript string `toml:"javascript"`
    Lua string `toml:"lua"`
    Ruby string `toml:"ruby"`
}

// =========
// EVALUATOR
// =========
type TextReplacement struct {
    From string `toml:"from"`
    To string `toml:"to"`
}

type EvaluatorConfig struct {
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
    Runtimes RuntimesConfig `toml:"runtimes"`
    Evaluator EvaluatorConfig `toml:"evaluator"`
    Components ComponentsConfig `toml:"components"`
}

func defaultConfig() *Config {
    return &Config{
        Parser: ParserConfig{
            LineBreakValue: "\n",
            DoubleLineBreakValue: "\n\n",
        },
        Runtimes: RuntimesConfig{
            Python: "python3",
            JavaScript: "node",
            Lua: "lua",
            Ruby: "ruby",
        },
        Evaluator: EvaluatorConfig{
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
    config := defaultConfig()
    
    if _, err := os.Stat("proxima.toml"); os.IsNotExist(err) {
        return config, nil
    }

    _, err := toml.DecodeFile("proxima.toml", config)
    return config, err
}
