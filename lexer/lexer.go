package lexer

// Lexer は構文解析器を表す。
type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  // これから読み込む位置
	ch           byte // 現在検査中の文字
}

// New は新しい構文解析器を返す。
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	// 初期化
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.reaition
	l.readPosition += 1
}
