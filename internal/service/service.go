package service

import (
	"github.com/Thunderbirrd/ShortLinksService/internal/repository"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/Thunderbirrd/ShortLinksService/pkg/utils"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateShortUrl(urlObject models.UrlObject) (string, error) {
	shortUrl, err := s.repo.CheckLongUrl(urlObject.LongUrl)
	if err != nil {
		return "", err
	}

	if shortUrl != "" {
		return shortUrl, err
	}

	urlObject.ShortUrl = utils.GenerateShortUrl(urlObject.LongUrl)
	err = s.repo.SaveNewUrl(urlObject)

	if err != nil {
		return "", err
	}

	return urlObject.ShortUrl, nil
}
