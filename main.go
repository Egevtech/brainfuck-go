package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alexflint/go-arg"
	"github.com/egevtech/brainfuck/computils"
)

var args struct {
	Input  string `arg:"positional,required" help:"File to compile"`
	Output string `arg:"-o" help:"Output file" default:"a.out"`

	CompileOnly     bool `arg:"-S" help:"Compile only, do not assembly or link"`
	CompileAssembly bool `arg:"-c" help:"Compile and assembly, but do not link"`
}

func main() {
	arg.MustParse(&args)

	content, err := os.ReadFile(args.Input)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		return
	}

	tokens, err := computils.Tokenize(string(content))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Tokenizer failed: %s\n", err)
		return
	}

	contents := computils.Codegen(tokens)

	err = os.WriteFile("./out.s", []byte(contents), 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write asm file: %s\n", err)
		return
	}

	if args.CompileOnly {
		return
	} else {
		assemble()

		err := os.Remove("out.s")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove assembler file: %s\n", err)
		}
	}

	if args.CompileAssembly {
		return
	} else {
		link()

		err := os.Remove("out.o")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove object file: %s\n", err)
		}
	}
}

func assemble() {
	nasm_command := exec.Command("nasm", "-felf64", "out.s", "-o", "out.o")
	nasm_command.Stderr = os.Stderr
	nasm_command.Stdout = os.Stdout
	if err := nasm_command.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run nasm: %s\n", err)
		return
	}
}

func link() {
	ld_command := exec.Command(
		"ld", "out.o", "bfstd.a",
		"-lc", "-dynamic-linker",
		"/lib64/ld-linux-x86-64.so.2",
		"-o", args.Output,
	)

	ld_command.Stderr = os.Stderr
	ld_command.Stdout = os.Stdout
	if err := ld_command.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run ld: %s\n", err)
		return
	}
}
