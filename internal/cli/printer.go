package cli

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/namreg/godown-v2/internal/api"
)

const (
	okString  = "OK"
	nilString = "(nil)"
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
	fmt.Fprintf(p.out, "%s\r\n", str)
}

func (p *printer) printError(err error) {
	fmt.Fprintf(p.out, "Error: %s\n", err.Error())
}

func (p *printer) printResponse(resp *api.Response) {
	switch resp.Result.Type {
	case api.Response_OK:
		p.println(okString)
	case api.Response_NIL:
		p.println(nilString)
	case api.Response_STRING:
		p.println(fmt.Sprintf("(string) %s", resp.Result.Item))
	case api.Response_INT:
		if n, err := strconv.Atoi(resp.Result.Item); err != nil {
			p.printError(err)
		} else {
			p.println(fmt.Sprintf("(integer) %d", n))
		}
	case api.Response_HELP:
		p.println(strings.Replace(resp.Result.Item, "\n", "\r\n", -1))
	case api.Response_ERR:
		p.println(fmt.Sprintf("(error) %s", resp.Result.Item))
	case api.Response_SLICE:
		items := resp.Result.Items
		buf := new(bytes.Buffer)
		for i, v := range resp.Result.Items {
			buf.WriteString(fmt.Sprintf("%d) %q", i+1, v))
			if i != len(items)-1 { // check whether the current item is not last
				buf.WriteString("\r\n")
			}
		}
		p.println(buf.String())
	default:
		fmt.Fprintf(p.out, "%v\n", resp)
	}
}
