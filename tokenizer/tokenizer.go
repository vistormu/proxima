package tokenizer

import (
    "proxima/token"
)

type Tokenizer struct {
    input []rune
    inputLen int
    position int
}

func New(input []rune) *Tokenizer {
    return &Tokenizer{
        input,
        len(input),
        0,
    }
}

func (t *Tokenizer) Tokenize() []token.Token {
    tokens := []token.Token{}

    for tok := t.token(); tok.Type != token.EOF; tok = t.token() {
        tokens = append(tokens, tok)
    }

    return tokens
}

func (t *Tokenizer) token() token.Token {
    // read char
    char, peekChar := t.readChar()

    // skip whitespace
    for isWhitespace(char) {
        char, peekChar = t.readChar()
    }

    // text
    if isText(char) {
        literal := ""
        for isText(char) {
            literal += string(char)

            if !isText(peekChar) {
                break
            }

            char, peekChar = t.readChar()
        }
        return token.New(literal)
    }

    // rest
    return token.New(char)
}

func (t *Tokenizer) readChar() (rune, rune) {
    if t.position >= t.inputLen {
        return 0, 0
    }

    currentChar := t.input[t.position]
    peekChar := rune(0)
    if t.position+1 < t.inputLen {
        peekChar = t.input[t.position+1]
    }

    t.position++

    return currentChar, peekChar
}

func isWhitespace(char rune) bool {
    return char == '\t' || char == '\r'
}
func isText(char rune) bool {
    _, ok := token.Characters[char]
    return (!ok || char == ' ') && !isWhitespace(char)
}
