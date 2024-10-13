package core

import (
	"time"
)

// TODO: possibly need a DpCollection struct
// to handle paginated fetches
type DpSource interface {
	Fetch(start_date time.Time, end_date time.Time) ([]Dp, error)
}
