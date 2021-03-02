use atty::Stream;
use chrono::{DateTime, Utc};
use clap::ArgMatches;
use std::path::{Path, PathBuf};

pub fn is_pipe() -> bool {
    !atty::is(Stream::Stdin)
}

// generate filename and path from type and config
pub fn get_new_filename(args: ArgMatches, config: super::config::Config) -> (String, PathBuf) {
    // determine if monthly, weekly, daily, new or default
    let now: DateTime<Utc> = Utc::now();

    let (format, path) = if args.is_present("monthly") {
        let path = if let Some(p) = config.monthly_path {
            p
        } else {
            config.default_path
        };
        (config.monthly_format, path)
    } else if args.is_present("weekly") {
        let path = if let Some(p) = config.weekly_path {
            p
        } else {
            config.default_path
        };
        (config.weekly_format, path)
    } else if args.is_present("daily") {
        let path = if let Some(p) = config.daily_path {
            p
        } else {
            config.default_path
        };
        (config.daily_format, path)
    } else if args.is_present("new") {
        let path = if let Some(p) = config.new_path {
            p
        } else {
            config.default_path
        };
        (config.new_format, path)
    } else {
        (config.default_format, config.default_path)
    };

    let path = Path::new(&path).to_path_buf();
    (now.format(&format).to_string(), path)
}
