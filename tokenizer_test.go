package main

import "testing"

func TestTokenizer(t *testing.T) {
	str := "+-<>.D[]"
	expected := []Token{TOKEN_ADD, TOKEN_SUB, TOKEN_PREV, TOKEN_NEXT, TOKEN_PRINT, TOKEN_DEBUG, TOKEN_LOOP_START, TOKEN_LOOP_END}
	tokens, err := Tokenize(str)

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
