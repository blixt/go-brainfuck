# Go Brainfuck

This is a simple module to run or debug Brainfuck code.

## Installation

    go get github.com/blixt/go-brainfuck/brainfuck

## Examples

### Hello World

    package main
    
    import (
    	"github.com/blixt/go-brainfuck/brainfuck"
    	"os"
    )
    
    func main() {
    	// Print "Hello" and exit.
    	brainfuck.Run("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.", "", os.Stdout)
    }

### More advanced & debugging

    package main
    
    import (
    	"fmt"
    	"github.com/blixt/go-brainfuck/brainfuck"
    	"io/ioutil"
    )
    
    func main() {
    	bf := &brainfuck.State{
    		Code:   "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.",
    		Memory: make([]byte, 100),
    		Output: ioutil.Discard, // Ignore output
    		Debug:  true,
    	}
    
    	// Do 100 iterations.
    	for i := 0; i < 100; i++ {
    		bf.Step()
    	}
    
    	// Print the current memory.
    	fmt.Println(bf.Memory)
    }

## Notes

In case of an EOF read from input, the `,` operator will be a no-op, unlike
some brainfuck implementations that will set the current value to 0 or -1.
