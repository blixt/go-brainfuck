# Go Brainfuck

This is a simple module to run or debug Brainfuck code.

## Installation

    go get github.com/blixt/go-brainfuck/brainfuck

## Example

    package main
    
    import (
    	"github.com/blixt/go-brainfuck/brainfuck"
    	"os"
    )
    
    func main() {
    	// Print "Hello" and exit.
    	brainfuck.Run("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.", "", os.Stdout)
    }

## Notes

In case of an EOF read from input, the `,` operator will be a no-op, unlike
some brainfuck implementations that will set the current value 0 or -1.
