
# zk

A simple command-line tool for taking notes stored in plain text files.

## Install

1. Download the [latest version](https://github.com/mkaz/ak/releases) or,

2. Create your notes directory, default at `~/Documents/Notes`

## Usage

### Creating notes

1. Use `zk` opens new note in editor (default: vim)

2. Add note inline as a command-line argument
```
$ zk "Here is my note"
```

3. Use STDIN by piping or slurping in data
```
$ git log -1 | zk
$ zk < file-to-import.txt
```

You can create a tag using @tag which can then be used to search later.
```
$ zk "Here is my note with a @tag"
```

### Viewing Notes

Show the past week of notes
```
$ zk --week
```

Show today's notes
```
$ zk --today
```

Show a specific date's notes
```
$ zk --date 2017-10-15
```

Search by term
```
$ zk -s searchterm
```

Use @tag to search. If you only include tags on command-line, then zk assumes you are searching by tag. You can include multiple tags and it will OR search for all of them.
```
$ zk @tag
```

### Editing zks

Edit today's note
```
$ zk --today --edit
```

Edit a specific date
```
$ zk --date 2017-11-15 --edit
```


## Configuration

zk creates a single text file a day, `zk-2017-01-06.txt` organizing them in monthly and yearly folders. By default it stores them in `~/Documents/Notes/` that directory must already exist or it will give a warning.

zk will look for a config file `.zkrc` in your home directory in TOML format.

```
# zk config file

# base directory that all zks are stored
notesdir = '~/Documents/zks'

# golang timestamp format for note delimiter
# See: https://golang.org/pkg/time/
timestamp = '2006-01-02 3:04PM'
```

Note: For Windows, use single quotes for the directory config, otherwise the `\` seperator will be considered an escape and considered invalid.

## Errata

**Syncing data** - You can use a cloud or sync provider to keep your notes in-sync across computers. I'm happy with the open-source [Syncthing](https://syncthing.net/) to keep my Documents folder in sync across computers. I hear Google Drive, Dropbox, and others might be able to do the same.


## License

zk is open source and free to use, it is licensed under a <a rel="license" href="https://opensource.org/licenses/MIT">MIT License</a>.

I welcome any contributions. Please feel free to open an issues to report a bug, submit a feature. Pull requests are also welcome.

An [mkaz](https://mkaz.blog/) contrivance.

