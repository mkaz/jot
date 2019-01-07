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
func filterFilesByCount(fs []string, n int) (files []string) {
	n = n - 1 // we want to include today
	for n >= 0 {
		dt := now.AddDate(0, 0, -1*n)
		fn, _ := getFilepathDate(dt)
		if elemExists(fn, fs) {
			files = append(files, fn)
		}
		n--
	}
	return files
}

// Filter file list return files from date forward
func filterFilesFromDate(fs []string, dt time.Time) (files []string) {
	for _, f := range fs {
		d := getDateFromFile(f, "23:59")
		if d.After(dt) {
			files = append(files, f)
		}
	}
	return files
}

// Filter file list return files to date backward
func filterFilesToDate(fs []string, dt time.Time) (files []string) {
	for _, f := range fs {
		d := getDateFromFile(f, "00:00")
		if d.Before(dt) {
			files = append(files, f)
		}
	}
	return files
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
					fmt.Println(chalk.Yellow("-------------------------------------------------"))
					// highlight search term
					words := strings.Split(note, " ")
					for _, word := range words {
						if word == term {
							fmt.Print(chalk.Green(word) + " ")
						} else {
							fmt.Print(word + " ")
						}
					}
				}
			}
			fmt.Println("\n")
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
