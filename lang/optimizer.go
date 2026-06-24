package lang

import (
	"fmt"
)

type ParCommand struct {
	Value int
}

type ParAdd ParCommand
type ParSub ParCommand

type ParMoveFor ParCommand
type ParMoveBack ParCommand

func (t ParAdd) String() string {
	return fmt.Sprintf("ParAdd{%d}", t.Value)
}

func (t ParSub) String() string {
	return fmt.Sprintf("ParSub{%d}", t.Value)
}

func (t ParMoveFor) String() string {
	return fmt.Sprintf("ParMoveFor{%d}", t.Value)
}

func (t ParMoveBack) String() string {
	return fmt.Sprintf("ParMoveBack{%d}", t.Value)
}

func Optimize(tokens []Token) []any {
	var res []any // ParCommand
	current := 0

	current_token := func() Token { return tokens[current] }

	for current < len(tokens) {
		switch current_token() {
		case TOKEN_ADD:
			counter := 0
			for current_token() == TOKEN_ADD {
				current++
				counter++
				if current >= len(tokens) {
					break
				}
			}
			res = append(res, ParAdd{Value: counter})
		case TOKEN_SUB:
			counter := 0
			for current_token() == TOKEN_SUB {
				current++
				counter++
				if current >= len(tokens) {
					break
				}
			}
			res = append(res, ParSub{Value: counter})
		case TOKEN_NEXT:
			counter := 0
			for current_token() == TOKEN_NEXT {
				current++
				counter++
				if current >= len(tokens) {
					break
				}
			}
			res = append(res, ParMoveFor{Value: counter})
		case TOKEN_PREV:
			counter := 0
			for current_token() == TOKEN_PREV {
				current++
				counter++
				if current >= len(tokens) {
					break
				}
			}
			res = append(res, ParMoveBack{Value: counter})
		}
	}

	return res
}
