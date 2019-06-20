package eval

import (
	"strings"
	"testing"
	"text/scanner"
)

/*
lex.nextとlex.textを使って文字を1文字づつ取得することができる
*/
func TestLexText(t *testing.T) {
	tests := []struct {
		input    string
		nexNum   int
		expected string
	}{
		{"a+b", 1, "a"},
		{"a+b", 2, "+"},
		{"a+b", 3, "b"},
	}

	for _, test := range tests {
		lexText(t, test.input, test.nexNum, test.expected)
	}
}

/*
nextNum = nextメソッドを実行する回数
*/
func lexText(t *testing.T, input string, nextNum int, expected string) {
	t.Helper()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats

	for i := 0; i < nextNum; i++ {
		lex.next()
	}

	actual := lex.text()
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
