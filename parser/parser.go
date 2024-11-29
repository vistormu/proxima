package parser

import (
    "strings"

    "proxima/config"
	"proxima/ast"
	"proxima/errors"
	"proxima/token"
)

type Parser struct {
    tokens []token.Token
    position int
    nTokens int

    currentLine int
    file string

    lineBreakValue string
    doubleLineBreakValue string

    Errors []error
}

func New(tokens []token.Token, file string, config *config.Config) *Parser {
    return &Parser{
        tokens,
        -1,
        len(tokens),
        1,
        file,
        config.Parser.LineBreakValue,
        config.Parser.DoubleLineBreakValue,
        []error{},
    }
}

func (p *Parser) Parse() []ast.Expression {
    expressions := []ast.Expression{}

    for p.position < p.nTokens {
        // parse expression
        expression := p.parseExpression()
        if expression != nil {
            expressions = append(expressions, expression)
        }

        // check for line breaks
        c1 := p.rawPeekToken(1).Type == token.LINEBREAK
        c2 := p.rawPeekToken(2).Type == token.LINEBREAK

        if c1 && c2 { // double line break
            expressions = append(expressions, &ast.Text{Value: p.doubleLineBreakValue})
        } else if c1 { // single line break
            expressions = append(expressions, &ast.Text{Value: p.lineBreakValue})
        }
    }

    return mergeConsecutiveTexts(expressions)
}

func mergeConsecutiveTexts(exprs []ast.Expression) []ast.Expression {
    var result []ast.Expression
    var prevText *ast.Text

    for _, expr := range exprs {
        switch e := expr.(type) {
        case *ast.Text:
            if prevText != nil {
                prevText.Value += e.Value
            } else {
                prevText = &ast.Text{
                    Value:      e.Value,
                    LineNumber: e.LineNumber,
                }
                result = append(result, prevText)
            }

        case *ast.Tag:
            prevText = nil
            for i := range e.Args {
                e.Args[i].Values = mergeConsecutiveTexts(e.Args[i].Values)
            }

            result = append(result, e)

        default:
            prevText = nil
            result = append(result, e)
        }
    }

    return result
}

// =======
// HELPERS
// =======
func (p *Parser) readToken() token.Token {
    p.position++

    if p.position >= p.nTokens {
        return token.New(rune(0))
    }

    t := p.tokens[p.position]

    if t.Type == token.LINEBREAK {
        p.currentLine++
        t = p.readToken()
    }

    return t
}

func (p *Parser) currentToken() token.Token {
    if p.position >= p.nTokens {
        return token.New(rune(0))
    }

    return p.tokens[p.position]
}

func (p *Parser) peekToken() token.Token {
    if p.position + 1 >= p.nTokens {
        return token.New(rune(0))
    }

    index := p.position + 1
    peekToken := p.tokens[index]
    for peekToken.Type == token.LINEBREAK {
        index++
        if index >= p.nTokens {
            return token.New(rune(0))
        }

        peekToken = p.tokens[index]
    }

    return peekToken
}

func (p *Parser) rawPeekToken(position int) token.Token {
    if p.position + position >= p.nTokens {
        return token.New(rune(0))
    }

    return p.tokens[p.position + position]
}

func (p *Parser) addError(errorType errors.ErrorType, args ...any) {
    p.Errors = append(p.Errors, errors.NewParseError(errorType, p.file, p.currentLine, args...))
}

func (p *Parser) expect(tokenType token.TokenType) (token.Token, bool) {
    t := p.readToken()

    if t.Type != tokenType {
        p.addError(errors.EXPECTED_TOKEN, tokenType, t)
        return t, false
    }

    return t, true
}

// =======
// PARSERS
// =======
func (p *Parser) parseExpression() ast.Expression {
    t := p.readToken()

    switch t.Type {
    case token.TEXT:
        return &ast.Text{Value: t.Literal, LineNumber: p.currentLine}

    case token.TAG:
        return p.parseTag()

    case token.EOF:
        return nil

    default:
        return &ast.Text{Value: t.Literal, LineNumber: p.currentLine}
    }
}

func (p *Parser) parseTag() *ast.Tag {
    tag := &ast.Tag{LineNumber: p.currentLine}

    // tag name
    t, ok := p.expect(token.TEXT)
    if !ok {
        return nil
    }
    if strings.Contains(t.Literal, " ") {
        p.addError(errors.WRONG_TAG_NAME, t.Literal)
        return nil
    }
    tag.Name = t.Literal

    // opening tag
    _, ok = p.expect(token.LBRACE)
    if !ok {
        return nil
    }

    // tag arguments
    tag.Args = p.parseArguments() // entry: LBRACE, exit: RBRACE

    return tag
}

func (p *Parser) parseArguments() []ast.Argument {
    args := []ast.Argument{}

    // args
    for p.currentToken().Type != token.RBRACE {
        if p.currentToken().Type == token.EOF {
            p.addError(errors.UNCLOSED_TAG)
            return args
        }

        arg := p.parseArgument() // entry: LBRACE, exit: RBRACE
        args = append(args, arg)

        if p.peekToken().Type == token.LBRACE {
            p.readToken()
            continue
        }
    }

    return args
}

func (p *Parser) parseArgument() ast.Argument {
    arg := ast.Argument{}

    t := p.currentToken()
    var ok bool
    if p.peekToken().Type == token.LANGLE {
        p.readToken() 
        t, ok = p.expect(token.TEXT)
        if !ok {
            return ast.Argument{}
        }
        arg.Name = t.Literal

        t, ok = p.expect(token.RANGLE)
        if !ok {
            return ast.Argument{}
        }
    }
    
    // values
    for p.peekToken().Type != token.RBRACE {
        if p.currentToken().Type == token.EOF {
            p.addError(errors.UNCLOSED_TAG)
            return ast.Argument{}
        }

        expression := p.parseExpression()
        if expression != nil {
            arg.Values = append(arg.Values, expression)
        }

        // check for line breaks
        c1 := p.rawPeekToken(1).Type == token.LINEBREAK
        c2 := p.rawPeekToken(2).Type == token.LINEBREAK

        if c1 && c2 { // double line break
            arg.Values = append(arg.Values, &ast.Text{Value: p.doubleLineBreakValue})
        } else if c1 { // single line break
            arg.Values = append(arg.Values, &ast.Text{Value: p.lineBreakValue})
        }
    }

    _, ok = p.expect(token.RBRACE)
    if !ok {
        return ast.Argument{}
    }

    return arg
}

