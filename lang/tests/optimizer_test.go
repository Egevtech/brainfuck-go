package tests

import (
	"testing"

	"github.com/egevtech/brainfuck/lang"
)

func TestOptimizer(t *testing.T) {
	src := []lang.Token{
		lang.TOKEN_ADD, lang.TOKEN_ADD, lang.TOKEN_ADD,
		lang.TOKEN_SUB, lang.TOKEN_SUB, lang.TOKEN_SUB,
		lang.TOKEN_PREV, lang.TOKEN_PREV, lang.TOKEN_PREV,
		lang.TOKEN_NEXT, lang.TOKEN_NEXT, lang.TOKEN_NEXT,
	}

	expected := []any{lang.ParAdd{3}, lang.ParSub{3}, lang.ParMoveBack{3}, lang.ParMoveFor{3}}
	got := lang.Optimize(src)

	if len(expected) != len(got) {
		t.Errorf("Lengths not match: expected %d != got %d\n", len(expected), len(got))
	}

	for index := range expected {
		if expected[index] != got[index] {
			t.Errorf("Match failed: %s != %s at index %d", expected[index], got[index], index)
		}
	}
}
