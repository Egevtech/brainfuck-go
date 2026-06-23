package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alexflint/go-arg"
)

var args struct {
	Input  string `arg:"positional,required" help:"File to compile"`
	Output string `arg:"-o" help:"Output file"`
}

func main() {
	arg.MustParse(&args)

	if args.Output == "" {
		args.Output = "a.out"
	}

	content, err := os.ReadFile(args.Input)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		return
	}

	tokens, err := Tokenize(string(content))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Tokenizer failed: %s\n", err)
		return
	}

	contents := Codegen(tokens)

	err = os.WriteFile("./out.s", []byte(contents), 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write asm file: %s\n", err)
		return
	}

	assemble()
	link()
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
