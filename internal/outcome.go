package internal

import "time"

type Outcome struct {
	ID          uint
	Description string
	Amount      float32
	Date        time.Time
}
