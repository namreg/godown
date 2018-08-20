package cli

import (
	"fmt"
	"io"

	"github.com/namreg/godown-v2/internal/api"
)

type printer struct {
	out io.Writer
}

func newPrinter(out io.Writer) *printer {
	return &printer{out: out}
}

//Close closes the printer
func (p *printer) Close() error {
	if cl, ok := p.out.(io.Closer); ok {
		return cl.Close()
	}
	return nil
}

func (p *printer) println(str string) {
	fmt.Fprintln(p.out, str)
}

func (p *printer) printError(err error) {
	fmt.Fprintf(p.out, "Error: %s\n", err.Error())
}

func (p *printer) printResponse(resp *api.Response) {
	fmt.Fprintf(p.out, "%v\n", resp)
}
