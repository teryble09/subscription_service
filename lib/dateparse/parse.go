package dateparse

import "time"

// date format in the techical specification
const layout = "01-2006"

func ParseMMYYYY(date string) (time.Time, error) {
	return time.Parse(layout, date)
}

func ParseIntoMMYYYY(date time.Time) string {
	return date.Format(layout)
}
