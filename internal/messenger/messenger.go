package messenger

import (
	"bufio"
	"net"

	"github.com/karpovicham/secured-tcp-server/internal/proto"
)

// Delimiter indicates the end of the data for the reader.
const Delimiter = '\n'

// Messenger interface to communicate between clint and server
type Messenger interface {
	// Send Writes data to pipeline
	Send(msgType proto.Type, msgData []byte) error
	// Receive Reads data from pipeline
	Receive() (*proto.Message, error)
}

type msgr struct {
	*bufio.ReadWriter
}

// MsgrFn is helper type for NewMessenger function
type MsgrFn func(conn net.Conn) Messenger

// NewMessenger returns Messenger with set up Reader and Writer
// consistent with MsgrFn function
func NewMessenger(conn net.Conn) Messenger {
	return &msgr{
		bufio.NewReadWriter(
			bufio.NewReader(conn),
			bufio.NewWriter(conn),
		),
	}
}

// Send Writes data
func (m *msgr) Send(msgType proto.Type, msgData []byte) error {
	msg := proto.Message{
		Type: msgType,
		Data: msgData,
	}

	msgJson, err := msg.MarshalJSON()
	if err != nil {
		return err
	}

	msgJson = append(msgJson, Delimiter)
	if _, err := m.Write(msgJson); err != nil {
		return err
	}

	return m.Flush()
}

// Receive Reads data
func (m *msgr) Receive() (*proto.Message, error) {
	data, err := m.ReadBytes(Delimiter)
	if err != nil {
		return nil, err
	}

	msg := new(proto.Message)
	if err = msg.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	return msg, nil
}
