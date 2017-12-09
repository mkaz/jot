/**
 * jot.go
 * reading, display, search functions
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	chalk "github.com/danielchatfield/go-chalk"
)

// Show Last N Days of Jots
func showLastDays(n int) {
	n = n - 1 // we want to include today
	for n >= 0 {
		d := now.AddDate(0, 0, -1*n)
		showFileDate(d)
		n--
	}
}

// Display a Jot by File date
func showFileDate(dt time.Time) {
	prevBlank := false

	fn, _ := getFilepathDate(dt)
	f, err := os.Open(fn)
	if err != nil {
		return
	}
	defer f.Close()

	// walk through file line by line, so can add formating
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if tsRe.MatchString(line) {
			if prevBlank { // add empty line above date line
				fmt.Println("")
				prevBlank = false
			}
			if noColor {
				fmt.Println(line)
			} else {
				fmt.Println(chalk.Yellow(line))
			}

		} else {
			if prevBlank { // if prevBlank still here write it
				fmt.Println("|")
				prevBlank = false
			}
			if line == "" { // set as prevBlank with no output
				prevBlank = true
			} else {
				prevBlank = false
				fmt.Println("| " + line)
			}
		}
	}
}
