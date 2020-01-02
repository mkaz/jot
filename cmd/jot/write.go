/**
 * jot.go
 * writing, editing functions
 */

package main

import (
	"os"
	"os/exec"
)

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
		var timestamp = now.Format(conf.Timestamp)
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
