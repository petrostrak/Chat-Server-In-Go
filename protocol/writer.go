package protocol

import (
	"fmt"
	"io"
)

type CommandeWriter struct {
	writer io.Writer
}

func NewCommandWriter(w io.Writer) *CommandeWriter {
	return &CommandeWriter{
		writer: w,
	}
}

func (w *CommandeWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	return err
}

func (w *CommandeWriter) Write(cmd interface{}) error {
	var err error
	switch v := cmd.(type) {
	case SendCommand:
		err = w.writeString(fmt.Sprintf("SEND %v\n", v.Message))
	case MessageCommand:
		err = w.writeString(fmt.Sprintf("MESSAGE %v %v\n", v.Name, v.Message))
	case NameCommand:
		err = w.writeString(fmt.Sprintf("NAME %v\n", v.Name))
	default:
		err = UnknownCommand
	}

	return err
}
