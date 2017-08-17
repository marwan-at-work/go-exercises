package chatty

import (
	"io"
	"time"
)

// Chat takes an io.Writer and talks about life for a few seconds.
func Chat(w io.Writer) {
	sentences := []string{
		"hello\n",
		"I hope your day is going well.\n",
		"It should be better when this exercise is finished!\n",
		"Interfaces let you abstract behavior.\n",
		"I'm writing to you, but you define what writing to you really means.\n",
	}

	for _, s := range sentences {
		time.Sleep(time.Second)

		w.Write([]byte(s))
	}
}
