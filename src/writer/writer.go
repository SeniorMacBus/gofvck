package writer

import (
	"fmt"
	"gofvck/token"
	"os"
)

func WriteTokens(tokens []token.Token) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(dir)
}
