use time::Duration;
use time::PrimitiveDateTime as DateTime;

const ONE_GIGASECOND: Duration = Duration::seconds(1000 * 1000 * 1000);

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    // todo!("What time is a gigasecond later than {start}");
    start.saturating_add(ONE_GIGASECOND)
}
