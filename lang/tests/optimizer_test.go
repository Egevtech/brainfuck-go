package tests

import (
	"fmt"
	"testing"

	"github.com/egevtech/brainfuck/lang"
)

func TestOptimizerShortQuery(t *testing.T) {
	src, err := lang.Tokenize("+-<>[].D")
	if err != nil {
		t.Errorf("Failed to tokenize test source: %s\n", err)
		return
	}

	expected := []any{lang.ParAdd{1}, lang.ParSub{1}, lang.ParMoveBack{1}, lang.ParMoveFor{1}, lang.ParLoopStart{0}, lang.ParLoopEnd{0}, lang.ParPrint{0}, lang.ParDebug{0}}
	got := lang.Optimize(src)

	if len(expected) != len(got) {
		t.Errorf("Lengths not match: expected %d, got %d\n", len(expected), len(got))
	}

	for index := range expected {
		if expected[index] != got[index] {
			t.Errorf("Match failed: %s != %s at %d", expected[index], got[index], index)
		} else {
			fmt.Printf("%s == %s\n", expected[index], got[index])
		}
	}
}
func TestOptimizerLongQuery(t *testing.T) {
	src, err := lang.Tokenize("+++---<<<>>>")
	if err != nil {
		t.Errorf("Failed to tokenize test source: %s\n", err)
		return
	}

	expected := []any{lang.ParAdd{3}, lang.ParSub{3}, lang.ParMoveBack{3}, lang.ParMoveFor{3}}
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
