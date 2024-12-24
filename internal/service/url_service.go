package service

import (
	"context"
	"fmt"
	"url-shortener-api/internal/cache"
	"url-shortener-api/internal/entity"
	"url-shortener-api/internal/repository"
)

type URLService struct {
	repository *repository.URLRepository
	cache      *cache.Cache
}

func (s *URLService) Shorten(ctx context.Context, url string) (string, error) {
	urlEntity := entity.NewURL(url)
	err := s.repository.Save(ctx, urlEntity)
	if err != nil {
		return "", err
	}
	return urlEntity.Hash, nil
}

func (s *URLService) GetTargetLink(ctx context.Context, hash string) (string, error) {
	url, err := s.cache.Get(ctx, hash)
	if err == nil {
		return url, nil
	}
	urlEntity, err := s.repository.FindByHash(ctx, hash)
	if err != nil {
		return "", err
	}
	err = s.cache.Set(ctx, hash, urlEntity.TargetLink)
	if err != nil {
		fmt.Printf("error setting cache: %v", err)
	}
	return urlEntity.TargetLink, nil
}

func NewURLService(repository *repository.URLRepository, cache *cache.Cache) *URLService {
	return &URLService{repository: repository, cache: cache}
}
