package tests

import (
	"fmt"
	"testing"

	"github.com/egevtech/brainfuck/lang"
)

func TestOptimizer(t *testing.T) {
	src, err := lang.Tokenize("+++---<<<>>>-")
	if err != nil {
		t.Error("Failed to tokenize test source")
		return
	}

	expected := []any{lang.ParAdd{3}, lang.ParSub{3}, lang.ParMoveBack{3}, lang.ParMoveFor{3}, lang.ParSub{1}}
	got := lang.Optimize(src)

	if len(expected) != len(got) {
		t.Errorf("Lengths not match: expected %d != got %d\n", len(expected), len(got))
	}

	for index := range expected {
		if expected[index] != got[index] {
			t.Errorf("Match failed: %s != %s at index %d", expected[index], got[index], index)
		} else {
			fmt.Printf("%s == %s\n", expected[index], got[index])
		}
	}

}
