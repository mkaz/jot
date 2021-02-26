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

    if args.is_present("monthly") {
        config.
    }
    now.format(&format).to_string()
}
