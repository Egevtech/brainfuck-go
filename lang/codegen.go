package lang

import (
	"fmt"

	"github.com/egevtech/foreach"
)

func Codegen(tokens []any) string {
	var contents string

	stdfunctions := "vector_init, move_forward, move_backward, add_to_cell, sub_from_cell, print_cell, print_cell_num, ln, vector_get"

	nesting_level := 0

	contents += "default rel\nsection .data\nglobal main\n\nvec dq 0\ncell dd 0\n\nextern "
	contents += stdfunctions
	contents += "\n\nsection .text\nmain:\n\tpush rbp\n\tmov rbp, rsp\n\n\tcall vector_init\n\tmov [vec], rax\n\n"

	foreach.ForEach(tokens, func(index int, token any) {
		switch t := token.(type) {
		case ParAdd:
			contents += fmt.Sprintf("\n\tmov rdi, [vec]\n\tmov rsi, %d\n\tcall add_to_cell\n", t.Value)
		case ParSub:
			contents += fmt.Sprintf("\n\tmov rdi, [vec]\n\tmov rsi, %d\n\tcall sub_from_cell\n", t.Value)
		case ParMoveFor:
			contents += fmt.Sprintf("\n\tmov rdi, [vec]\n\tmov rsi, %d\n\tcall move_forward\n", t.Value)
		case ParMoveBack:
			contents += fmt.Sprintf("\n\tmov rdi, [vec]\n\tmov rsi, %d\n\tcall move_backward\n", t.Value)
		case ParDebug:
			contents += "\n\tmov rdi, [vec]\n\tcall print_cell_num\n"
		case ParPrint:
			contents += "\n\tmov rdi, [vec]\n\tcall print_cell\n"
		case ParLoopStart:
			nesting_level++
			contents += fmt.Sprintf("\n;LOOP%d-START\n\tloop%d_start:\n\n\tmov rdi, [vec]\n\tcall vector_get\n\tcmp rax, 0\n\tje loop%d_end\n", nesting_level, nesting_level, nesting_level)
		case ParLoopEnd:
			contents += fmt.Sprintf("\n\tjmp loop%d_start\n\tloop%d_end:\n", nesting_level, nesting_level)

			nesting_level--
			if nesting_level < 0 {
				panic("Unclosed loop")
			}

		default:
			panic("Unexpected ParCommand while generating code")
		}
	})

	contents += "mov rsp, rbp\n\tpop rbp\n\n\tcall ln\n\n\tmov rax, 0\n\tret"

	return contents
}
