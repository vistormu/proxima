package parser

import (
    "fmt"
    "testing"

    "proxima/tokenizer"
    "proxima/ast"
    "proxima/config"
)

// ========
// CHECKERS
// ========
func checkParserErrors(t *testing.T, parser *Parser) {
    if len(parser.Errors) == 0 {
        return
    }

    for _, err := range parser.Errors {
        t.Errorf(err.Error())
    }

    t.FailNow()
}

func checkExpressionLength(t *testing.T, expressions []ast.Expression, expected int) {
    if len(expressions) != expected {
        message := fmt.Sprintf("Wrong number of expressions\nexpected=%d\ngot=%d", expected, len(expressions))

        for i, expression := range expressions {
            if _, ok := expression.(*ast.Text); ok {
                message += fmt.Sprintf("\ntext %d: %s", i, expression.(*ast.Text).Value)
                continue
            } else if _, ok := expression.(*ast.Tag); ok {
                message += fmt.Sprintf("\ntag %d: %s", i, expression.(*ast.Tag).Name)
            } else {
                message += fmt.Sprintf("\nunknown %d\n", i)
            }
        }

        t.Fatalf(message)
    }
}

func checkExpressionIsText(t *testing.T, expression ast.Expression, expected string) {
    text, ok := expression.(*ast.Text)
    if !ok {
        t.Fatalf("Wrong expression type\nexpected=%T\ngot=%T", &ast.Text{}, expression)
    }

    if text.Value != expected {
        t.Fatalf("Wrong text value\nexpected=%s\ngot=%s", expected, text.Value)
    }
}

func checkExpressionIsTag(t *testing.T, expression ast.Expression, expectedName string, expectedArgs int) *ast.Tag {
    tag, ok := expression.(*ast.Tag)
    if !ok {
        t.Fatalf("Wrong expression type\nexpected=%T\ngot=%T", &ast.Tag{}, expression)
    }

    if tag.Name != expectedName {
        t.Fatalf("Wrong tag name\nexpected=%s\ngot=%s", expectedName, tag.Name)
    }

    if len(tag.Args) != expectedArgs {
        t.Fatalf("Wrong number of arguments\nexpected=%d\ngot=%d", expectedArgs, len(tag.Args))
    }

    return tag
}

func checkArgument(t *testing.T, argument ast.Argument, expectedName string, expectedValue []string) {
    if argument.Name != expectedName {
        t.Fatalf("Wrong argument name\nexpected=%s\ngot=%s", expectedName, argument.Name)
    }

    if len(argument.Values) != len(expectedValue) {
        message := fmt.Sprintf("Wrong number of argument values\nexpected=%d\ngot=%d", len(expectedValue), len(argument.Values))

        for i, value := range argument.Values {
            message += fmt.Sprintf("\nvalue %d: %s", i, value.(*ast.Text).Value)
        }

        t.Fatalf(message)
    }

    for i, value := range expectedValue {
        if argument.Values[i].(*ast.Text).Value != value {
            t.Fatalf("Wrong argument value\nexpected=%s\ngot=%s", value, argument.Values[i])
        }
    }
}

// =======
// HELPERS
// =======
func getExpressions(t *testing.T, input string) []ast.Expression {
    tokenizer := tokenizer.New([]rune(input))
    tokens := tokenizer.Tokenize()

    parser := New(tokens, "test", config.GetDefaultConfig())
    expressions := parser.Parse()

    checkParserErrors(t, parser)

    return expressions
}

// =====
// TESTS
// =====
func TestText(t *testing.T) {
    input := `This is a text`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    checkExpressionIsText(t, expressions[0], "This is a text")
}

func TestDoubleText(t *testing.T) {
    input := `This is a text
    This is another text`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    checkExpressionIsText(t, expressions[0], "This is a text\nThis is another text")
}

func TestTagWithOneArg(t *testing.T) {
    input := `@tag{arg}`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 1)

    checkArgument(t, tag.Args[0], "", []string{"arg"})
}

func TestTagWithNoArgs(t *testing.T) {
    input := `@tag{}`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 1)

    checkArgument(t, tag.Args[0], "", []string{})
}

func TestTagWithLinebreak(t *testing.T) {
    input := `@tag{
        arg
    }`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 1)

    checkArgument(t, tag.Args[0], "", []string{"arg\n"})
}

func TestTagWithNamedArgs(t *testing.T) {
    input := `@tag{ # with a random comment
        <name> value
    }`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 1)

    checkArgument(t, tag.Args[0], "name", []string{"value\n"})
}

func TestTagWithMultipleArgs(t *testing.T) {
    input := `@tag{arg1}{arg2}{arg3}`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 3)

    checkArgument(t, tag.Args[0], "", []string{"arg1"})
    checkArgument(t, tag.Args[1], "", []string{"arg2"})
    checkArgument(t, tag.Args[2], "", []string{"arg3"})
}

func TestTagWithOneMiddleEmptyArg(t *testing.T) {
    input := `@tag{arg1}{}{arg3}`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 3)
    
    checkArgument(t, tag.Args[0], "", []string{"arg1"})
    checkArgument(t, tag.Args[1], "", []string{})
    checkArgument(t, tag.Args[2], "", []string{"arg3"})
}

func TestTagWithMultipleArgsAndNamedArgs(t *testing.T) {
    input := `@tag{
        <name1> value1
    }{
        <name2> value2
    }{
        <name3> value3
    }`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 3)

    checkArgument(t, tag.Args[0], "name1", []string{"value1\n"})
    checkArgument(t, tag.Args[1], "name2", []string{"value2\n"})
    checkArgument(t, tag.Args[2], "name3", []string{"value3\n"})
}

func TestTagWithEmptyNamedArg(t *testing.T) {
    input := `@tag{
        <name1> value1
    }{
        <name2>
    }`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 2)

    checkArgument(t, tag.Args[0], "name1", []string{"value1\n"})
    checkArgument(t, tag.Args[1], "name2", []string{})
}

func TestErrors(t *testing.T) {
    input := `@tag{
        <name1> value1
    }{
        <name2> value2
    }{
        <name3> value3
    }`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 3)

    checkArgument(t, tag.Args[0], "name1", []string{"value1\n"})
    checkArgument(t, tag.Args[1], "name2", []string{"value2\n"})
    checkArgument(t, tag.Args[2], "name3", []string{"value3\n"})
}

func TestNestedTags(t *testing.T) {
    input := `@tag{
        @nested{}
    }`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 2)
    tag := checkExpressionIsTag(t, expressions[0], "tag", 1)

    checkExpressionLength(t, tag.Args[0].Values, 1)
    tagArg := checkExpressionIsTag(t, tag.Args[0].Values[0], "nested", 1)
    
    checkArgument(t, tagArg.Args[0], "", []string{})
}

func TestNestedTags2(t *testing.T) {
    input := `@test{This is a @test2{}}`
    expressions := getExpressions(t, input)

    checkExpressionLength(t, expressions, 1)
    tag := checkExpressionIsTag(t, expressions[0], "test", 1)

    checkExpressionLength(t, tag.Args[0].Values, 2)
    checkExpressionIsText(t, tag.Args[0].Values[0], "This is a ")

    tagArg := checkExpressionIsTag(t, tag.Args[0].Values[1], "test2", 1)
    checkArgument(t, tagArg.Args[0], "", []string{})
}
