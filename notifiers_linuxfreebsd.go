// +build linux freebsd

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const (
	specificPart = `
    NOTI_VOICE
        Name of voice used for speech notifications. See "espeak --voices" for
        available voices.`
)

func init() {
	flag.Usage = func() {
		fmt.Printf(manual, specificPart)
	}
}

// bannerNotify triggers a Notify notification.
func bannerNotify() {
	_, err := exec.LookPath("notify-send")
	if err != nil {
		log.Fatal("Install 'notify-send' and try again")
	}

	cmd := exec.Command("notify-send", *title, *message)
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// speechNotify triggers an eSpeak notification.
func speechNotify() {
	_, err := exec.LookPath("espeak")
	if err != nil {
		log.Println("Install 'espeak' and try again")

		var errStr string
		if runtime.GOOS == "freebsd" {
			errStr = "On FreeBSD this might be: 'sudo pkg install --yes espeak'"
		} else if runtime.GOOS == "linux" {
			errStr = "On Linux this might be: 'sudo apt-get install --yes espeak'"
		}

		log.Fatal(errStr)
	}

	voice := os.Getenv(voiceEnv)
	if voice == "" {
		voice = "english-us"
	}

	*message = fmt.Sprintf("%s %s", *title, *message)

	cmd := exec.Command("espeak", "-v", voice, "--", *message)
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
}