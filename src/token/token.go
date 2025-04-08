package token

import "fmt"

type TokenType int

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

type Token struct {
	Type  TokenType
	Value uint64
}

var Token_map map[byte]TokenType = map[byte]TokenType{
	'>': MOVE_R,
	'<': MOVE_L,
	'+': INC,
	'-': DEC,
	'.': PRINT,
	',': INPUT,
	'[': LOOP_START,
	']': LOOP_END,
}

// Operations connected to tokens
func (t Token) Move_r(p *uint64, size uint64) {
	// moves the value pointer right
	if *p+t.Value <= size-1 {
		*p += t.Value
	} else {
		*p = *p + t.Value - size
	}
}

func (t Token) Move_l(p *uint64, size uint64) {
	// moves the value pointer left
	if *p-t.Value >= 0 {
		*p -= t.Value
	} else {
		*p = *p + size - t.Value
	}
}

func (t Token) Inc(v *uint8) {
	// increments the value in the buffer under the value pointer
	if *v+uint8(t.Value) <= 255 {
		*v += uint8(t.Value)
	} else {
		*v = *v + uint8(t.Value) - 255 - 1
	}
}

func (t Token) Dec(v *uint8) {
	// decrements the value in the buffer under the value pointer
	if *v-uint8(t.Value) >= 0 {
		*v -= uint8(t.Value)
	} else {
		*v = *v + 255 - uint8(t.Value) + 1
	}
}

func (t Token) Print_value(c byte) {
	// prints out the value under the value pointer
	for i := 0; uint64(i) < t.Value; i++ {
		fmt.Printf("%c", c)
	}
}

func (t Token) Jump_if_zero(v uint8, i *int, d uint64) {
	if v == 0 {
		*i = int(d)
	} else {
		*i++
	}
}

func (t Token) Jump_if_non_zero(v uint8, i *int, d uint64) {
	if v != 0 {
		*i = int(d)
	} else {
		*i++
	}
}
