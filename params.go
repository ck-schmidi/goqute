package goqute

import (
	"fmt"
	"os"
)

// Mode is the qb mode, in which the userscript
// was started
type Mode int

const (
	// ModeHints is the value for the qb hint mode
	ModeHints = iota
	// ModeCommand is the value for the qb command mode
	ModeCommand
)

// QBParams represents the parameters
// set by qutebrowser
type QBParams struct {
	Mode            Mode   `json:"mode"`
	UserAgent       string `json:"user_agent"`
	Fifo            string `json:"fifo"`
	HTML            string `json:"html"`
	Text            string `json:"text"`
	ConfigDir       string `json:"config_dir"`
	DataDir         string `json:"data_dir"`
	DownloadDir     string `json:"download_dir"`
	CommandlineText string `json:"commandline_text"`
	URL             string `json:"url"`
	SelectedText    string `json:"selected_text"`

	// only available in command mode
	Title string `json:"title"`

	// only available in hints mode
	SelectedHTML string `json:"selected_html"`
}

// CreateQBParams creates a QBParams objects
// with the values set by qutebrowser
func CreateQBParams() (*QBParams, error) {

	// Mode
	var mode Mode
	m := os.Getenv("QUTE_MODE")
	if m == "hints" {
		mode = ModeHints
	} else if m == "command" {
		mode = ModeCommand
	} else {
		return nil, fmt.Errorf("could not find qutebrowser environment")
	}

	return &QBParams{
		Mode:            mode,
		UserAgent:       os.Getenv("QUTE_USER_AGENT"),
		Fifo:            os.Getenv("QUTE_FIFO"),
		HTML:            os.Getenv("QUTE_HTML"),
		Text:            os.Getenv("QUTE_TEXT"),
		ConfigDir:       os.Getenv("QUTE_CONFIG_DIR"),
		DataDir:         os.Getenv("QUTE_DATA_DIR"),
		DownloadDir:     os.Getenv("QUTE_DOWNLOAD_DIR"),
		CommandlineText: os.Getenv("QUTE_COMMANDLINE_TEXT"),
		URL:             os.Getenv("QUTE_URL"),
		Title:           os.Getenv("QUTE_TITLE"),
		SelectedText:    os.Getenv("QUTE_SELECTED_TEXT"),
		SelectedHTML:    os.Getenv("QUTE_SELECTED_HTML"),
	}, nil
}
