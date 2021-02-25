package phrases

// PhraseService defines author service behavior.
type PhraseService interface {
	CreatePhrase(*Phrase) (*Phrase, error)
	ReadPhrase(id int) (*Phrase, error)
	ReadPhrases() (*[]Phrase, error)
	ObtainPhrase() (*Phrase, error)
}

type service struct {
	repository PhraseRepository
}

//CreatePhrase is used to create a phrase in the repository
func (service *service) CreatePhrase(phrase *Phrase) (*Phrase, error) {
	return service.repository.CreatePhrase(phrase)
}

//ReadPhrases is used to get all the phrases from the repository
func (service *service) ReadPhrases() (*[]Phrase, error) {
	return service.repository.ReadPhrases()
}

//ReadPhrase is used to get a single phrase by id from the repository
func (service *service) ReadPhrase(id int) (*Phrase, error) {
	return service.repository.ReadPhrase(id)
}

//ObtainPhrase is used to get a singles random phrase from the repository
func (service *service) ObtainPhrase() (*Phrase, error) {
	return service.repository.ObtainPhrase()
}

//NewService is used to create a single instance of the service
func NewService(repository PhraseRepository) PhraseService {
	return &service{
		repository: repository,
	}
}
