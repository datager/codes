package utils

import (
	"io"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func HTTPDownload(url, file string) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()
	f, err := os.Create(file)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return errors.WithStack(err)
}
