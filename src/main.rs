// extern crate toml;

use clap::{crate_version, App, Arg};
use std::env;
use std::fs;
use std::io::prelude::*;
use std::io::{self, Read};
use std::process::Command;

mod config;
mod utils;

fn main() {
    let args = App::new("jot")
        .version(crate_version!())
        .about("Jot notes on the command-line")
        .author("Marcus Kazmierczak")
        .arg(
            Arg::new("config")
                .about("Configuration file")
                .short('c')
                .long("config")
                .takes_value(true),
        )
        .arg(Arg::new("monthly").about("Monthly note").long("monthly"))
        .arg(Arg::new("weekly").about("Weekly note").long("weekly"))
        .arg(Arg::new("daily").about("Daily note").long("daily"))
        .arg(Arg::new("new").about("New note").long("new"))
        .arg(
            Arg::new("content")
                .about("Create note from command-line")
                .multiple_values(true),
        )
        .after_help(
            "Create notes:
    1. Pipe into jot
        echo 'Hello' | jot

    2. Command-line args:
        jot 'This is my note'

    3. No args or pipe opens new note in $EDITOR",
        )
        .get_matches();

    let mut content = String::new();

    // read in config
    let config = config::get_config(args.value_of("config"));

    // get new filename
    let (filename, notes_path) = utils::get_new_filename(args.clone(), config.clone());
    if !notes_path.exists() {
        println!("Notes directory not found: {:?}", notes_path);
        println!("To make sure notes are not created in some random spot, the notes directory must already exist. Please create or change 'notes_dir' config in jot.conf to an existing directory");
        std::process::exit(1);
    }
    let file_path = notes_path.join(filename);

    // get file content from pipe
    if utils::is_pipe() {
        let mut stdin = io::stdin();
        match stdin.read_to_string(&mut content) {
            Ok(_) => {}
            Err(e) => println!("Error reading stdin: {:?}", e),
        };
    }

    // get file content from command-line
    if content == "" {
        match args.values_of("content") {
            Some(msg) => {
                let v: Vec<&str> = msg.collect();
                content = v.join(" ");
            }
            _ => {}
        };
    }

    // no content open file in EDITOR
    if content == "" {
        let editor_cmd = match env::var("EDITOR") {
            Ok(val) => val,
            Err(_) => "vim".to_string(),
        };
        match file_path.to_str() {
            Some(s) => {
                Command::new(editor_cmd)
                    .arg(s)
                    .status()
                    .expect("Error editing in vim");
            }
            None => panic!("Error with file_path: {:?}", file_path),
        };
    }
    // file content exists create file
    else {
        let mut file = match fs::OpenOptions::new()
            .read(true)
            .append(true)
            .create(true)
            .open(&file_path)
        {
            Ok(file) => file,
            Err(e) => panic!("Error creating file. {}", e),
        };

        // We want to guarentee a line ending but not double up
        // So subsequent notes don't append to same line
        // 1. Remove last line ending (if exists)
        let content = content.strip_suffix("\n").unwrap_or(&content);
        // 2. Add back the line ending
        let content = content.to_owned() + "\n";

        match file.write_all(content.as_bytes()) {
            Ok(_) => println!("Note added to: {}", file_path.to_str().unwrap()),
            Err(e) => panic!("Error writing to file. {}", e),
        }
    }
}
