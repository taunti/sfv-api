package cfn

import (
	"fmt"
	"time"
)

const dateFormat = "Jan.02.2006"

func ParseDate(date string) time.Time {
	parsed, err := time.Parse(dateFormat, date)
	if err != nil {
		fmt.Println(err)
	}
	return parsed
}

func DiffInDays(d1, d2 time.Time) float64 {
	return d2.Sub(d1).Hours() / 24
}
