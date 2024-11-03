package config

import (
    "os"
    "github.com/BurntSushi/toml"
)

const DefaultConfig = `[parser]
line_break_value = "\n"
double_line_break_value = "\n\n"

[evaluator]
python_cmd = "python3 -c"
javascript_cmd = "node -e"
lua_cmd = "lua -e"
ruby_cmd = "ruby -e"
# text_replacement = [
#     { from = "", to = "" },
# ]

[components]
components_dir = "./components/"
use_modules = false
exclude = []
`

type ParserConfig struct {
    LineBreakValue string `toml:"line_break_value"`
    DoubleLineBreakValue string `toml:"double_line_break_value"`
}

type TextReplacement struct {
    From string `toml:"from"`
    To string `toml:"to"`
}

type EvaluatorConfig struct {
    PythonCmd string `toml:"python_cmd"`
    JavaScriptCmd string `toml:"javascript_cmd"`
    LuaCmd string `toml:"lua_cmd"`
    RubyCmd string `toml:"ruby_cmd"`
    TextReplacements []TextReplacement `toml:"text_replacement"`
}

type ComponentsConfig struct {
    ComponentsDir string `toml:"components_dir"`
    UseModules bool `toml:"use_modules"`
    Exclude []string `toml:"exclude"`
}

type Config struct {
    Parser ParserConfig `toml:"parser"`
    Evaluator EvaluatorConfig `toml:"evaluator"`
    Components ComponentsConfig `toml:"components"`
}

func defaultConfig() *Config {
    return &Config{
        Parser: ParserConfig{
            LineBreakValue: "\n",
            DoubleLineBreakValue: "\n\n",
        },
        Evaluator: EvaluatorConfig{
            PythonCmd: "python3 -c",
            JavaScriptCmd: "node -e",
            LuaCmd: "lua -e",
            RubyCmd: "ruby -e",
            TextReplacements: []TextReplacement{},
        },
        Components: ComponentsConfig{
            ComponentsDir: "./components/",
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
