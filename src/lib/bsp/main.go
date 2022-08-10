package main

import (
	"errors"
	"math/rand"
	"strings"
	"syscall/js"
	"time"

	"github.com/Nv7-Github/bsharp/backends/interpreter"
	"github.com/Nv7-Github/bsharp/ir"
	"github.com/Nv7-Github/bsharp/parser"
	"github.com/Nv7-Github/bsharp/tokens"
)

type fs struct{}

func (*fs) Parse(f string) (*parser.Parser, error) {
	return nil, errors.New("not implemented")
}

func runBsp(src string) string {
	rand.Seed(time.Now().UnixNano())

	// Build IR
	stream := tokens.NewStream("main.bsp", src)
	tok := tokens.NewTokenizer(stream)
	err := tok.Tokenize()
	if err != nil {
		return err.Error()
	}
	parse := parser.NewParser(tok)
	err = parse.Parse()
	if err != nil {
		return err.Error()
	}
	ir := ir.NewBuilder()
	err = ir.Build(parse, &fs{})
	if err != nil {
		if len(ir.Errors) > 0 {
			out := ""
			for _, err := range ir.Errors {
				out += err.Message + "\n"
			}
			return out
		} else {
			return err.Error()
		}
	}

	// Run IR
	out := &strings.Builder{}
	interp := interpreter.NewInterpreter(ir.IR(), out)
	err = interp.Run()
	if err != nil {
		return err.Error()
	}
	return out.String()
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("runBspWasm", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return js.ValueOf(runBsp(args[0].String()))
	}))

	<-c // Keep it running
}
