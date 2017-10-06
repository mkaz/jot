
## Jot.py

A simple command-line tool for taking notes stored in plain text files.


### Setup

* Copy jot.py to $HOME/bin
* Requires python 


### Usage

#### Creating Jots

1. Jot as argument

```
$ jot.py "Here is my note"
```

2. Free form jot
```
$ jot.py
<< Enter Text to Jot (ctrl-d to end) >>
Here is my 
multi-line note.
```

3. Using STDIN
```
$ git log -1 | jot
```


#### Viewing Jots

Show the full week of notes
```
$ jot --week 
```

Show today's notes
```
$ jot --today
```

Show a specific dates notes (only works in current week (TODO))
```
$ jot --date 2017-10-15
```

### Settings

The idea is for jot to create a week of notes and it stores each week in a plain text file. By default it uses `~/Documents/jots/` but will need that directory created. Jot creates weekly files something like `jot-2017-w40.md` with w40 for week #40.

Jot determines where to read/write the file by looking at:

1. Is command-line parameter `--filename` use that specific file
2. Is environment variable JOT_FILE set to the full path, use that specific file
3. Is environment variable JOT_DIR set and direcotry exists to create files in
4. Default to ~/Documents/jots/

Environment variables are an easy way to switch where notes are written depending on environment, for example creating work notes or home notes.

### Errata

**Syncing data** - You can use a cloud or sync provider to keep your notes in-sync across computers. I'm happy with the open-source [Syncthing](https://syncthing.net/) to keep my Documents folder in sync across computers. I hear Google Drive, Dropbox, and others might be able to do the same.

### License

Jot.py is open source and free to use, it is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.
