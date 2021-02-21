use atty::Stream;
use chrono::{DateTime, Utc};

pub fn is_pipe() -> bool {
    !atty::is(Stream::Stdin)
}

// generate filename from date/time
pub fn get_new_filename(format: String) -> String {
    let now: DateTime<Utc> = Utc::now();
    now.format(&format).to_string()
}
