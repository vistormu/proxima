package token

type TokenType int

const (
    ILLEGAL TokenType = iota
    EOF

    TAG
    TEXT
    LINEBREAK
    LBRACE
    RBRACE
    DASH
)

var TypeToString = map[TokenType]string{
    ILLEGAL: "ILLEGAL",
    EOF: "EOF",

    TAG: "TAG",
    TEXT: "TEXT",
    LINEBREAK: "LINEBREAK",
    LBRACE: "LBRACE",
    RBRACE: "RBRACE",
    DASH: "DASH",
}

type Token struct {
    Type TokenType
    Literal string
}

var characters = map[byte]Token{
    0: {EOF, ""},
    '\n': {LINEBREAK, "\n"},
    '{': {LBRACE, "{"},
    '}': {RBRACE, "}"},
    '-': {DASH, "-"},
}

func NewCharToken(char byte) Token {
    token, ok := characters[char]
    if !ok {
        return Token{ILLEGAL, string(char)}
    }
    return token
}
func NewTagToken(tag string) Token {
    return Token{TAG, tag}
}
func NewTextToken(text string) Token {
    return Token{TEXT, text}
}
