package token

type TokenType string
const (
    ILLEGAL TokenType = "ILLEGAL"
    EOF = "EOF"

    TAG = "TAG"
    TEXT = "TEXT"

    LINEBREAK = "LINEBREAK"

    LBRACE = "LBRACE"
    RBRACE = "RBRACE"
    LANGLE = "LANGLE"
    RANGLE = "RANGLE"

    // ...
    HASHTAG = "HASHTAG"
    SPACE = "SPACE"
    BACKSLASH = "BACKSLASH"
)

type Token struct {
    Type TokenType
    Literal string
}

var Characters = map[rune]Token{
    0: {EOF, ""},
    '@': {TAG, "@"},
    '\n': {LINEBREAK, "\n"},
    '{': {LBRACE, "{"},
    '}': {RBRACE, "}"},
    '<': {LANGLE, "<"},
    '>': {RANGLE, ">"},
    '#': {HASHTAG, "#"},
    ' ': {SPACE, " "},
    '\\': {BACKSLASH, "\\"},
}

func New(literal any) Token {
    switch literal.(type) {
    case rune:
        token, ok := Characters[literal.(rune)]
        if !ok {
            return Token{ILLEGAL, string(literal.(rune))}
        }
        return token

    case string:
        return Token{TEXT, literal.(string)}

    default:
        return Token{ILLEGAL, ""}
    }
}
