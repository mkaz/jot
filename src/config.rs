use serde::Deserialize;
use std::env;
use std::fs;
use std::path::Path;
use toml;

#[derive(Clone, Deserialize)]
pub struct Config {
    pub notesdir: String,
}

// determine config file:
//  - command-line
//  - environment variable
//  - XDG_CONFIG_HOME
//  - Default: $HOME/.config/zk.conf
pub fn get_config(filearg: Option<&str>) -> Config {
    let filename = determine_filename(filearg);

    // check filename exists
    // XDG_CONFIG_HOME
    let config_path = Path::new(&filename);
    if !config_path.exists() {
        println!("Config file not found. {:?}", config_path);
        //TODO: better error message
        std::process::exit(1);
    }

    // finally $HOME/.config/zk.conf
    let config_file = fs::read_to_string(filename).unwrap();
    return toml::from_str(&config_file).unwrap();
}

fn determine_filename(filearg: Option<&str>) -> String {
    // from command-line
    if let Some(f) = filearg {
        return f.to_string();
    }

    // check enviornment variable
    if let Ok(val) = env::var("ZK_CONFIG_FILE") {
        return val;
    }

    if let Ok(xdg_dir) = env::var("XDG_CONFIG_HOME") {
        // check that xdg dir/zk.conf exists
        let filepath = format!("{}/zk.conf", xdg_dir);
        // TODO: check valid
        return filepath;
    }

    if let Ok(home_dir) = env::var("HOME") {
        let filepath = format!("{}/.config/zk.conf", home_dir);
        //TODO: check valid
        return filepath;
    }

    return "zk.conf".to_string();
}
