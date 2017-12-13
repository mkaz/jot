/**
 * jot.go
 * utility functions
 */

package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

// Determine file to edit by date
func getFilepathDate(dt time.Time) (filename, dir string) {

	// build directory from date
	dir = filepath.Join(conf.Jotsdir, dt.Format("2006"), dt.Format("01"))

	// build filename from date
	filename = "jot-" + dt.Format("2006-01-02") + ".txt"
	var file = filepath.Join(dir, filename)

	return file, dir
}

func getJotsConfig() (conf Config) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configFile := filepath.Join(usr.HomeDir, ".jotrc")
	toml.DecodeFile(configFile, &conf)

	// set default diretory if not in config
	if conf.Jotsdir == "" {
		conf.Jotsdir = filepath.Join(usr.HomeDir, "Documents", "jots")
	}

	// fail out if directory does not exist, cant do anything
	// without the directory and its a good idea to let the
	// user create the first one, so they know where it is
	if _, err := os.Stat(conf.Jotsdir); os.IsNotExist(err) {
		errlog.Fatalln("Base directory does not exist", conf.Jotsdir)
	}

	// set default timestamp if not in config
	if conf.Timestamp == "" {
		conf.Timestamp = "2006-01-02 3:04PM"
	}

	return conf
}

func getJotFiles() (fa []string) {
	filepath.Walk(conf.Jotsdir, func(path string, fi os.FileInfo, err error) error {
		if !fi.IsDir() {
			if strings.Contains(path, "jot-") && strings.Contains(path, ".txt") {
				fa = append(fa, path)
			}
		}
		return nil
	})
	return fa
}
