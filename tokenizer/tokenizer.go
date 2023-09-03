package tokenizer

import (
    "proxima/token"
)

type Tokenizer struct {
    input string
    position int
    peekPosition int
    char byte
    skip bool
}

// PUBLIC
func New(input string) *Tokenizer {
    t := &Tokenizer{input: input}
    t.readChar()
    return t
}

func (t *Tokenizer) GetToken() token.Token {
    if t.skip { t.skipWhitespace() }
    if t.char == '}' { t.skip = false } else { t.skip = true }

    if t.char == '#' { t.skipComment() }

    if t.char == '@' {
        tag := t.readTag()
        return token.NewTagToken(tag)
    }
    if isText(t.char) {
        return token.NewTextToken(t.readText())
    }

    token := token.NewCharToken(t.char)
    t.readChar()

    return token
}

// PRIVATE
func (t *Tokenizer) readChar() {
    if t.peekPosition >= len(t.input) {
        t.char = 0
    } else {
        t.char = t.input[t.peekPosition]
    }
    t.position = t.peekPosition
    t.peekPosition += 1
}
func (t *Tokenizer) readText() string {
    start := t.position
    for isText(t.char) {
        t.readChar()
    }
    return t.input[start:t.position]
}
func (t *Tokenizer) readTag() string {
    t.readChar() // skip @
    start := t.position
    for t.char != ' ' && t.char != '\n' && t.char != 0 && t.char != '{' {
        t.readChar()
    }
    return t.input[start:t.position]
} 
func (t *Tokenizer) skipWhitespace() {
    for t.char == ' ' || t.char == '\t' {
        t.readChar()
    }
}
func (t *Tokenizer) skipComment() {
    for t.char != '\n' {
        t.readChar()
    }
}

// HELPERS
func isText(char byte) bool {
    return char != 0 && char != '\n' && char != '@' && char != '{' && char != '}'
}
