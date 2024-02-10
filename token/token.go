package token

type TokenType int

const (
    ILLEGAL TokenType = iota
    EOF

    TAG
    TEXT

    LINEBREAK
    DOUBLE_LINEBREAK
    LBRACE
    RBRACE
)

var TypeToString = map[TokenType]string{
    ILLEGAL: "ILLEGAL",
    EOF: "EOF",

    TAG: "TAG",
    TEXT: "TEXT",

    LINEBREAK: "LINEBREAK",
    DOUBLE_LINEBREAK: "DOUBLE_LINEBREAK",
    LBRACE: "LBRACE",
    RBRACE: "RBRACE",
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
}

func NewCharToken(char byte) Token {
    token, ok := Characters[char]
    if !ok {
        return Token{ILLEGAL, string(char)}
    }
    return token
}
func NewTwoCharToken(char1, char2 byte) Token {
    return Token{DOUBLE_LINEBREAK, "\n\n"} // WIP
}
func NewTagToken(tag string) Token {
    return Token{TAG, tag}
}
func NewTextToken(text string) Token {
    return Token{TEXT, text}
}
