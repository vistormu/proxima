package config

import (
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

[components]
components_dir = "./components/"
use_modules = false
exclude = []
`

type ParserConfig struct {
    LineBreakValue *string `toml:"line_break_value"`
    DoubleLineBreakValue *string `toml:"double_line_break_value"`
}

type EvaluatorConfig struct {
    PythonCmd *string `toml:"python_cmd"`
    JavaScriptCmd *string `toml:"javascript_cmd"`
    LuaCmd *string `toml:"lua_cmd"`
    RubyCmd *string `toml:"ruby_cmd"`
}

type ComponentsConfig struct {
    ComponentsDir *string `toml:"components_dir"`
    UseModules *bool `toml:"use_modules"`
    Exclude []string `toml:"exclude"`
}

type Config struct {
    Parser ParserConfig `toml:"parser"`
    Evaluator EvaluatorConfig `toml:"evaluator"`
    Components ComponentsConfig `toml:"components"`
}

func (c *Config) setDefaults() {
    // parser
    if c.Parser.LineBreakValue == nil {
        defaultLineBreakValue := "\n"
        c.Parser.LineBreakValue = &defaultLineBreakValue
    }
    if c.Parser.DoubleLineBreakValue == nil {
        defaultDoubleLineBreakValue := "\n\n"
        c.Parser.DoubleLineBreakValue = &defaultDoubleLineBreakValue
    }
    
    // evaluator
    if c.Evaluator.PythonCmd == nil {
        defaultPythonCmd := "python3 -c"
        c.Evaluator.PythonCmd = &defaultPythonCmd
    }
    if c.Evaluator.JavaScriptCmd == nil {
        defaultJavaScriptCmd := "node -e"
        c.Evaluator.JavaScriptCmd = &defaultJavaScriptCmd
    }
    if c.Evaluator.LuaCmd == nil {
        defaultLuaCmd := "lua -e"
        c.Evaluator.LuaCmd = &defaultLuaCmd
    }
    if c.Evaluator.RubyCmd == nil {
        defaultRubyCmd := "ruby -e"
        c.Evaluator.RubyCmd = &defaultRubyCmd
    }

    // components
    if c.Components.ComponentsDir == nil {
        defaultComponentsDir := "./components/"
        c.Components.ComponentsDir = &defaultComponentsDir
    }
    if c.Components.UseModules == nil {
        defaultUseModules := false
        c.Components.UseModules = &defaultUseModules
    }
    if c.Components.Exclude == nil {
        defaultExclude := []string{}
        c.Components.Exclude = defaultExclude
    }
}

func LoadConfig() (*Config, error) {
    var config Config
    _, err := toml.DecodeFile("proxima.toml", &config)
    if err != nil {
        return nil, err
    }

    config.setDefaults()

    return &config, err
}
