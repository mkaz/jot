// extern crate toml;

use chrono::{DateTime, Utc};
use clap::{App, Arg};
use std::fs;
use std::path::Path;

use atty::Stream;
use std::io::{self, Read};

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
        .get_matches();

    let notes_path = Path::new("/home/mkaz/Documents/Notes/Zk");
    // let config_data = fs::read_to_string(config_file).unwrap();
    // let config = config_data.parse::<toml::Value>().unwrap();

    // check for command line args
    // check for nothing
    // check for stdin
    let mut content = String::new();

    if is_pipe() {
        let mut stdin = io::stdin(); // We get `Stdin` here.
        match stdin.read_to_string(&mut content) {
            Ok(_) => {}
            Err(e) => println!("Error reading stdin: {:?}", e),
        };
    } else {
        // read args from command-line
        println!("Not piped");
    }

    // get new filename
    if content != "" {
        let filename = get_new_filename();
        let file_path = notes_path.join(filename);
        println!("Creating file: {:?}", file_path);
    }
}

fn is_pipe() -> bool {
    !atty::is(Stream::Stdin)
}

// generate filename from date/time
fn get_new_filename() -> String {
    let now: DateTime<Utc> = Utc::now();
    now.format("%y%m%d%H%M.md").to_string()
}
