package token

type TokenType int

const (
    ILLEGAL TokenType = iota
    EOF

    // Main types
    TAG
    TEXT
    CHAR

    // Characters
    LINEBREAK
    LBRACE
    RBRACE
    DASH

    // Tag types
    // Alignment
    JUSTIFY
    CENTER
    LEFT
    RIGHT

    // Headers
    H0
    H1
    H2
    H3

    // Text styles
    BOLD
    ITALIC
    UNDERLINE
    STRIKETHROUGH

    // Lists
    BULLETLIST
    NUMBERLIST

    // Links
    URL

    // Images
    IMAGE

    // Other
    BREAK
)

var TypeToString = map[TokenType]string{
    ILLEGAL: "ILLEGAL",
    EOF: "EOF",

    TAG: "TAG",
    TEXT: "TEXT",
    CHAR: "CHAR",

    LINEBREAK: "LINEBREAK",
    LBRACE: "LBRACE",
    RBRACE: "RBRACE",
    DASH: "DASH",

    JUSTIFY: "JUSTIFY",
    CENTER: "CENTER",
    LEFT: "LEFT",
    RIGHT: "RIGHT",

    H0: "H0",
    H1: "H1",
    H2: "H2",
    H3: "H3",

    BOLD: "BOLD",
    ITALIC: "ITALIC",
    UNDERLINE: "UNDERLINE",
    STRIKETHROUGH: "STRIKETHROUGH",

    BULLETLIST: "BULLETLIST",
    NUMBERLIST: "NUMBERLIST",

    URL: "URL",

    IMAGE: "IMAGE",

    BREAK: "BREAK",
}

type Token struct {
    Type TokenType
    Subtype TokenType
    Content string
}

var characters = map[byte]Token{
    0: {Type: CHAR, Subtype: EOF},
    '\n': {Type: CHAR, Subtype: LINEBREAK},
    '{': {Type: CHAR, Subtype: LBRACE},
    '}': {Type: CHAR, Subtype: RBRACE},
    '-': {Type: CHAR, Subtype: DASH},
}

var tags = map[string]Token{
    "justify": {Type: TAG, Subtype: JUSTIFY},
    "center": {Type: TAG, Subtype: CENTER},
    "left": {Type: TAG, Subtype: LEFT},
    "right": {Type: TAG, Subtype: RIGHT},

    "h0": {Type: TAG, Subtype: H0},
    "h1": {Type: TAG, Subtype: H1},
    "h2": {Type: TAG, Subtype: H2},
    "h3": {Type: TAG, Subtype: H3},

    "bold": {Type: TAG, Subtype: BOLD},
    "italic": {Type: TAG, Subtype: ITALIC},
    "underline": {Type: TAG, Subtype: UNDERLINE},
    "strikethrough": {Type: TAG, Subtype: STRIKETHROUGH},

    "bulletlist": {Type: TAG, Subtype: BULLETLIST},
    "numberlist": {Type: TAG, Subtype: NUMBERLIST},

    "url": {Type: TAG, Subtype: URL},

    "image": {Type: TAG, Subtype: IMAGE},

    "break": {Type: TAG, Subtype: BREAK},
}

func NewCharToken(char byte) Token {
    token, ok := characters[char]
    if ok {
        return token
    }
    return Token{Type: ILLEGAL}
}
func NewTagToken(tag string) Token {
    token, ok := tags[tag]
    if ok {
        return token
    }
    return Token{Type: ILLEGAL}
}
func NewTextToken(text string) Token {
    return Token{Type: TEXT, Subtype: TEXT, Content: text}
}
