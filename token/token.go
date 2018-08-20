package token

// 定数
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234567890

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Type はトークンの種類を表す。
type Type string

// Token は字句解析器が出力するトークン。
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent は引数の IDENT がキーワードでないか判定する。
// もし、キーワードであれば、キーワードの Type を返す。
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
