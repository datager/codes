package utils

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

type FileInfo struct {
	Name    string      // base name of the file
	Size    int64       // length in bytes for regular files; system-dependent for others
	Mode    os.FileMode // file mode bits
	ModTime time.Time   // modification time
	IsDir   bool        // abbreviation for Mode.IsDir
}

func FileInfoToStruct(fi os.FileInfo) *FileInfo {
	return &FileInfo{
		Name:    fi.Name(),
		Size:    fi.Size(),
		Mode:    fi.Mode(),
		ModTime: fi.ModTime(),
		IsDir:   fi.IsDir(),
	}
}

func Stat(name string) (*FileInfo, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return FileInfoToStruct(fi), errors.WithStack(err)
}

func StatDir(name string) ([]*FileInfo, error) {
	fis, err := ioutil.ReadDir(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	results := make([]*FileInfo, len(fis))
	for i, fi := range fis {
		results[i] = FileInfoToStruct(fi)
	}
	return results, nil
}
