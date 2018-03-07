package goqute

import (
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Session represents the current session (opened tabs)
type Session struct {
	Windows []struct {
		Geometry interface{} `yaml:"geometry"`
		Tabs     []struct {
			History []struct {
				Active    bool `yaml:"active"`
				Pinned    bool `yaml:"pinned"`
				ScrollPos struct {
					X int `yaml:"x"`
					Y int `yaml:"y"`
				} `yaml:"scroll-pos"`
				Title string  `yaml:"title"`
				URL   string  `yaml:"url"`
				Zoom  float64 `yaml:"zoom"`
			} `yaml:"history"`
			Active bool `yaml:"active,omitempty"`
		} `yaml:"tabs"`
		Active bool `yaml:"active,omitempty"`
	} `yaml:"windows"`
}

// ExtractSession extract Session from an io.Reader source
// which provides a Session object as a YAML document
func ExtractSession(source io.Reader) (*Session, error) {

	// read from reader before unmarshaling
	content, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}

	var session Session
	err = yaml.Unmarshal(content, &session)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return &session, nil
}
