/**
 * jot.go
 *
 * A simple command-line tool for taking notes stored in plain text files.
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// globals
var helpFlag bool
var editFlag bool
var now time.Time
var errlog *log.Logger

func main() {
	errlog = log.New(os.Stderr, "", 0)
	now = time.Now()

	// parse command-line parameters
	flag.BoolVar(&helpFlag, "help", false, "Display Help")
	flag.BoolVar(&editFlag, "edit", false, "Edit flag, suppresses timestamp")
	flag.Parse()

	if helpFlag {
		usage()
	}

	// Read / Search
	// TODO: add reading, searching notes

	// Writing Jot
	// TODO: templates

	// File: `${ jotsdir }/2017/12/jot-2017-12-06.txt`
	file := getFilename()

	// check if received a command-line jot
	args := flag.Args()
	if len(args) > 0 {
		jot := strings.Join(args[0:], " ")
		writeFile(file, jot)
		os.Exit(1)
	}

	// terminal checks if stdin is associated with terminal or pipe
	// if we just try to read from os.Stdin directly it hangs waiting
	// for user input
	if !terminal.IsTerminal(0) {
		data, _ := ioutil.ReadAll(os.Stdin)
		writeFile(file, string(data))
		os.Exit(1)
	}

	// open file in editor
	var editorArgs []string

	if !editFlag {
		writeFile(file, "")                           // call writeFile which creates/appends timestamp
		editorArgs = append(editorArgs, "+normal G$") // position at end of file open cursor
	}

	// TODO: use env var EDITOR
	// "+" for vim positions cursor on last line
	editorArgs = append(editorArgs, file)
	cmd := exec.Command("vim", editorArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func writeFile(filename, text string) {

	// append timestamp
	if !editFlag {
		// TODO: config for timestamp date string
		var timestamp = now.Format("2006-01-02 3:04PM")
		text = "\n" + timestamp + " :: " + text
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		errlog.Fatalln("Error opening file to append:", err)
	}
	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		errlog.Fatalln("Error writing file:", err)
	}
}

// Determine file to edit by date
func getFilename() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// check if directories exists
	// TODO: config for base directory
	var base = filepath.Join(usr.HomeDir, "Documents", "jots")
	if _, err := os.Stat(base); os.IsNotExist(err) {
		errlog.Fatalln("Base directory: ~/Documents/jots/ does not exist")
	}

	var dir = filepath.Join(base, now.Format("2006"), now.Format("01"))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// build filename from date
	// TODO: config for filename
	filename := "jot-" + now.Format("2006-01-02") + ".txt"
	var file = filepath.Join(dir, filename)

	return file
}

// Display Usage
func usage() {

	fmt.Println("usage: jot [-help]\n")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
	os.Exit(1)
}
