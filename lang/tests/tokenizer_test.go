package tests

import (
	"testing"

	"github.com/egevtech/brainfuck/lang"
)

func TestTokenizer(t *testing.T) {
	str := "+-<>.D[]"
	expected := []lang.Token{lang.TOKEN_ADD, lang.TOKEN_SUB, lang.TOKEN_PREV, lang.TOKEN_NEXT, lang.TOKEN_PRINT, lang.TOKEN_DEBUG, lang.TOKEN_LOOP_START, lang.TOKEN_LOOP_END}
	tokens, err := lang.Tokenize(str)

	if err != nil {
		t.Fatalf("Tokenizer failed: %s", err)
	}

	if len(expected) != len(tokens) {
		t.Errorf("Lenghts not equals, get %d, expected %d", len(expected), len(tokens))
	}

	for index := range expected {
		if expected[index] != tokens[index] {
			t.Errorf("Error on symbol %d, expected %d, got %d", index, expected[index], tokens[index])
		}
	}
}
