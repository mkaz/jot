/**
 * jot.go
 * reading, display, search functions
 */

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
	fn, _ := getFilepathDate(dt)
	showFileByPath(fn)
}

func showFileByPath(fn string) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return
	}

	notes := parseDayToNotes(string(data))
	for _, note := range notes {
		lines := strings.Split(note, "\n")
		for idx, line := range lines {
			if idx == 0 {
				fmt.Println(chalk.Yellow(line))
			} else {
				fmt.Println("| " + line)
			}
		}
		fmt.Println()
	}
}

func searchFiles(term string) {
	for _, fn := range files {
		data, _ := ioutil.ReadFile(fn)
		if strings.Contains(string(data), term) {
			// parse all data into individual notes
			notes := parseDayToNotes(string(data))

			// display note with term
			for _, note := range notes {
				if strings.Contains(note, term) {
					// highlight search term
					words := strings.Split(note, " ")
					for _, word := range words {
						if word == term {
							fmt.Print(chalk.Red(word) + " ")
						} else {
							fmt.Print(word + " ")
						}
					}
				}
			}
		}
	}
}

func parseDayToNotes(str string) (notes []string) {
	note := ""
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if tsRe.MatchString(line) {
			if note != "" { // append previous
				notes = append(notes, strings.Trim(note, "\n"))
			}
			// start new note
			note = line + "\n"
		} else {
			note = note + line + "\n"
		}
	}
	notes = append(notes, strings.Trim(note, "\n"))
	return notes
}
