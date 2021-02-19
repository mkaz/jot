// extern crate toml;

use clap::{App, Arg};
use std::fs::File;
use std::io::prelude::*;
use std::io::{self, Read};
use std::path::Path;
use std::process::Command;

mod config;
mod utils;

fn main() {
    let matches = App::new("zk")
        .version("0.1.0")
        .about("Zettlekasten on the command-line")
        .author("Marcus Kazmierczak")
        .arg(
            Arg::new("length")
                .about("Number of words")
                .short('l')
                .long("length")
                .takes_value(true),
        )
        .arg(
            Arg::new("config")
                .about("Configuration file")
                .short('c')
                .long("config"),
        )
        .arg(
            Arg::new("content")
                .about("Create note from command-line")
                .multiple(true),
        )
        .after_help(
            "Create notes:
    1. Pipe into zk
        echo 'Hello' | zk

    2. Command-line args:
        zk 'This is my note'

    3. No args or pipe opens new note in $EDITOR",
        )
        .get_matches();

    // read in config
    let config = config::get_config("zk.conf");

    let notes_path = Path::new(&config.notes_path);
    // let config_data = fs::read_to_string(config_file).unwrap();
    // let config = config_data.parse::<toml::Value>().unwrap();

    // No View
    // Creating New File
    let mut content = String::new();

    // get new filename
    // TODO: append to existing file
    let filename = utils::get_new_filename();
    let file_path = notes_path.join(filename);
    if file_path.exists() {
        panic!("File already exists");
    }

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
        match matches.values_of("content") {
            Some(msg) => {
                let v: Vec<&str> = msg.collect();
                content = v.join(" ");
            }
            _ => {}
        };
    }

    // no content open file in EDITOR
    if content == "" {
        // TODO: get editor command
        match file_path.to_str() {
            Some(s) => {
                Command::new("vim")
                    .arg(s)
                    .status()
                    .expect("Error editing in vim");
            }
            None => panic!("Error with file_path: {:?}", file_path),
        };
    }
    // file content exists create file
    else {
        let mut file = match File::create(&file_path) {
            Ok(file) => file,
            Err(e) => panic!("Error creating file. {}", e),
        };

        match file.write_all(content.as_bytes()) {
            Ok(_) => println!("File created: {:?}", file_path),
            Err(e) => panic!("Error writing to file. {}", e),
        }
    }
}
