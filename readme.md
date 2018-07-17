
## Jot

A simple command-line tool for taking notes stored in plain text files.

I revived my simple Python script jot.py and rewrote in Golang, including a few features inspired by [jrnl.sh](http://jrnl.sh/) which is a really nice tool. However, a few things I wanted to do differently. First, I like storing the notes in a text file per day, whereas jrnl stores all notes together in a single file.

Also another key feature was the ability to use templates, so you can open a note with a prefilled template. Jot supports any number of user defined templates. See [#using templates](#using-templates)

### Install

1. Download the latest version
* [From the Jot releases page](https://github.com/mkaz/jot/releases) or,
* if you golang, install using:  `go get github.com/mkaz/jot`

2. Create your jots directory, default at `~/Documents/jots`


### Usage

#### Creating Jots

1. Using `jot` opens in editor (default: vim)

2. Jot a note inline as a command-line argument
```
$ jot "Here is my note"
```

3. Use STDIN by piping or slurping in data
```
$ git log -1 | jot
$ jot < import.txt
```

You can create a tag using @tag which can then be used to search later.
```
$ jot "Here is my note with a @tag"
```

#### Viewing Jots

Show the past week of notes
```
$ jot --week
```

Show today's notes
```
$ jot --today
```

Show a specific date's notes
```
$ jot --date 2017-10-15
```

Search by term
```
$ jot -s searchterm
```

Use @tag to search. If you only include tags on command-line, then jot assumes you are searching by tag. You can include multiple tags and it will OR search for all of them.
```
$ jot @tag
```

#### Editing Jots

Edit today's note
```
$ jot --today --edit
```

Edit a specific date
```
$ jot --date 2017-11-15 --edit
```

#### Using templates

Create a text file with you template named tmpl1.txt in the Jots directory (~/Documents/jots/). You can then create a new jot with that template using:
```
$ jot -t 1
```

You can create any number of templates, also the number is just a convenience. You can name your template tmplJoe.txt and then you would open it using `jot -t Joe`

I use the following template when I am trying to debug a tricky problem.
```
	PROBLEM:
	HYPOTHESIS:
	EXPERIMENT:
	RESULTS:
	CONCLUSION:
```

### Settings

Jot creates a single text file a day, `jot-2017-01-06.txt` organizing them in monthly and yearly folders. By default it stores them in `~/Documents/jots/` that directory must already exist or it will give a warning.

Jot will look for a config file `.jotrc` in your home directory in TOML format.

```
# Jot config file

# base directory that all jots are stored
jotsdir = '~/Documents/jots'

# golang timestamp format for note delimiter
# See: https://golang.org/pkg/time/
timestamp = '2006-01-02 3:04PM'
```

Note: For Windows, use single quotes for the directory config, otherwise the `\` seperator will be considered an escape and considered invalid.

### Errata

**Syncing data** - You can use a cloud or sync provider to keep your notes in-sync across computers. I'm happy with the open-source [Syncthing](https://syncthing.net/) to keep my Documents folder in sync across computers. I hear Google Drive, Dropbox, and others might be able to do the same.


### License

Jot is open source and free to use, it is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.

I welcome any contributions. Please feel free to open an issues to report a bug, submit a feature. Pull requests are also welcome.

An [mkaz](https://mkaz.blog/) contrivance.

