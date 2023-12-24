package utils

import (
	"time"

	"github.com/xeonx/timeago"
)

func GetTimeAgo() timeago.Config {
	c := timeago.Config{
		Max:          time.Duration(365*10) * 24 * time.Hour,
		PastPrefix:   "",
		PastSuffix:   " ago",
		FuturePrefix: "in ",
		FutureSuffix: "",

		Periods: []timeago.FormatPeriod{
			{D: time.Second, One: "about a second", Many: "%d seconds"},
			{D: time.Minute, One: "about a minute", Many: "%d minutes"},
			{D: time.Hour, One: "about an hour", Many: "%d hours"},
			{D: timeago.Day, One: "one day", Many: "%d days"},
			{D: timeago.Month, One: "one month", Many: "%d months"},
			{D: timeago.Year, One: "one year", Many: "%d years"},
		},
	}

	return c
}
