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

    // skip comments
    if char == '#' {
        for char != '\n' {
            char, peekChar = t.readChar()
        }
    }

    // special characters
    tok, ok := token.Characters[char]
    if ok {
        // keep white space after }
        if char == '}' && peekChar == ' ' {
            return tok
        }

        // skip white space after special characters
        for isWhiteSpace(peekChar) {
            char, peekChar = t.readChar()
        }

        return tok
    }

    // text and escape characters
    if isText(char) || char == '\\' {
        literal := ""
        for isText(char) || char == '\\' {
            // escape character
            if char == '\\' {
                literal += string(peekChar)
                char, peekChar = t.readChar()
                char, peekChar = t.readChar()

                continue
            }

            // text
            literal += string(char)

            if !isText(peekChar) && peekChar != '\\' {
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

func isWhiteSpace(char rune) bool {
    return char == '\t' || char == '\r' || char == ' '
}
func isText(char rune) bool {
    _, ok := token.Characters[char]
    return !ok && !isWhiteSpace(char) || char == ' '
}
