package timeago

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	testcases := []struct {
		name string
		msg  string
		now  time.Time
		past time.Time
	}{
		{"swapped time", "a year ago", time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"a year ago", "a year ago", time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC), time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"10 years ago", "10 years ago", time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC), time.Date(2011, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"a month ago", "a month ago", time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"10 months ago", "10 months ago", time.Date(2021, 10, 0, 0, 0, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"yesterday", "yesterday", time.Date(2021, 0, 1, 0, 0, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"10 days ago", "10 days ago", time.Date(2021, 0, 10, 0, 0, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"a hour ago", "a hour ago", time.Date(2021, 0, 0, 1, 0, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"10 hours ago", "10 hours ago", time.Date(2021, 0, 0, 10, 0, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"a minute ago", "a minute ago", time.Date(2021, 0, 0, 0, 1, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"10 minutes ago", "10 minutes ago", time.Date(2021, 0, 0, 0, 10, 0, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"a seconds ago", "a second ago", time.Date(2021, 0, 0, 0, 0, 1, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"10 seconds ago", "10 seconds ago", time.Date(2021, 0, 0, 0, 0, 10, 0, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"a moment ago", "a moment ago", time.Date(2021, 0, 0, 0, 0, 0, 50, time.UTC), time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"zero date", "", time.Time{}, time.Time{}},
	}
	for _, testcase := range testcases {
		msg, _ := Format(testcase.now, testcase.past)
		if msg != testcase.msg {
			t.Errorf("case: %v should be %v; got %v", testcase.name, testcase.msg, msg)
		}
	}
}

func TestFormatNow(t *testing.T) {
	testcases := []struct {
		name string
		msg  string
		past time.Time
	}{
		{"a year ago", "a year ago", time.Now().AddDate(-1, 0, 0)},
		{"10 year ago", "10 years ago", time.Now().AddDate(-10, 0, 0)},
		{"a month ago", "a month ago", time.Now().AddDate(0, -1, 0)},
		{"10 months ago", "10 months ago", time.Now().AddDate(0, -10, 0)},
		{"yesterday", "yesterday", time.Now().AddDate(0, 0, -1)},
		{"10 days ago", "10 days ago", time.Now().AddDate(0, 0, -10)},
		{"zero date", "", time.Time{}},
	}
	for _, testcase := range testcases {
		msg, _ := FormatNow(testcase.past)
		if msg != testcase.msg {
			t.Errorf("case: %v should be %v; got %v", testcase.name, testcase.msg, msg)
		}
	}
}
