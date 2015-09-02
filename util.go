package palantir

import (
	"time"
)

func (r *Registration) HasExpired(tryDurationDays int) bool {
	duration := int64(tryDurationDays) * 24 * 3600
	return r.Date+duration < time.Now().Unix()
}
