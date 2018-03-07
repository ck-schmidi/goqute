package goqute

import (
	"fmt"
	"log"
	"os/exec"
)

// Notifier can be used for string based notifications
type Notifier interface {
	Notify(message string, urgency UrgencyLevel)
}

// UrgencyLevel represents notify urgancy level
type UrgencyLevel string

const (
	// UrgencyLow for messages with low priority (debugging,…)
	UrgencyLow UrgencyLevel = "low"
	// UrgencyNormal for messages with normal priority (normal messages for user)
	UrgencyNormal UrgencyLevel = "normal"
	// UrgencyCritical for messages with high priority (errors, warnings,…)
	UrgencyCritical UrgencyLevel = "critical"
)

// Desktop Notifier

type desktopNotifier struct{}

func (dn desktopNotifier) Notify(message string, urgency UrgencyLevel) {
	cmd := exec.Command("notify-send", "-u", string(urgency), message)
	if err := cmd.Run(); err != nil {
		log.Printf("error while starting notify-send: %s\n", err)
	}
}

// NewDesktopNotifier returns a notifier for desktop notifications
func NewDesktopNotifier() Notifier {
	return desktopNotifier{}
}

// Command Line Notifier for tests

type clNotifier struct{}

func (cln clNotifier) Notify(message string, urgency UrgencyLevel) {
	log.Println(string(urgency) + ": " + message)
}

// NewCommandLineNotifier returns a notifier for desktop notifications
func NewCommandLineNotifier() Notifier {
	return clNotifier{}
}

func createNotifyMessage(message string, urgency UrgencyLevel) []byte {
	var messageCommands = map[UrgencyLevel]string{
		UrgencyLow:      "message-info",
		UrgencyNormal:   "message-info",
		UrgencyCritical: "message-error",
	}
	return []byte(fmt.Sprintf("%s \"%s\" \n", messageCommands[urgency], message))
}
