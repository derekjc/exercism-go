package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var count int;
	for _, v := range birdsPerDay {
        count += v
    }
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var count int
	start := (week - 1) * 7
    for i := start; i < start + 7; i++ {
        count += birdsPerDay[i]
    }
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range(birdsPerDay) {
        if i % 2 == 0 {
            birdsPerDay[i] = birdsPerDay[i] + 1
        }
    }
	return birdsPerDay
}
