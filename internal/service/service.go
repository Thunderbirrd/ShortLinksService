package service

import (
	"github.com/Thunderbirrd/ShortLinksService/internal/repository"
	"github.com/Thunderbirrd/ShortLinksService/pkg/models"
	"github.com/Thunderbirrd/ShortLinksService/pkg/utils"
	"github.com/sirupsen/logrus"
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
		logrus.Errorf("Error while getting short url from db: %s", err.Error())
		return "", err
	}

	if shortUrl != "" {
		return shortUrl, err
	}

	urlObject.ShortUrl = utils.GenerateShortUrl(urlObject.LongUrl)
	err = s.repo.SaveNewUrl(urlObject.LongUrl, urlObject.ShortUrl)

	if err != nil {
		logrus.Errorf("Error while saving urls in db: %s", err.Error())
		return "", err
	}

	return urlObject.ShortUrl, nil
}

func (s *Service) GetLongUrlByShortUrl(shortUrl string) (string, error) {
	longUrl, err := s.repo.GetLongUrlByShortUrl(shortUrl)

	if err != nil {
		logrus.Errorf("Error while getting long url from db: %s", err.Error())
		return "", err
	}

	return longUrl, nil
}
