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

	"github.com/marcusolsson/tui-go"
	"golang.org/x/crypto/ssh/terminal"
)

// globals
var editFlag bool
var noColor bool
var n int
var now time.Time
var tsRe *regexp.Regexp
var errlog *log.Logger
var template string
var args []string

type Config struct {
	Jotsdir   string
	Timestamp string
}

var ui tui.UI
var conf Config
var panel *tui.Box
var fileList *tui.List

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
	flag.StringVar(&template, "t", "", "Template name to use")
	flag.Parse()
	args = flag.Args()

	if *helpFlag {
		usage()
	}

	if *versionFlag {
		fmt.Println("jot v0.5.0")
		os.Exit(0)
	}

	// read in config file if exists
	conf = getJotsConfig()

	// Read in files
	files := getJotFiles()

	// TUI

	fileList = tui.NewList()
	fileList.OnItemActivated(func(l *tui.List) {
		sel := l.SelectedItem()
		fp := filepath.Join(conf.Jotsdir, sel)
		openInEditor(fp, nil)
		ui.Repaint() // doesnt work :(
	})

	updateFileList(files)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	// Perform search as typing
	input.OnChanged(func(e *tui.Entry) {
		// if more than two-chars
		// search files
		text := e.Text()
		if len(text) >= 2 {
			found := searchFiles(files, text)
			updateFileList(found)
		}
	})

	// Create new document onSubmit
	input.OnSubmit(func(e *tui.Entry) {
		// fmt.Println(?)
	})

	tui.DefaultFocusChain.Set(input, fileList)
	panel = tui.NewVBox(inputBox, fileList)
	panel.SetSizePolicy(tui.Expanding, tui.Expanding)

	ui, _ = tui.New(panel)
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Tab", func() {
		// what to do when tab
	})

	if err := ui.Run(); err != nil {
		log.Fatalln("Error starting UI", err)
	}
}

func updateFileList(files []string) {
	// scrub file names for display
	files = displayFilenames(files)
	fileList.RemoveItems() // start fresh
	fileList.AddItems(files...)
}

func displayFilenames(files []string) (rtn []string) {
	for _, f := range files {
		f = strings.Replace(f, conf.Jotsdir, "", 1)
		f = strings.TrimLeft(f, "/")
		rtn = append(rtn, f)
	}
	return rtn
}

func cli() {

	// TODO: parse command-line args and create if something
	//      is being piped in
	file := "/tmp/foo"

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
}

// Display Usage
func usage() {
	fmt.Println("usage: jot [args] [note]\n")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
	os.Exit(0)
}
