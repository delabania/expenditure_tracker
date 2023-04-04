package expenditure_tracker

import (
	"time"
)

type Person struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Expenditure struct {
	ID          int64     `json:"id"`
	Person      Person    `json:"person"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	SplitRatio  float32   `json:"split_ratio"`
}

type Settlement struct {
	ID     int64     `json:"id"`
	Person Person    `json:"person"`
	Date   time.Time `json:"date"`
	Amount int64     `json:"amount"`
}
