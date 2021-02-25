package phrases

import (
	"gorm.io/gorm"
)

//PhraseRepository provides an abstraction on top of the phrase data source
type PhraseRepository interface {
	CreatePhrase(*Phrase) (*Phrase, error)
	ReadPhrase(int) (*Phrase, error)
	ObtainPhrase() (*Phrase, error)
}

//Repository struct recive an instance of the database to used
type Repository struct {
	db *gorm.DB
}

//New is use to create an instance of the repository to interact with the database
func New(db *gorm.DB) *Repository {
	db.AutoMigrate(&Phrase{})

	return &Repository{
		db: db,
	}
}

//CreatePhrase is used to create a phrase inside the database
func (repository *Repository) CreatePhrase(phrase *Phrase) (*Phrase, error) {
	if err := repository.db.Create(&phrase).Error; err != nil {
		return nil, err
	}

	return phrase, nil
}

//ReadPhrase is used to get a phrase from the database by its ID
func (repository *Repository) ReadPhrase(id int) (*Phrase, error) {
	phrase := &Phrase{}

	if query := repository.db.Where("id = ?", id).First(phrase); query.Error != nil {
		return nil, query.Error
	}

	return phrase, nil
}

//ObtainPhrase is used to obtain a random phrase from the database
func (repository *Repository) ObtainPhrase() (*Phrase, error) {
	phrase := &Phrase{}

	if query := repository.db.Order("RAMDOM()").First(phrase); query.Error != nil {
		return nil, query.Error
	}

	return phrase, nil
}
