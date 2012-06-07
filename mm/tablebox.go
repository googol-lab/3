package mm

import (
	"fmt"
	"io"
	. "nimble-cube/nc"
	"os"
)

//CONCEPT: Send interfaces over the fannels
//Slice, Slice3, Scalar, Vector, GPUSlice, ...
//Use type info for nice output, recycling decissions...

type TableBox struct {
	input  <-chan float64
	time   <-chan float64
	writer io.Writer
}

func NewTableBox(file string) *TableBox {
	box := new(TableBox)
	var err error
	box.writer, err = os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	CheckIO(err)
	return box
}

func (box *TableBox) Run() {
	ok := true
	for ok {
		var time float64
		time, ok = <-box.time
		value := <-box.input
		fmt.Fprintln(box.writer, time, "\t", value)
	}
}
