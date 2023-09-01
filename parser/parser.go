package parser

import (
    "fmt"
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
}

// PUBLIC
func New(input string) *Parser {
    p := &Parser{tokenizer: tokenizer.New(input), currentLine: 1}
    p.nextToken()
    p.nextToken()
    return p
}
func (p *Parser) Parse() *ast.Document {
    document := &ast.Document{}

    for !p.currentTokenIs(token.EOF) {
        paragraph := p.parseParagraph()
        document.Paragraphs = append(document.Paragraphs, paragraph)
        p.nextToken()
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
func (p *Parser) expectPeekTokenToBe(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        p.addError(fmt.Sprintf("expected next token to be %d, got %d", t, p.peekToken.Type))
        return false
    }
}

// ERRORS
func (p *Parser) addError(message string) {
    p.Errors = append(p.Errors, error.Error{
        Stage: "Parser",
        Line: p.currentLine,
        Message: message,
    })
}

// HELPERS
func (p *Parser) paragraphIsTerminated() bool {
    return p.currentTokenIs(token.LINEBREAK) && p.peekTokenIs(token.LINEBREAK) || 
    p.currentTokenIs(token.LINEBREAK) && p.peekTokenIs(token.EOF) ||
    p.currentTokenIs(token.EOF)
}

// PARSE
func (p *Parser) parseParagraph() *ast.Paragraph {
    paragraph := &ast.Paragraph{}

    for !p.paragraphIsTerminated() {
        paragraph.Content = append(paragraph.Content, p.parseInline())
        p.nextToken()
    }

    return paragraph
}

func (p *Parser) parseInline() ast.Inline {
    if p.currentTokenIs(token.LINEBREAK) {
        p.nextToken()
    }
    if p.currentTokenIs(token.TEXT) {
        return p.parseText()
    }
    if p.currentTokenIs(token.TAG) {
        return p.parseTag()
    }
    
    p.addError(fmt.Sprintf("Unexpected token: %s", token.TypeToString[p.currentToken.Type]))
    return nil
}
func (p *Parser) parseText() *ast.Text {
    return &ast.Text{Content: p.currentToken.Literal}
}
func (p *Parser) parseTag() *ast.Tag {
    tag := &ast.Tag{Name: p.currentToken.Literal}

    if !p.expectPeekTokenToBe(token.LBRACE) { return nil }
    p.nextToken()

    for !p.currentTokenIs(token.RBRACE) {
        tag.Content = append(tag.Content, p.parseInline())
        p.nextToken()
    }

    return tag
}
