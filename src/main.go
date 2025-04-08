package main

import (
	"fmt"
	"gofvck/executor"
	"gofvck/parser"
	"gofvck/tokenizer"
	"os"
)

func main() {
	// Setting up a clean project
	buffer := make([]uint8, 30000)
	//var p uint64 = 0

	// Reading the source code
	file_name := os.Args[1]
	content, err := parser.ReadSource(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Tokenizing
	tokens, loop_start, loop_end, token_err := tokenizer.Tokenize(content)
	if token_err != nil {
		fmt.Println(token_err)
		return
	}

	// Executing operations
	executor.Execute(tokens, buffer, loop_start, loop_end)

}
