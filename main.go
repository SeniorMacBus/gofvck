package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ParseError struct{}
type Stack struct {
	data []uint64
}
type Token int64

const (
	MOVE_R = iota
	MOVE_L
	INC
	DEC
	PRINT
	INPUT
	LOOP_START
	LOOP_END
	CODE_END
)

func (t Token) move_r(p *uint64, size uint64) {
	if *p == size-1 {
		*p = 0
	} else {
		*p++
	}
}

func (t Token) move_l(p *uint64, size uint64) {
	if *p == 0 {
		*p = size - 1
	} else {
		*p--
	}
}

func (t Token) inc(v *uint8) {
	if *v == 255 {
		*v = 0
	} else {
		*v++
	}
}

func (t Token) dec(v *uint8) {
	if *v == 0 {
		*v = 255
	} else {
		*v--
	}
}

func (t Token) print_token(c byte) {
	fmt.Printf("%c", c)
}

func (t Token) jump_if_zero(v uint8, i *int, d uint64) {
	if v == 0 {
		*i = int(d)
	} else {
		*i += 1
	}
}

func (t Token) jump_if_non_zero(v uint8, i *int, d uint64) {
	if v != 0 {
		*i = int(d)
	} else {
		*i += 1
	}
}

func (e *ParseError) Error() string {
	return "The loops in the code are wrong!"
}

func (s *Stack) append(val uint64) {
	s.data = append(s.data, val)
}

func (s *Stack) pop() uint64 {
	var elm uint64
	elm, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return elm
}

func (s Stack) isempty() bool {
	return len(s.data) == 0
}

func ReadSource(file_name string) ([]byte, error) {
	f, err := os.Open(file_name)
	if err != nil {
		return make([]byte, 0), err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	content, read_err := io.ReadAll(reader)
	if read_err != nil {
		return make([]byte, 0), read_err
	}

	return content, nil
}

func Tokenize(code []byte) []Token {
	tokens := make([]Token, len(code))
	for i, t := range code {
		switch t {
		case byte('>'):
			tokens[i] = MOVE_R
		case byte('<'):
			tokens[i] = MOVE_L
		case byte('+'):
			tokens[i] = INC
		case byte('-'):
			tokens[i] = DEC
		case byte(','):
			tokens[i] = INPUT
		case byte('.'):
			tokens[i] = PRINT
		case byte('['):
			tokens[i] = LOOP_START
		case byte(']'):
			tokens[i] = LOOP_END
		default:
			tokens[i] = CODE_END
		}
	}

	return tokens
}

func ParseTokens(tokens []Token) (map[uint64]uint64, map[uint64]uint64, error) {
	s := Stack{data: make([]uint64, 0)}
	s_indices := make(map[uint64]uint64, 0)
	e_indices := make(map[uint64]uint64, 0)
	for i, t := range tokens {
		switch t {
		case LOOP_START:
			s.append(uint64(i))
		case LOOP_END:
			if s.isempty() {
				return nil, nil, &ParseError{}
			}
			s_idx := s.pop()
			s_indices[s_idx] = uint64(i)
			e_indices[uint64(i)] = s_idx
		}
	}

	if !s.isempty() {
		return nil, nil, &ParseError{}
	}

	return s_indices, e_indices, nil
}

func Execute(tokens []Token, buffer []uint8, loop_start map[uint64]uint64, loop_end map[uint64]uint64) {
	var p uint64 = 0
	var i int = 0
	size := len(buffer)
	for tokens[i] != CODE_END {
		t := tokens[i]
		switch t {
		case MOVE_R:
			t.move_r(&p, uint64(size))
			i++
		case MOVE_L:
			t.move_l(&p, uint64(size))
			i++
		case INC:
			t.inc(&buffer[p])
			i++
		case DEC:
			t.dec(&buffer[p])
			i++
		case PRINT:
			t.print_token(buffer[p])
			i++
		case LOOP_START:
			t.jump_if_zero(buffer[p], &i, loop_start[uint64(i)]+1)
		case LOOP_END:
			t.jump_if_non_zero(buffer[p], &i, loop_end[uint64(i)]+1)
		}
	}
}

func main() {
	// Setting up a clean project
	buffer := make([]uint8, 5)
	//var p uint64 = 0

	// Reading the source code
	file_name := os.Args[1]
	content, err := ReadSource(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Tokenizing
	tokens := Tokenize(content)
	loop_start, loop_end, _ := ParseTokens(tokens)
	Execute(tokens, buffer, loop_start, loop_end)

}
