use atty::Stream;
use chrono::{DateTime, Utc};
use clap::ArgMatches;

pub fn is_pipe() -> bool {
    !atty::is(Stream::Stdin)
}

// generate filename from date/time
pub fn get_new_filename(args: ArgMatches, config: super::config::Config) -> String {
    // determine if monthly, weekly, daily, new or default
    let now: DateTime<Utc> = Utc::now();

    let format = if args.is_present("monthly") {
        config.monthly_format
    } else if args.is_present("weekly") {
        config.weekly_format
    } else if args.is_present("daily") {
        config.daily_format
    } else if args.is_present("new") {
        config.new_format
    } else {
        config.default_format
    };

    now.format(&format).to_string()
}
