/**
 * jot.go
 *
 * A simple command-line tool for taking notes stored in plain text files.
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	chalk "github.com/danielchatfield/go-chalk"
	"golang.org/x/crypto/ssh/terminal"
)

// globals
var helpFlag bool
var editFlag bool
var noColor bool
var n int
var now time.Time
var tsRe *regexp.Regexp
var errlog *log.Logger
var jotsdir string
var template string

func main() {
	errlog = log.New(os.Stderr, "", 0)
	now = time.Now()
	tsRe = regexp.MustCompile(`[0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+[APM]+ ::.*`)

	// parse command-line parameters
	flag.BoolVar(&helpFlag, "help", false, "Display Help")
	flag.BoolVar(&editFlag, "edit", false, "Edit flag, suppresses timestamp")
	flag.BoolVar(&noColor, "nocolor", false, "Use to disable colors")
	flag.IntVar(&n, "n", 0, "Last N days to show")
	var ds = flag.String("date", "", "Show speific date yyyy-mm-dd")
	var today = flag.Bool("today", false, "Show todays note, alias -n 1")
	var week = flag.Bool("week", false, "Show last week, alias -n 7")
	flag.StringVar(&template, "t", "", "Template name to use")
	flag.Parse()

	if helpFlag {
		usage()
	}

	// TODO: config for jots dir
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	jotsdir = filepath.Join(usr.HomeDir, "Documents", "jots")
	// check if directories exists
	if _, err := os.Stat(jotsdir); os.IsNotExist(err) {
		errlog.Fatalln("Base directory does not exist", jotsdir)
	}

	if *today {
		n = 1
	}

	if *week {
		n = 7
	}

	// Read / Search
	// TODO: add reading, searching notes

	// show jots by last N number of days
	if n > 0 {
		if editFlag {
			if n > 1 {
				errlog.Fatalln("Editing multiple files at once is not yet supported.")
			}
			file, _ := getFilepathDate(now)
			openInEditor(file, nil)
		}
		showLastDays(n)
		os.Exit(0)
	}

	// show jots by date
	if *ds != "" {
		dt, err := time.Parse("2006-01-02", *ds)
		if err != nil {
			errlog.Fatalln("Error parsing date, try format yyyy-mm-dd", err)
		}
		if editFlag {
			file, _ := getFilepathDate(dt)
			openInEditor(file, nil)
		} else {
			showFileDate(dt)
		}
		os.Exit(0)
	}

	// Writing Jot
	// TODO: templates

	// File: `${ jotsdir }/2017/12/jot-2017-12-06.txt`
	file, dir := getFilepathDate(now)

	// create directory if does not exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// check if received a command-line jot
	args := flag.Args()
	if len(args) > 0 {
		jot := strings.Join(args[0:], " ")
		writeFile(file, jot)
		os.Exit(0)
	}

	// terminal checks if stdin is associated with terminal or pipe
	// if we just try to read from os.Stdin directly it hangs waiting
	// for user input
	if !terminal.IsTerminal(0) {
		data, _ := ioutil.ReadAll(os.Stdin)
		writeFile(file, string(data))
		os.Exit(0)
	}

	// open file in editor
	var editorArgs []string
	if !editFlag {
		tpl := ""
		if template != "" {
			// attempt to read in template
			templateFile := filepath.Join(jotsdir, "tmpl"+template+".txt")
			content, err := ioutil.ReadFile(templateFile)
			if err == nil {
				tpl = string(content)
			}
		}
		writeFile(file, tpl)                          // call writeFile which creates/appends timestamp
		editorArgs = append(editorArgs, "+normal G$") // position at end of file open cursor
	}
	openInEditor(file, editorArgs)
}

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

func openInEditor(file string, args []string) {
	// TODO: use env var EDITOR
	// "+" for vim positions cursor on last line
	args = append(args, file)
	cmd := exec.Command("vim", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Write Jot to File
func writeFile(filename, text string) {
	if !editFlag { // append timestamp
		// TODO: config for timestamp date string
		var timestamp = now.Format("2006-01-02 3:04PM")
		text = "\n" + timestamp + " :: " + text
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		errlog.Fatalln("Error opening file to append:", err)
	}
	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		errlog.Fatalln("Error writing file:", err)
	}
}

// Determine file to edit by date
func getFilepathDate(dt time.Time) (filename, dir string) {

	// build directory from date
	dir = filepath.Join(jotsdir, dt.Format("2006"), dt.Format("01"))

	// build filename from date
	filename = "jot-" + dt.Format("2006-01-02") + ".txt"
	var file = filepath.Join(dir, filename)

	return file, dir
}

// Display Usage
func usage() {
	fmt.Println("usage: jot [-help]\n")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
	os.Exit(0)
}
