package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/woothee/woothee-go"
)

type Log struct {
	UserAgent string          `json:"ua"`
	Result    *woothee.Result `json:"result"`
	Error     string          `json:"error"`
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for stdin.Scan() {
		ua := stdin.Text()
		r, err := woothee.Parse(ua)

		log := Log{
			UserAgent: ua,
			Result:    r,
		}
		if err != nil {
			log.Error = err.Error()
		}

		if err := enc.Encode(log); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}
