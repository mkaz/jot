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
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// globals
var editFlag bool
var noColor bool
var n int
var now time.Time
var tsRe *regexp.Regexp
var errlog *log.Logger
var files []string
var template string

type Config struct {
	Jotsdir   string
	Timestamp string
}

var conf Config

func main() {
	errlog = log.New(os.Stderr, "", 0)
	now = time.Now()
	tsRe = regexp.MustCompile(`[0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+[APM]+ ::.*`)

	// parse command-line parameters
	var helpFlag = flag.Bool("help", false, "Display Help")
	var versionFlag = flag.Bool("version", false, "Display version")
	flag.BoolVar(&editFlag, "edit", false, "Edit flag, suppresses timestamp")
	flag.BoolVar(&noColor, "nocolor", false, "Use to disable colors")

	// date arguments
	flag.IntVar(&n, "n", 0, "Last N days to show")
	var ds = flag.String("date", "", "Show speific date yyyy-mm-dd")
	var today = flag.Bool("today", false, "Show todays note, alias -n 1")
	var week = flag.Bool("week", false, "Show last week, alias -n 7")
	var yesterday = flag.Bool("yesterday", false, "Set date to yesterday")
	var from = flag.String("from", "", "Show notes from date yyyy-mm-dd")
	var to = flag.String("to", "", "Show notes to date yyyy-mm-dd")

	var search = flag.String("s", "", "Search for term")
	flag.StringVar(&template, "t", "", "Template name to use")
	flag.Parse()
	args := flag.Args()

	if *helpFlag {
		usage()
	}

	if *versionFlag {
		fmt.Println("jot v0.4.1")
		os.Exit(0)
	}

	// read in config file if exists
	conf = getJotsConfig()

	if *today {
		n = 1
	}

	if *week {
		n = 7
	}

	if *yesterday {
		*ds = now.AddDate(0, 0, -1).Format("2006-01-02")
	}

	// retrieve the base jots directory
	files = getJotFiles()
	showFiles := false

	// --------------------------------------------------
	// Read / Search
	// --------------------------------------------------

	// show jots by last N number of days
	if n > 0 {
		if editFlag {
			if n > 1 {
				errlog.Fatalln("Editing multiple files at once is not yet supported.")
			}
			file, _ := getFilepathDate(now)
			openInEditor(file, nil)
			os.Exit(0)
		}
		// filter file list2
		files = filterFilesByCount(files, n)
		showFiles = true
	}

	// show jots for specific date
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

	if *from != "" {
		fdt, err := time.Parse("2006-01-02 15:04", *from+" 00:00")
		if err != nil {
			errlog.Fatalln("Error parsing from date, try format yyyy-mm-dd", err)
		}
		// filter files from
		files = filterFilesFromDate(files, fdt)
		showFiles = true
	}

	if *to != "" {
		tdt, err := time.Parse("2006-01-02 15:04", *to+" 23:59")
		if err != nil {
			errlog.Fatalln("Error parse to date, try format yyyy-mm-dd", err)
		}
		// filter files to
		files = filterFilesToDate(files, tdt)
		showFiles = true
	}

	if *search != "" {
		searchFiles(*search)
		os.Exit(0)
	}

	// check if tag search, the only arguments start with @
	if len(args) > 0 {
		tagSearch := true
		for _, a := range args {
			if !strings.HasPrefix(a, "@") {
				tagSearch = false
			}
		}
		if tagSearch {
			for _, a := range args {
				searchFiles(a)
			}
			os.Exit(0)
		}
	}

	if showFiles {
		for _, f := range files {
			showFileByPath(f)
		}
		os.Exit(0)
	}

	// --------------------------------------------------
	// Writing Jot
	// --------------------------------------------------

	// File: `${ jotsdir }/2017/12/jot-2017-12-06.txt`
	file, dir := getFilepathDate(now)

	// create directory if does not exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// check if received a command-line jot
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
	// check for edit flag, if editing we dont need
	// the template, or editor args to move to end
	if !editFlag {
		tpl := ""
		if template != "" {
			// attempt to read in template
			templateFile := filepath.Join(conf.Jotsdir, "tmpl"+template+".txt")
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

// Display Usage
func usage() {
	fmt.Println("usage: jot [args] [note]\n")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
	os.Exit(0)
}
