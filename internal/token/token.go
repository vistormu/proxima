package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	TAG  TokenType = "TAG"
	TEXT TokenType = "TEXT"

	LINEBREAK TokenType = "LINEBREAK"

	LBRACE  TokenType = "LBRACE"
	RBRACE  TokenType = "RBRACE"
	LANGLE  TokenType = "LANGLE"
	RANGLE  TokenType = "RANGLE"
	HASHTAG TokenType = "HASHTAG"
)

type Token struct {
	Type    TokenType
	Literal string
}

var Characters = map[rune]Token{
	0:    {EOF, ""},
	'@':  {TAG, "@"},
	'\n': {LINEBREAK, "\\n"},
	'{':  {LBRACE, "{"},
	'}':  {RBRACE, "}"},
	'<':  {LANGLE, "<"},
	'>':  {RANGLE, ">"},
	'#':  {HASHTAG, "#"},
}

func New(literal any) Token {
	switch literal := literal.(type) {
	case rune:
		token, ok := Characters[literal]
		if !ok {
			return Token{ILLEGAL, string(literal)}
		}
		return token

	case string:
		return Token{TEXT, literal}

	default:
		return Token{ILLEGAL, ""}
	}
}
