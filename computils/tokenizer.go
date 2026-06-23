package computils

import (
	"fmt"
	"unicode"
)

type Token int

const (
	TOKEN_NEXT Token = iota
	TOKEN_PREV

	TOKEN_ADD
	TOKEN_SUB

	TOKEN_PRINT
	TOKEN_DEBUG

	TOKEN_LOOP_START
	TOKEN_LOOP_END
)

func (t Token) String() string {
	switch t {
	case TOKEN_ADD:
		return "TOKEN_ADD"
	case TOKEN_SUB:
		return "TOKEN_SUB"
	case TOKEN_NEXT:
		return "TOKEN_NEXT"
	case TOKEN_PREV:
		return "TOKEN_PREV"
	case TOKEN_PRINT:
		return "TOKEN_PRINT"
	case TOKEN_DEBUG:
		return "TOKEN_DEBUG"
	case TOKEN_LOOP_START:
		return "TOKEN_LOOP_START"
	case TOKEN_LOOP_END:
		return "TOKEN_LOOP_END"
	}

	return ""
}

func ForEach[T any](i []T, lb func(int, T)) {
	for index, curr := range i {
		lb(index, curr)
	}
}

func tok_append(vec *[]Token, token Token) {
	*vec = append(*vec, token)
}

func Tokenize(input string) ([]Token, error) {
	tokens := []Token{}

	for _, char := range input {
		if unicode.IsSpace(char) {
			continue
		}

		switch char {
		case '.':
			tok_append(&tokens, TOKEN_PRINT)
		case 'D':
			tok_append(&tokens, TOKEN_DEBUG)

		case '+':
			tok_append(&tokens, TOKEN_ADD)
		case '-':
			tok_append(&tokens, TOKEN_SUB)

		case '>':
			tok_append(&tokens, TOKEN_NEXT)
		case '<':
			tok_append(&tokens, TOKEN_PREV)

		case '[':
			tok_append(&tokens, TOKEN_LOOP_START)
		case ']':
			tok_append(&tokens, TOKEN_LOOP_END)

		default:
			return []Token{}, fmt.Errorf("Unexpected token: %c", char)
		}
	}

	return tokens, nil
}
