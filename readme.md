
## Jot

A simple command-line tool for taking notes stored in plain text files.

I decided to revive my simple Python script and rewrite in Golang. Plus
including new features inspired by [jrnl.sh](http://jrnl.sh/) which is
a really nice tool.

A few things I wanted to do differently. First, I like storing the notes
in a text file per day, whereas jrnl stores all notes together in a
single file.

Another key feature was the ability to use templates, so you can open a
note with a prefilled template. Jot supports any number of user defined
templates.



### Usage

#### Creating Jots

1. Jot as argument

```
$ jot "Here is my note"
```

2. Type just `jot` to open editor (default: vim)

3. Use STDIN
```
$ git log -1 | jot
$ jot < import.txt
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

Show a specific dates notes
```
$ jot --date 2017-10-15
```


#### Editing Jots

Edit today's note
````
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



### Settings

Jot creates a single text file a day, `jot-2017-01-06.txt` organizing them in monthly and yearly folders. By default it stores them in `~/Documents/jots/` but that directory must already exist.


### Errata

**Syncing data** - You can use a cloud or sync provider to keep your notes in-sync across computers. I'm happy with the open-source [Syncthing](https://syncthing.net/) to keep my Documents folder in sync across computers. I hear Google Drive, Dropbox, and others might be able to do the same.


### License

Jot is open source and free to use, it is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.

I welcome any contributions. Please feel free to open an issues to report a bug, submit a feature. Pull requests are also welcome.

An [mkaz](https://mkaz.tech/) contrivance.

