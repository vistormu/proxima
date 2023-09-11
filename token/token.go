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
    BACKSLASH
    HASH
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
    BACKSLASH: "BACKSLASH",
    HASH: "HASH",
}

type Token struct {
    Type TokenType
    Literal string
}

var Characters = map[byte]Token{
    0: {EOF, ""},
    '\n': {LINEBREAK, "\n"},
    '{': {LBRACE, "{"},
    '}': {RBRACE, "}"},
    '-': {DASH, "-"},
    '\\': {BACKSLASH, "\\"},
    '#': {HASH, "#"},
}

func NewCharToken(char byte) Token {
    token, ok := Characters[char]
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
