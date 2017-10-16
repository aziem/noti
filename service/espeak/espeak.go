package espeak

import (
	"os"
	"os/exec"
)

// Notification is an espeak notification.
type Notification struct {
	// -g
	WordGap int
	// -p
	PitchAdjustment int
	// -s
	Rate int
	// -v
	VoiceName string

	Text string
}

// Send triggers a spoken notification.
func (n *Notification) Send() error {
	cmd := exec.Command("espeak", "-v", n.VoiceName, "--", n.Text)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// SetMessage sets a notification's message.
func (n *Notification) SetMessage(m string) {
	n.Text = m
}

// GetMessage gets a notification's message.
func (n *Notification) GetMessage() string {
	return n.Text
}
