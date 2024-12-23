package commands

import (
    "testing"
    // "fmt"
)

func TestFuzzy(t *testing.T) {
    wrongCommands := []string{
        "inis",
        "held",
        "verion",
        "makee",
        "male",
    }

    correctCommands := []string{
        "init",
        "help",
        "version",
        "make",
        "make",
    }

    keys := keys(commands)

    t.Logf("\nfinding closest match for %v\n", keys)

    for i, wrongCommand := range wrongCommands {
        closestMatch := findClosestMatch(wrongCommand, keys)
        if closestMatch != correctCommands[i] {
            t.Errorf("\n-> wrong match\n   |> wrong: %s\n   |> expected: %v\n   |> fuzzy: %v\n", wrongCommand, correctCommands[i], closestMatch)
        }
    }
}
