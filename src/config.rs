use serde::Deserialize;
use std::fs;
use toml;

#[derive(Clone, Deserialize)]
pub struct Config {
    pub notesdir: String,
}

pub fn get_config(filename: &str) -> Config {
    let config_file = fs::read_to_string(filename).unwrap();
    return toml::from_str(&config_file).unwrap();
}
