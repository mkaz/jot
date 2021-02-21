use serde::Deserialize;
use std::env;
use std::fs;
use std::path::Path;
use toml;

#[derive(Clone, Deserialize)]
pub struct Config {
    pub notes_dir: String,
    #[serde(default = "default_format")]
    pub filename_format: String,
}

fn default_format() -> String {
    "%Y%m%d%H%M.md".to_string()
}

pub fn get_config(filearg: Option<&str>) -> Config {
    let filename = determine_filename(filearg);

    // check filename exists
    let config_path = Path::new(&filename);
    if !config_path.exists() {
        println!("Config file not found. {:?}", config_path);
        // TODO: better error message
        std::process::exit(1);
    }

    // Read and parse config
    let config_file = fs::read_to_string(filename).unwrap();
    return toml::from_str(&config_file).unwrap();
}

// Determine config file:
//   1. command-line parameter
//   2. ZK environment variable
//   3. XDG_CONFIG_HOME env variable
//   4. $HOME/.config/zk.conf
//   5. Default current dir
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
        // TODO: check valid
        return filepath;
    }

    // default
    return "zk.conf".to_string();
}
