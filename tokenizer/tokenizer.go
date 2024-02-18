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
    if t.char == '#' {
        t.skipComment()
    }
    if t.char == '\\' {
        t.readChar()
        switch t.char {
        case 'n':
            t.readChar()
            return token.NewTextToken("<br>")
        case '<':
            t.readChar()
            return token.NewTextToken("&lt;")
        case '>':
            t.readChar()
            return token.NewTextToken("&gt;")
        }
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
    if t.char == '\n' && t.peekChar() == '\n' {
        t.readChar()
        t.readChar()
        return token.NewTwoCharToken('\n', '\n')
    }

    token := token.NewCharToken(t.char)
    t.readChar()

    return token
}

// PRIVATE
func (t *Tokenizer) peekChar() byte {
    if t.peekPosition >= len(t.input) {
        return 0
    }
    return t.input[t.peekPosition]
}
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
func (t *Tokenizer) skipComment() {
    for t.char != '\n' && t.char != 0 {
        t.readChar()
    }
}

// HELPERS
func isText(char byte) bool {
    _, ok := token.Characters[char]
    return !ok && char != '@' && char != '\\' && char != '#'
}
