package utils

var Version = &VersionInfo{}

type VersionInfo struct {
	Version string
	Commit  string
	Date    string
}
