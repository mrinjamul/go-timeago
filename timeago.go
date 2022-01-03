/*
Package timeago ...
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
	// If time is in future, return an error with a message.
	if time.Now().Before(past) {
		return "", errors.New("date is in future")
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
	if past.IsZero() {
		return "", errors.New("no date has been set")
	}
	// If time is in future, return an error with a message.
	if now.Before(past) {
		return "", errors.New("date is in future")
	}
	// If time is in past, return time in words.
	// Get the difference between the two times.
	diff := now.Sub(past)
	// If difference is less than a second, return a moment ago.
	if diff < time.Second {
		return "a moment ago", nil
	}
	// If difference is less than a minute, return seconds ago.
	if diff < time.Minute {
		if int(diff.Seconds()) == 1 {
			return "a second ago", nil
		}
		return strconv.Itoa(int(diff.Seconds())) + " seconds ago", nil
	}
	// If difference is less than an hour, return minutes ago.
	if diff < time.Hour {
		if int(diff.Minutes()) == 1 {
			return "a minute ago", nil
		}
		return strconv.Itoa(int(diff.Minutes())) + " minutes ago", nil
	}
	// If difference is less than a day, return hours ago.
	if diff < 24*time.Hour {
		if int(diff.Hours()) == 1 {
			return "an hour ago", nil
		}
		return strconv.Itoa(int(diff.Hours())) + " hours ago", nil
	}
	// If difference is less than a week, return days ago.
	if diff < 7*24*time.Hour {
		if int(diff.Hours()/24) == 1 {
			return "yesterday", nil
		}
		return strconv.Itoa(int(diff.Hours()/24)) + " days ago", nil
	}
	// If difference is less than a month, return weeks ago.
	if diff < 30*24*time.Hour {
		if int(diff.Hours()/24/7) == 1 {
			return "a week ago", nil
		}
		return strconv.Itoa(int(diff.Hours()/24/7)) + " weeks ago", nil
	}
	// If difference is less than a year, return months ago.
	if diff < 365*24*time.Hour {
		if int(diff.Hours()/24/30) == 1 {
			return "a month ago", nil
		}
		return strconv.Itoa(int(diff.Hours()/24/30)) + " months ago", nil
	}
	// If difference is more than a year, return years ago.
	if int(diff.Hours()/24/365) == 1 {
		return "a year ago", nil
	}
	return strconv.Itoa(int(diff.Hours()/24/365)) + " years ago", nil
}
