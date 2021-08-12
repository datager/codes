package clients

import (
	"time"

	"github.com/pkg/errors"
)

type FseRuntimeClient struct {
	FseCli
}

func NewFseRuntimeClient(addr string, timeout time.Duration) (*FseRuntimeClient, error) {
	client, err := NewFseClient(addr, timeout)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &FseRuntimeClient{
		*client,
	}, nil
}

func (frc *FseRuntimeClient) QueryRepo(repoID string) bool {
	return frc.FseCli.QueryRepo(repoID)
}

func (frc *FseRuntimeClient) AddReop(repo RepoInfo) error {
	return frc.FseCli.AddRepo(repo)
}

func (frc *FseRuntimeClient) RankFeature(query RepoSearchCriteria) ([]SearchResultItem, error) {
	return frc.FseCli.RankFeature(query)
}

func (frc *FseRuntimeClient) MutliRankFeature(query RepoSearchCriteria) ([]SearchResultItem, error) {
	return frc.FseCli.MutliRankFeature(query)
}
