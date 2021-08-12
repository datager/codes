package wolf

import (
	"bufio"
	"bytes"
	"github.com/golang/glog"
	"os"
)

func Urlownload() {
	path := "configs.getInstance().urlDownloadPath"

	file, err := os.Open(path)
	if err != nil {
		glog.Fatal(err)
	}

	reader := bufio.NewReader(file)
	buf := bytes.NewBuffer(make([]byte, 1024))
	buf.Reset()
	_ = reader
}
