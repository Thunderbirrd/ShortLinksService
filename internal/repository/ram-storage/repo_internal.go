package ramstorage

import "sync"

type RepositoryInternal struct {
	mux sync.RWMutex

	storage map[string]string
}

func NewRepositoryInternal(storage map[string]string) *RepositoryInternal {
	return &RepositoryInternal{storage: storage}
}

func (r *RepositoryInternal) SaveNewUrl(longUrl, shortUrl string) error {
	r.mux.Lock()

	r.storage[longUrl] = shortUrl

	r.mux.Unlock()

	return nil
}

func (r *RepositoryInternal) CheckLongUrl(longUrl string) (string, error) {
	return r.storage[longUrl], nil
}

func (r *RepositoryInternal) GetLongUrlByShortUrl(shortUrl string) (string, error) {
	r.mux.Lock()
	for k, v := range r.storage {
		if v == shortUrl {
			r.mux.Unlock()
			return k, nil
		}
	}
	r.mux.Unlock()

	return "", nil
}
