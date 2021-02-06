// extern crate toml;

use clap::{App, Arg};
use std::fs::File;
use std::path::Path;

use std::io::prelude::*;
use std::io::{self, Read};

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
        .arg(Arg::new("message").multiple(true))
        .get_matches();

    let notes_path = Path::new("/home/mkaz/Documents/Notes/Zk");
    // let config_data = fs::read_to_string(config_file).unwrap();
    // let config = config_data.parse::<toml::Value>().unwrap();

    // No View
    // Create New File
    let mut content = String::new();

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
        match matches.values_of("message") {
            Some(msg) => {
                let v: Vec<&str> = msg.collect();
                content = v.join(" ");
            }
            _ => {}
        };
    }

    // get new filename
    if content != "" {
        let filename = utils::get_new_filename();
        let file_path = notes_path.join(filename);

        if file_path.exists() {
            panic!("File already exists");
        }

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
