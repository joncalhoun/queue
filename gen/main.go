package main

import (
	"flag"
	"io"
	"os"
	"os/exec"
	"text/template"

	"github.com/joncalhoun/pipe"
)

type data struct {
	Type string
	Name string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "type", "", "The subtype used for the queue being generated")
	flag.StringVar(&d.Name, "name", "", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.Parse()

	// Create our template + other commands we want to run
	t := template.Must(template.New("queue").Parse(queueTemplate))

	rc, wc, _ := pipe.Commands(
		exec.Command("gofmt"),
		exec.Command("goimports"),
	)
	t.Execute(wc, d)
	wc.Close()
	io.Copy(os.Stdout, rc)
}

var queueTemplate = `
package queue

import (
	"container/list"
)

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{list.New()}
}

// {{.Name}} is a queue implementation for the {{.Type}} type.
// Behind the scenes it is a linked list FIFO queue
// that uses container/list under the hood. The primary
// motivation in creating this type is to allow the compiler
// to verify that we are using the correct types with our
// queue rather than dealing with the interface{} type in
// the rest of our code.
type {{.Name}} struct {
	list *list.List
}

// Len returns the total length of the queue
func (q *{{.Name}}) Len() int {
	return q.list.Len()
}

// Enqueue adds an item to the back of the queue
func (q *{{.Name}}) Enqueue(i {{.Type}}) {
	q.list.PushBack(i)
}

// Dequeue removes and returns the front item in the queue
func (q *{{.Name}}) Dequeue() {{.Type}} {
	if q.list.Len() == 0 {
		// You could opt to return errors here, but I personally
		// prefer to leave length checking up to end users kinda
		// like bounds checking in slices.
		panic(ErrEmptyQueue)
	}

	raw := q.list.Remove(q.list.Front())
	if typed, ok := raw.({{.Type}}); ok {
		return typed
	}

	// This won't ever happen unless someone has access to
	// insert things into the list with an invalid type or
	// your code has bug.
	panic(ErrInvalidType)
}
`
