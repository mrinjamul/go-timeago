/*Package timeago ...
timeago will evaluate time and return that time in words.

Example: an hour ago, a moment ago etc.
*/
package timeago

import (
	"errors"
	"strconv"
	"time"
)

// FormatNow takes previous time and return that time in words.
func FormatNow(past time.Time) (string, error) {
	// If time is zero, return an error with a message.
	if past.IsZero() {
		return "", errors.New("no date has been set")
	}
	now := time.Now()
	msg, err := Format(now, past)
	if err != nil {
		return "", err
	}
	return msg, nil
}

// Format takes current and previous time and return that time in words.
func Format(now, past time.Time) (string, error) {
	// If time is zero, return an error with a message.
	if now.IsZero() || past.IsZero() {
		return "", errors.New("no date has been set")
	}
	msg := ""
	seconds := toSeconds(now) - toSeconds(past)
	minutes := toMinutes(now) - toMinutes(past)
	hours := toHours(now) - toHours(past)
	days := toDays(now) - toDays(past)
	// create absolute values
	seconds = abs(seconds)
	minutes = abs(minutes)
	hours = abs(hours)
	days = abs(days)

	if days > 0 && days < 30 {
		if days > 1 {
			msg = strconv.Itoa(int(days)) + " days ago"
		} else {
			msg = "yesterday"
		}
	} else if days >= 30 && days < 365 {
		months := days / 30
		if months > 1 {
			msg = strconv.Itoa(int(months)) + " months ago"
		} else {
			msg = "a month ago"
		}
	} else if days >= 365 {
		years := days / 365
		if years > 1 {
			msg = strconv.Itoa(int(years)) + " years ago"
		} else {
			msg = "a year ago"
		}
	} else if days == 0 && hours > 0 {
		if hours > 1 {
			msg = strconv.Itoa(int(hours)) + " hours ago"
		} else {
			msg = "an hour ago"
		}
	} else if hours == 0 && minutes > 0 {
		if minutes > 1 {
			msg = strconv.Itoa(int(minutes)) + " minutes ago"
		} else {
			msg = "a minute ago"
		}
	} else if minutes == 0 && seconds > 0 {
		if seconds > 1 {
			msg = strconv.Itoa(int(seconds)) + " seconds ago"
		} else {
			msg = "a second ago"
		}
	} else {
		msg = "a moment ago"

	}
	return msg, nil
}

func toSeconds(t time.Time) int64 {
	seconds := t.Unix()
	return seconds
}
func toMinutes(t time.Time) int64 {
	seconds := toSeconds(t)
	minutes := seconds / 60
	return minutes
}

func toHours(t time.Time) int64 {
	minutes := toMinutes(t)
	hours := minutes / 60
	return hours
}

func toDays(t time.Time) int64 {
	hours := toHours(t)
	days := hours / 24
	return days
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
