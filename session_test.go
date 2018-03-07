package goqute

import (
	"fmt"
	"os"
	"testing"
)

func TestExtractSession(t *testing.T) {
	file, err := os.Open("session_test.yml")
	if err != nil {
		t.Error(err)

	}

	session, err := ExtractSession(file)
	if err != nil {
		t.Error(err)
	}

	for i, window := range session.Windows {
		fmt.Println("window", i)
		for j, tab := range window.Tabs {
			fmt.Println("  tab", j)
			for _, historyEntry := range tab.History {
				fmt.Println("    url", historyEntry.URL)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	_ = session

}
