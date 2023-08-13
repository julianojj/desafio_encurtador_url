package domain

import (
	cryptorand "crypto/rand"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type URLShortener struct {
	Id        string
	LongURL   string
	ShortURL  string
	ExpiresIn time.Time
}

func NewURLShortener(id string, expiresIn time.Time) (*URLShortener, error) {
	shortener := &URLShortener{
		Id:        id,
		ExpiresIn: expiresIn,
	}
	err := shortener.Validate()
	if err != nil {
		return nil, err
	}
	return shortener, nil
}

func (u *URLShortener) Shorten(url string) (string, error) {
	if url == "" {
		return "", errors.New("url is required")
	}
	u.LongURL = url
	shortenedURL := randomString(6)
	u.ShortURL = shortenedURL
	return shortenedURL, nil
}

func randomString(length int) string {
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	cryptorand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func (u *URLShortener) Validate() error {
	if u.IsExpired() {
		return errors.New("expired url")
	}
	return nil
}

func (u *URLShortener) IsExpired() bool {
	today := time.Now()
	return today.After(u.ExpiresIn)
}
