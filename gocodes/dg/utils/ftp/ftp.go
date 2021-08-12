package ftp

import (
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/pkg/errors"
)

type ConnOpts struct {
	Host     string
	Port     int
	UserName string
	Password string
	Timeout  int
}

func (opts *ConnOpts) getAddr() string {
	port := opts.Port
	if port == 0 {
		port = 21
	}
	return fmt.Sprintf("%v:%v", opts.Host, opts.Port)
}

type Params struct {
	*ConnOpts
	Path string
}

type ListParams struct {
	*Params
	DirOnly bool
}

type RmDirParams struct {
	*Params
	Recursive bool
}

type RenameParams struct {
	*Params
	From, To string
}

func List(params *ListParams) ([]*ftp.Entry, error) {
	conn, err := DialAndLogin(params.ConnOpts)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer conn.Quit()
	entries, err := conn.List(params.Path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(entries) == 0 {
		return []*ftp.Entry{}, nil
	}
	if !params.DirOnly {
		return entries, nil
	}
	dirs := make([]*ftp.Entry, 0)
	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFolder {
			dirs = append(dirs, entry)
		}
	}
	return dirs, nil
}

func MkDir(params *Params) error {
	conn, err := DialAndLogin(params.ConnOpts)
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Quit()
	return conn.MakeDir(params.Path)
}

func RmDir(params *RmDirParams) error {
	conn, err := DialAndLogin(params.ConnOpts)
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Quit()
	if params.Recursive {
		return conn.RemoveDirRecur(params.Path)
	}
	return conn.RemoveDir(params.Path)
}

func Rename(params *RenameParams) error {
	conn, err := DialAndLogin(params.ConnOpts)
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Quit()
	return conn.Rename(params.From, params.To)
}

func DialAndLogin(opts *ConnOpts) (*ftp.ServerConn, error) {
	conn, err := ftp.DialTimeout(opts.getAddr(), time.Duration(opts.Timeout)*time.Millisecond)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return conn, conn.Login(opts.UserName, opts.Password)
}
