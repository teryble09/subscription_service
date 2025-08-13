package dateparse

import "time"

const layout = "01-2006"

func ParseMMYYYY(date string) (time.Time, error) {
	// in case if user wants to set end date to null
	if date == "" {
		return time.Time{}, nil
	}
	return time.Parse(layout, date)
}

func ParseIntoMMYYYY(date time.Time) string {
	return date.Format(layout)
}
