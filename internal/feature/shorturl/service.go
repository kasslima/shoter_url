package shorturl

import "errors"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateShortURL(url string) (string, error) {
	if len(url) < 5 {
		return "", errors.New("Url is too short")
	}
	return url[:5], nil
}
