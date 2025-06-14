package token

import "testing"

func TestLookupIdent(t *testing.T) {
	// table–driven tests for all keywords defined in the package
	tests := []struct {
		input    string
		expected TokenType
	}{
		{"fn", FUNCTION},
		{"let", LET},
		{"true", TRUE},
		{"false", FALSE},
		{"if", IF},
		{"else", ELSE},
		{"return", RETURN},
		// an identifier that is not a keyword should be IDENT
		{"foobar", IDENT},
		{"add", IDENT},
	}

	for i, tt := range tests {
		got := LookupIdent(tt.input)
		if got != tt.expected {
			t.Fatalf("tests[%d] - wrong token type for %q. expected=%q, got=%q", i, tt.input, tt.expected, got)
		}
	}
}
