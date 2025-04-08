package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadSource(file_name string) ([]byte, error) {
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	content, read_err := io.ReadAll(reader)
	if read_err != nil {
		return nil, read_err
	}

	return content, nil
}
