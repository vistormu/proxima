package parser

import (
    "fmt"
    "strings"
    "proxima/token"
    "proxima/tokenizer"
    "proxima/ast"
    "proxima/error"
)

type Parser struct {
    tokenizer *tokenizer.Tokenizer

    currentToken token.Token
    peekToken token.Token

    currentLine int

    Errors []error.Error
    File string
}

// PUBLIC
func New(input string, file string) *Parser {
    p := &Parser{tokenizer: tokenizer.New(input), currentLine: 1, File: file}
    p.nextToken()
    p.nextToken()
    return p
}
func (p *Parser) Parse() *ast.Document {
    document := &ast.Document{}

    for !p.currentTokenIs(token.EOF) {
        paragraph := p.parseParagraph()
        if len(paragraph.Content) > 0 {
            document.Paragraphs = append(document.Paragraphs, paragraph)
        }
    }

    return document
}

// TOKENS
func (p *Parser) nextToken() {
    p.currentToken = p.peekToken
    p.peekToken = p.tokenizer.GetToken()

    if p.currentTokenIs(token.LINEBREAK) {
        p.currentLine += 1
    }
}
func (p *Parser) currentTokenIs(t token.TokenType) bool {
    return p.currentToken.Type == t 
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return p.peekToken.Type == t
}

// ERRORS
func (p *Parser) addError(message string) {
    p.Errors = append(p.Errors, error.Error{
        Stage: "Parser",
        Line: p.currentLine,
        Message: message,
        File: p.File,
    })
}

// HELPERS
func (p *Parser) paragraphIsTerminated() bool {
    condition1 := p.currentTokenIs(token.DOUBLE_LINEBREAK)
    condition2 := p.currentTokenIs(token.LINEBREAK) && p.peekTokenIs(token.EOF)
    condition3 := p.currentTokenIs(token.EOF)

    return condition1 || condition2 || condition3
}

// PARSE
func (p *Parser) parseParagraph() *ast.Paragraph {
    paragraph := &ast.Paragraph{}

    for !p.paragraphIsTerminated() {
        expression := p.parseInline()
        if expression != nil {
            paragraph.Content = append(paragraph.Content, expression)
        }
    }
    if !p.currentTokenIs(token.EOF) {
        p.nextToken()
    }

    return paragraph
}

func (p *Parser) parseInline() ast.Inline {
    switch p.currentToken.Type {
    case token.LINEBREAK, token.DOUBLE_LINEBREAK:
        p.nextToken()
        return nil
    case token.TEXT:
        return p.parseText()
    case token.TAG:
        return p.parseTag()
    default:
        p.addError(fmt.Sprintf("Unexpected token: %s", token.TypeToString[p.currentToken.Type]))
        return nil
    }
}
func (p *Parser) parseText() *ast.Text {
    text := &ast.Text{Content: p.currentToken.Literal}
    p.nextToken()

    for p.currentTokenIs(token.LINEBREAK) && p.peekTokenIs(token.TEXT) {
        p.nextToken()
        text.Content += "\n" + p.currentToken.Literal
        p.nextToken()
    }

    return text
}
func (p *Parser) parseTag() *ast.Tag {
    tag := &ast.Tag{Name: strings.TrimPrefix(p.currentToken.Literal, "@"), LineNumber: p.currentLine}
    p.nextToken()

    if p.paragraphIsTerminated() {
        return tag
    }

    var inlineExpressions []ast.Inline
    if p.currentTokenIs(token.LBRACE) {
        p.nextToken()
        
        counter := 0
        for !p.currentTokenIs(token.RBRACE) {
            counter += 1
            if counter > 100 {
                p.addError("Right brace not found. Probably an unescaoed '#' char?")
                return nil
            }

            expression := p.parseInline()
            if expression != nil {
                inlineExpressions = append(inlineExpressions, expression)
            }
            
            if p.currentTokenIs(token.RBRACE) && p.peekTokenIs(token.LBRACE) {
                tag.Arguments = append(tag.Arguments, inlineExpressions)
                inlineExpressions = []ast.Inline{}
                p.nextToken()
                p.nextToken()
            }
        }
        tag.Arguments = append(tag.Arguments, inlineExpressions)

        p.nextToken()

        return tag
    }

    if p.currentTokenIs(token.LINEBREAK) { 
        p.nextToken()

        for !p.paragraphIsTerminated() {
            expression := p.parseInline()
            if expression != nil {
                inlineExpressions = append(inlineExpressions, expression)
            }
        }
        tag.Arguments = append(tag.Arguments, inlineExpressions)
         
        return tag
    }
    
    return tag
}
