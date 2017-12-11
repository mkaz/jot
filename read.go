/**
 * jot.go
 * reading, display, search functions
 */

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	prevBlank := false

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

func showSearchFor(term string) {
	fmt.Println("Searching for: ", term)
	filepath.Walk(jotsdir, searchFiles(term))
}

func searchFiles(term string) filepath.WalkFunc {
	return func(fn string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			if strings.Contains(fn, "jot-") && strings.Contains(fn, ".txt") {
				data, _ := ioutil.ReadFile(fn)
				if strings.Contains(string(data), term) {
					// parse all data into individual notes
					notes := parseDayToNotes(string(data))

					// display note with term
					for _, note := range notes {
						if strings.Contains(note, term) {
							fmt.Println(note)
						}
					}
				}
			}
		}
		return nil
	}
}

func parseDayToNotes(str string) (notes []string) {
	note := ""
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if tsRe.MatchString(line) {
			if note != "" { // append previous
				notes = append(notes, note)
			}
			// start new note
			note = line
		} else {
			note = note + line
		}
	}
	notes = append(notes, note)
	return notes
}
