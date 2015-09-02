package palantir

import (
	"time"
)

func (r *Registration) HasExpired() bool {
	duration := int64(r.TryDuration) * 24 * 3600
	return r.Date+duration < time.Now().Unix()
}
