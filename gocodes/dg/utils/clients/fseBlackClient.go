package clients

import (
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type FseBlackClient struct {
	FseCli
}

func NewFseBlackClient(addr string, timeout time.Duration) (*FseBlackClient, error) {
	client, err := NewFseClient(addr, timeout)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &FseBlackClient{
		*client,
	}, nil
}

func (fbc *FseBlackClient) QueryRepo(repoID string) bool {
	return fbc.FseCli.QueryRepo(repoID)
}

func (fbc *FseBlackClient) AddRepo(repo RepoInfo) error {
	return fbc.FseCli.AddRepo(repo)
}

func (fbc *FseBlackClient) DeleteRepo(repoID string) error {
	return fbc.FseCli.DeleteRepo(repoID)
}

func (fbc *FseBlackClient) AddFeature(repoID string, feature EntityObject) error {
	return fbc.FseCli.AddFeature(repoID, feature)
}

func (fbc *FseBlackClient) DeleteFeature(repoID string, featureID string) error {
	return fbc.FseCli.DeleteFeature(repoID, featureID)
}

func (fbc *FseBlackClient) RankFeature(query RepoSearchCriteria) ([]SearchResultItem, error) {
	return fbc.FseCli.RankFeature(query)
}

func (fbc *FseBlackClient) MutliRankFeature(query RepoSearchCriteria) ([]SearchResultItem, error) {
	return fbc.FseCli.MutliRankFeature(query)
}

func (fbc *FseBlackClient) BatchDeleteFeature(repoID string, featureIDs []string) error {
	var g errgroup.Group
	for _, v := range featureIDs {
		id := v
		g.Go(func() error {
			return fbc.DeleteFeature(repoID, id)
		})
	}

	if err := g.Wait(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
