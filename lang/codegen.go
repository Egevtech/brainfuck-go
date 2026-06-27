package lang

import (
	"fmt"

	"github.com/egevtech/brainfuck/util"
)

func Codegen(tokens []Token) string {
	var contents string

	stdfunctions := "vector_init, read_cell, next_cell, prev_cell, add_cell, sub_cell, print_cell, print_cell_num"

	nesting_level := 0

	contents += "section .data\nglobal _start\n\nvec dq 0\ncell dd 0\n\nextern "
	contents += stdfunctions
	contents += "\n\nsection .text\n_start:\n\tcall vector_init\n\tmov [vec], rax\n\n"

	util.ForEach(tokens, func(index int, token Token) {
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

		case TOKEN_LOOP_START:
			nesting_level++
			contents += fmt.Sprintf("\n;----LOOP%d-START----\n\tloop%d:\n", nesting_level, nesting_level)
			contents += "mov rdi, [vec]\n\tmov rsi, cell\n"
			contents += fmt.Sprintf("\tcall read_cell\n\tcmp rax, 0\n\tje loop%d_end\n\n", nesting_level)
			return
		case TOKEN_LOOP_END:
			contents += fmt.Sprintf("\tjmp loop%d\n\tloop%d_end:\n;----LOOP%d-END----\n\n", nesting_level, nesting_level, nesting_level)
			nesting_level--
			return
		}

		contents += fmt.Sprintf("\tmov rdi, [vec]\n\tmov rsi, cell\n\tcall %s\n\n", stdcall)
	})

	contents += "\tmov rax, 60\n\tmov rdi, 0\n\tsyscall\n"

	return contents
}
