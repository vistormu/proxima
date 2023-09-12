package tokenizer

import (
    "proxima/token"
)

type Tokenizer struct {
    input string
    position int
    peekPosition int
    char byte
}

// PUBLIC
func New(input string) *Tokenizer {
    t := &Tokenizer{input: input}
    t.readChar()
    return t
}

func (t *Tokenizer) GetToken() token.Token {
    if t.char == '\\' {
        t.readChar()
        text := string(t.char)
        t.readChar()
        return token.NewTextToken(text)
    }
    if t.char == '@' {
        return token.NewTagToken(t.readTag())
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
    start := t.position
    for t.char != ' ' && t.char != '\n' && t.char != 0 && t.char != '{' {
        t.readChar()
    }
    return t.input[start:t.position]
} 

// HELPERS
func isText(char byte) bool {
    _, ok := token.Characters[char]
    return !ok && char != '@' && char != '\\'
}
