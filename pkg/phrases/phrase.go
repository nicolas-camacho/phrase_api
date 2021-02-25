package phrases

import "time"

//Phrase struct defines the database model for a Phrase
type Phrase struct {
	ID        uint `gorm:"primary_key;autoIncrement"`
	CreatedAt time.Time
	Content   string
}
