package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/begenov/tg-bot-golang/lib/e"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(UserName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavePages = errors.New("no seved page")

type Page struct {
	URL      string
	UserName string
}

const (
	ErrMsg = "can't calulcate hash"
)

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap(ErrMsg, err)
	}
	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap(ErrMsg, err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
