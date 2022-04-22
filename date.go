// https://gosamples.dev/date-format-yyyy-mm-dd/

package main

import (
	"fmt"
	"time"
)

const (
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDhhmmss = "2006-01-02 15:04:05"
)

func main() {
	now := time.Now().UTC()
	fmt.Println(now.Format(YYYYMMDD))
	fmt.Println(now.Format(YYYYMMDDhhmmss))
}
