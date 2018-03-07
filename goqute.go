package goqute

import (
	"io"
	"net/url"
	"os"
)

// QuteBrowserInstance represents the single qutebrowser instance during
// the execution of the userscript. QuteBrowserInstance provides parameters set by
// qutebrowser and functionality for interacting with qutebrowser
type QuteBrowserInstance struct {
	Params  QBParams
	Session Session
	pipe    io.Writer
}

// Tab represents a single qutebrowser tab
type Tab struct {
	Title string
	URL   url.URL
}

// Notify notifies qutebrowser with given message and urgengy level
func (q QuteBrowserInstance) Notify(message string, urgency UrgencyLevel) {
	q.ExecuteCommand(createNotifyMessage(message, urgency))
}

// ExecuteCommand executes any command which can be handled by qutebrowser
func (q QuteBrowserInstance) ExecuteCommand(command []byte) {
	q.pipe.Write(command)
}

// NewQuteBrowserInstance returns an already intialized QuteBrowserInstance
// by reading values from environment set by qutebrowser
func NewQuteBrowserInstance() (*QuteBrowserInstance, error) {
	// Create parameters from QB
	params, err := CreateQBParams()
	if err != nil {
		return nil, err
	}
	// Open named pipe for communication back to qutebrowser
	pipe, err := os.OpenFile(params.Fifo, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return nil, err
	}
	// Read current session
	return &QuteBrowserInstance{
		Params: *params,
		pipe:   pipe,
	}, nil
}
