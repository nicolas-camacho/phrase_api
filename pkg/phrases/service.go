package phrases

// PhraseService defines author service behavior.
type PhraseService interface {
	CreatePhrase(*Phrase) (*Phrase, error)
	ReadPhrase(id int) (*Phrase, error)
	ObtainPhrase() (*Phrase, error)
}

// Service struct handles phrase business logic tasks.
type Service struct {
	repository PhraseRepository
}

//CreatePhrase is used to create a phrase in the repository
func (service *Service) CreatePhrase(phrase *Phrase) (*Phrase, error) {
	return service.repository.CreatePhrase(phrase)
}

//ReadPhrase is used to get a single phrase by id from the repository
func (service *Service) ReadPhrase(id int) (*Phrase, error) {
	return service.repository.ReadPhrase(id)
}

//ObtainPhrase is used to get a singles random phrase from the repository
func (service *Service) ObtainPhrase() (*Phrase, error) {
	return service.repository.ObtainPhrase()
}

//NewService is used to create a single instance of the service
func NewService(repository PhraseRepository) *Service {
	return &Service{repository: repository}
}
