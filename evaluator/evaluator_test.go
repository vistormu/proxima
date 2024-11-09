package evaluator

import (
    "fmt"
    "testing"
    "proxima/config"
)

func TestInterpreter(t *testing.T) {
    // Example usage
    config := config.GetDefaultConfig()
    interpreter, err := NewInterpreter(config)
    if err != nil {
        t.Error(err)
    }
    defer interpreter.Close()

    args := []struct{ Name string; Value string }{
        {"arg1", "value1"},
    }
    component := Component{
        language: PYTHON,
        content:  "def test_func(arg1): return arg1",
        name:     "test_func",
    }

    result, err := interpreter.Evaluate(args, component)
    if err != nil {
        t.Error(err)
    }
    fmt.Println("Result:", result)
}
