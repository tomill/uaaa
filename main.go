package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/woothee/woothee-go"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		ua := stdin.Text()

		r, err := woothee.Parse(ua)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		if err := json.NewEncoder(os.Stdout).Encode(r); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}
