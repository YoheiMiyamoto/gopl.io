package eval

import (
	"strings"
	"testing"
	"text/scanner"
)

/*
初めのlexを取得
*/
func TestParsePrimary(t *testing.T) {
	tests := []struct {
		in       string
		expected Expr
	}{
		{"a+b", Var("a")},
		{"(a+b)*2", binary{[]rune("+")[0], Var("a"), Var("b")}},
		{"((a+b)*3)*2", binary{[]rune("*")[0], binary{[]rune("+")[0], Var("a"), Var("b")}, literal(3)}},
	}
	for _, test := range tests {
		testParsePrimary(t, test.in, test.expected)
	}
}

func testParsePrimary(t *testing.T, in string, expected Expr) {
	t.Helper()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(in))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next()

	actual := parsePrimary(lex)
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
