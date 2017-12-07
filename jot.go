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
var help bool
var now time.Time
var errlog *log.Logger

func main() {
	errlog = log.New(os.Stderr, "", 0)
	now = time.Now()

	// parse command-line parameters
	flag.BoolVar(&help, "help", false, "Display Help")
	flag.Parse()
	if help {
		usage()
	}

	// File: `${ jotsdir }/2017/12/jot-2017-12-06.txt`
	var file = getFilename()

	// check if received a command-line note
	// if so append to file and finish
	args := flag.Args()
	if len(args) > 0 {
		jot := strings.Join(args[0:], " ")
		fmt.Println(">> Input:", jot)
		os.Exit(1)
	}

	// terminal checks if stdin is associated with terminal or pipe
	// if we just try to read from os.Stdin directly it hangs waiting
	// for user input
	if !terminal.IsTerminal(0) {
		data, _ := ioutil.ReadAll(os.Stdin)
		fmt.Println(string(data))
		os.Exit(1)
	}

	// open file in editor
	cmd := exec.Command("vim", file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()

}

// Determine file to edit by date
func getFilename() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// check if directories exists
	var base = filepath.Join(usr.HomeDir, "Documents", "jots")
	if _, err := os.Stat(base); os.IsNotExist(err) {
		errlog.Fatalln("Base directory: ~/Documents/jots/ does not exist")
	}

	var dir = filepath.Join(base, now.Format("2006"), now.Format("01"))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// build filename from date
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
