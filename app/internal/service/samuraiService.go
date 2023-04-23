package service

type SamuraiRights interface {
	SetTurnover() error
	BindToDamiyo() error
}

type SamuraiService struct {
}

func (s *SamuraiService) SetTurnover() error {
	return nil
}

func (s *SamuraiService) BindToDamiyo() error {
	return nil
}
