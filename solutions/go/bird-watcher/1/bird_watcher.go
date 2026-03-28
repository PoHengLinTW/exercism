package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for _, cnt := range(birdsPerDay) {
        sum += cnt
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	days := len(birdsPerDay)
    sum := 0
    startDay := (week - 1) * 7
    endDay := min(days, week * 7)

    for startDay < endDay {
        sum += birdsPerDay[startDay]
        startDay += 1
    }

    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day := 0; day < len(birdsPerDay); day += 2 {
        birdsPerDay[day] += 1
    }
    return birdsPerDay
}
