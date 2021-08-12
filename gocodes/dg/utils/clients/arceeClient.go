package clients

import (
	"io"

	"codes/gocodes/dg/models"
)

type ArceeClient interface {
	io.Closer
	BatchAddImages(imgs []*models.ImageQuery) ([]string, error)
	BatchDeleteImages(imageURLs []string) error
}
