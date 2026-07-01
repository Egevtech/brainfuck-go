package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alexflint/go-arg"
	"github.com/egevtech/brainfuck/lang"
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

	tokens, err := lang.Tokenize(string(content))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Tokenizer failed: %s\n", err)
		return
	}

	optTokens := lang.Optimize(tokens)

	contents := lang.Codegen(optTokens)

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
		buildStdlib()
		link()

		err := os.Remove("out.o")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove object file: %s\n", err)
		}
	}
}

func buildStdlib() {
	var cs_command *exec.Cmd

	cs_command = exec.Command("bash", "./build.sh")

	cs_command.Dir = "external/bfstdlib"
	cs_command.Stderr = os.Stderr
	cs_command.Stdout = os.Stdout

	if err := cs_command.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build stdlib: %s\n", err)
		panic("Build failed")
	}
}

func assemble() {
	var asm_command *exec.Cmd

	asm_command = exec.Command("nasm", "-felf64", "out.s", "-o", "out.o")

	asm_command.Stderr = os.Stderr
	asm_command.Stdout = os.Stdout

	if err := asm_command.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run assemler: %s\n", err)
		panic("Build failed")
	}
}

func link() {
	var ld_command *exec.Cmd

	ld_command = exec.Command("clang", "out.o", "external/bfstdlib/stdlib.a")

	ld_command.Stderr = os.Stderr
	ld_command.Stdout = os.Stdout
	if err := ld_command.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run ld: %s\n", err)
		panic("Build failed")
	}
}
