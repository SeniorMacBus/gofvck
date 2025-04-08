package executor

import (
	"gofvck/token"
)

func Execute(tokens []token.Token, buffer []uint8, loop_start map[uint64]uint64, loop_end map[uint64]uint64) {
	var p uint64 = 0
	var i int = 0
	size := len(buffer)
	for tokens[i].Type != token.CODE_END {
		t := tokens[i]
		switch t.Type {
		case token.MOVE_R:
			t.Move_r(&p, uint64(size))
			i++
		case token.MOVE_L:
			t.Move_l(&p, uint64(size))
			i++
		case token.INC:
			t.Inc(&buffer[p])
			i++
		case token.DEC:
			t.Dec(&buffer[p])
			i++
		case token.PRINT:
			t.Print_value(buffer[p])
			i++
		case token.LOOP_START:
			t.Jump_if_zero(buffer[p], &i, loop_start[uint64(i)]+1)
		case token.LOOP_END:
			t.Jump_if_non_zero(buffer[p], &i, loop_end[uint64(i)]+1)
		}
	}
}
