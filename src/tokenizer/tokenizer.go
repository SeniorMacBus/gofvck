package tokenizer

import "gofvck/token"

type ParseError struct{}
type Stack struct {
	data []uint64
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

func createToken(token_type byte, value uint64) token.Token {
	token := token.Token{Type: token.Token_map[token_type], Value: value}

	return token
}

func Tokenize(code []byte) ([]token.Token, map[uint64]uint64, map[uint64]uint64, error) {
	tokens := make([]token.Token, 0)
	var prev_char byte = code[0]
	var counter uint64 = 0
	for _, t := range code {
		if t == prev_char {
			counter++
		} else {
			tokens = append(tokens, createToken(prev_char, counter))
			counter = 1
		}
		prev_char = t
	}

	end_token := token.Token{Type: token.CODE_END, Value: 1}
	tokens = append(tokens, end_token)
	loop_starts, loop_ends, e := parseTokensForLoops(tokens)
	if e != nil {
		return nil, nil, nil, e
	}

	return tokens, loop_starts, loop_ends, nil
}

func parseTokensForLoops(tokens []token.Token) (map[uint64]uint64, map[uint64]uint64, error) {
	s := Stack{data: make([]uint64, 0)}
	s_indices := make(map[uint64]uint64, 0)
	e_indices := make(map[uint64]uint64, 0)
	for i, t := range tokens {
		switch t.Type {
		case token.LOOP_START:
			s.append(uint64(i))
		case token.LOOP_END:
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
