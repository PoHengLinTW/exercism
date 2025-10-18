use std::fmt::Display;

#[derive(PartialEq, Debug)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Display for Clock {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}

const TWENTY_FOUR_HOURS_IN_MINUTES: i32 = 24 * 60;

fn adjust_clock(hours: i32, minutes: i32) -> Clock {
    let total_minutes = ((hours * 60 + minutes) % TWENTY_FOUR_HOURS_IN_MINUTES
        + TWENTY_FOUR_HOURS_IN_MINUTES)
        % TWENTY_FOUR_HOURS_IN_MINUTES;

    Clock {
        hours: total_minutes / 60,
        minutes: total_minutes % 60,
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        adjust_clock(hours, minutes)
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        adjust_clock(self.hours, self.minutes + minutes)
    }
}
