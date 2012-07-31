package brainfuck

import (
	"fmt"
	"io"
	"log"
	"strings"
)

const (
	MAX_ITERATIONS = 90000
	MEMORY_SIZE    = 30000
)

type State struct {
	Code        string
	Instruction int
	Iterations  int
	Pointer     int
	Memory      []byte
	Input       io.Reader
	Output      io.Writer
	Debug       bool
}

func (s *State) Reset() {
	s.Instruction = 0
	s.Iterations = 0
	s.Pointer = 0
	s.Memory = make([]byte, MEMORY_SIZE)
}

func (s *State) Step() (running bool, err error) {
	if s.Instruction < 0 || s.Instruction >= len(s.Code) {
		err = fmt.Errorf("Tried to step but not in an executing state")
		return
	}

	if s.Iterations > MAX_ITERATIONS {
		err = fmt.Errorf("Execution timeout reached")
		return
	}
	s.Iterations++

	if s.Debug {
		log.Printf("[%05d] @%05d(=0x%02x)  %04d: %c\n", s.Iterations, s.Pointer, s.Memory[s.Pointer], s.Instruction, s.Code[s.Instruction])
	}

	switch s.Code[s.Instruction] {
	case '>':
		if s.Pointer >= len(s.Memory)-1 {
			err = fmt.Errorf("Memory pointer out of bounds")
			return
		}
		s.Pointer++
	case '<':
		if s.Pointer <= 0 {
			err = fmt.Errorf("Memory pointer out of bounds")
			return
		}
		s.Pointer--
	case '+':
		s.Memory[s.Pointer]++
	case '-':
		s.Memory[s.Pointer]--
	case '.':
		var n int
		if n, err = s.Output.Write([]byte{s.Memory[s.Pointer]}); err != nil {
			return
		} else if n != 1 {
			err = fmt.Errorf("Failed to write one byte to output")
			return
		}
	case ',':
		in := make([]byte, 1)
		if n, err := s.Input.Read(in); err == nil && n == 1 {
			s.Memory[s.Pointer] = in[0]
		}
	case '[':
		var depth int
		if s.Memory[s.Pointer] == 0 {
			s.Instruction++
			for depth > 0 || s.Code[s.Instruction] != ']' {
				if s.Code[s.Instruction] == '[' {
					depth++
				} else if s.Code[s.Instruction] == ']' {
					depth--
				}
				s.Instruction++
			}
		}
	case ']':
		var depth int
		s.Instruction--
		for depth > 0 || s.Code[s.Instruction] != '[' {
			if s.Code[s.Instruction] == ']' {
				depth++
			} else if s.Code[s.Instruction] == '[' {
				depth--
			}
			s.Instruction--
		}
		s.Instruction--
	}

	s.Instruction++

	if s.Instruction >= len(s.Code) {
		if s.Debug {
			log.Printf("[%05d] @%05d(=0x%02x)  END\n", s.Iterations, s.Pointer, s.Memory[s.Pointer])
		}
	} else {
		running = true
	}

	return
}

func Run(code string, input string, output io.Writer) (state *State, err error) {
	state = &State{
		Code:   code,
		Memory: make([]byte, MEMORY_SIZE),
		Input:  strings.NewReader(input),
		Output: output,
	}

	for {
		var running bool
		if running, err = state.Step(); !running {
			break
		}
	}

	return
}
