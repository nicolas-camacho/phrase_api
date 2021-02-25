package phrases

import "time"

//Phrase struct defines the database model for a Phrase
type Phrase struct {
	ID        uint      `json:"id" gorm:"primary_key;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content" gorm:"unique;not null;default:null"`
}
