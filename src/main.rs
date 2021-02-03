// extern crate toml;

use clap::{App, Arg};
// use std::fs;

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
            Arg::new("digits")
                .about("Include numbers")
                .short('d')
                .long("digits"),
        )
        .arg(
            Arg::new("stdin")
                .about("From stdin")
                .short('s')
                .long("stdin"),
        )
        .get_matches();

    let notes_dir = "/home/mkaz/Documents/Notes/Zk";
    // let config_data = fs::read_to_string(config_file).unwrap();
    // let config = config_data.parse::<toml::Value>().unwrap();

    // check for command line args
    // check for nothing
    // check for stdin

    if is_pipe() {
        let mut buffer = String::new();
        let mut stdin = io::stdin(); // We get `Stdin` here.
        match stdin.read_to_string(&mut buffer) {
            Ok(_) => println!("Read to buffer"),
            Err(_) => println!("An error happened"),
        };
        println!("Buffer: {}", buffer);
    } else {
        println!("Not piped");
    }

    println!("Done");
}

fn is_pipe() -> bool {
    !atty::is(Stream::Stdin)
}
