/**
 * jot.go
 * reading, display, search functions
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	chalk "github.com/danielchatfield/go-chalk"
)

// getJotFiles - reads in all files returning an array filenames
func getJotFiles() (fa []string) {
	filepath.Walk(conf.Jotsdir, func(path string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			if filepath.Ext(path) == ".md" || filepath.Ext(path) == ".txt" {
				fa = append(fa, path)
			}
		}
		return nil
	})
	return fa
}

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

// showFileDate displays a Jot by file date
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
		displayNote(note, "")
	}
}

func searchFiles(term string, files []string) {
	for _, fn := range files {
		data, _ := ioutil.ReadFile(fn)
		notes := parseDayToNotes(string(data))
		for _, note := range notes {
			if strings.Contains(strings.ToLower(note), strings.ToLower(term)) {
				displayNote(note, term)
				fmt.Println("\n")
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
	note = strings.TrimSpace(note)
	if note != "" {
		notes = append(notes, note)
	}
	return notes
}

// displayNote (output note to screen)
func displayNote(note string, term string) {
	lines := strings.Split(note, "\n")
	for idx, line := range lines {
		if idx == 0 {
			fmt.Println(chalk.Cyan(line))
		} else {
			line = highlightWord(line, term)
			fmt.Println("| " + line)
		}
	}
}

// highlightWord looks for term in str and adds color code
func highlightWord(str string, term string) (rtn string) {
	if term == "" {
		return str
	}

	if !strings.Contains(str, term) {
		return str
	}

	words := strings.Fields(str)
	for _, word := range words {
		if word == term {
			rtn += chalk.Magenta(word) + " "
		} else {
			rtn += word + " "
		}
	}

	return strings.TrimSpace(rtn)
}
