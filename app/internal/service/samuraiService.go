package service

type SamuraiRights interface {
	SetTurnover() error
	BindToDamiyo() error
}

type SamuraiService struct {
	relationService RelationsServiceMethods
	cardsService    CardRights
}

func NewSamuraiService(relationService RelationsServiceMethods, cardsService CardRights) *SamuraiService {
	return &SamuraiService{relationService: relationService, cardsService: cardsService}
}

func (s *SamuraiService) SetTurnover() error {
	return nil
}

func (s *SamuraiService) BindToDamiyo() error {
	return nil
}
