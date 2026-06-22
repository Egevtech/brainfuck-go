package main

import "fmt"

func Codegen(tokens []Token) string {
	var contents string

	stdfunctions := "vector_init, next_cell, prev_cell, add_cell, sub_cell, print_cell, print_cell_num"

	contents += "section .data\nglobal _start\n\nvec dq 0\ncell dd 0\n\nextern "
	contents += stdfunctions
	contents += "\n\nsection .text\n_start:\n\tcall vector_init\n\tmov [vec], rax\n\n"

	ForEach(tokens, func(index int, token Token) {
		var stdcall string

		switch token {
		case TOKEN_ADD:
			stdcall = "add_cell"
		case TOKEN_SUB:
			stdcall = "sub_cell"
		case TOKEN_NEXT:
			stdcall = "next_cell"
		case TOKEN_PREV:
			stdcall = "prev_cell"
		case TOKEN_PRINT:
			stdcall = "print_cell"
		case TOKEN_DEBUG:
			stdcall = "print_cell_num"
		}

		contents += fmt.Sprintf("\tmov rdi, [vec]\n\tmov rsi, cell\n\tcall %s\n\n", stdcall)
	})

	contents += "\tmov rax, 60\n\tmov rdi, 0\n\tsyscall"

	return contents
}
