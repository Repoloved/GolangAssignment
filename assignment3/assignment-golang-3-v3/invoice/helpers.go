package invoice

import (
	"strings"
	"time"
)

func formatDate(defaultDate string) string {

	tmp := strings.Split(defaultDate, "/")

	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	date, _ := time.Parse("2006-01-02", strings.Join(tmp, "-"))
	return date.Format("02-January-2006")
}
