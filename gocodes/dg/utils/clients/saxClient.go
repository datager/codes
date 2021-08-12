package clients

import (
	"context"
	"fmt"
	"time"

	"codes/gocodes/dg/utils/json"

	"golang.org/x/sync/errgroup"

	"github.com/emirpasic/gods/maps/treebidimap"
	godsutils "github.com/emirpasic/gods/utils"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type SaxClient struct {
	*HTTPClient
}

func NewSaxClient(baseAddr string, timeout time.Duration) *SaxClient {
	return &SaxClient{
		HTTPClient: NewHTTPClient(baseAddr, timeout, nil),
	}
}

func (this *SaxClient) BatchSearch(queries []*SaxQuery) (*treebidimap.Map, error) {
	results := make([][]*SaxCandidate, len(queries))
	g, _ := errgroup.WithContext(context.Background())
	for i, query := range queries {
		i, query := i, query
		g.Go(func() error {
			items, err := this.SearchEx(query)
			if err != nil {
				return errors.WithStack(err)
			}
			results[i] = items
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, errors.WithStack(err)
	}

	// sort value by score
	m := treebidimap.NewWith(godsutils.StringComparator, func(a, b interface{}) int {
		s1 := a.(*SaxCandidate).SimilarityDegree
		s2 := b.(*SaxCandidate).SimilarityDegree
		if s1 > s2 {
			return -1
		} else if s1 < s2 {
			return 1
		} else {
			return godsutils.StringComparator(a.(*SaxCandidate).Vid, b.(*SaxCandidate).Vid)
		}
	})
	for _, r := range results {
		for _, item := range r {
			if val, exist := m.Get(item.Vid); exist {
				if item.SimilarityDegree > val.(*SaxCandidate).SimilarityDegree {
					m.Put(item.Vid, item)
				}
			} else {
				m.Put(item.Vid, item)
			}
		}
	}
	return m, nil
}

func (this *SaxClient) SearchEx(query *SaxQuery) ([]*SaxCandidate, error) {
	req, err := newSaxSearchReq(query.Feature)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	items, err := this.Search(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	filtered := make([]*SaxCandidate, 0)
	for _, item := range items {
		if item.SimilarityDegree < query.Confidence {
			break
		}
		filtered = append(filtered, item)
	}
	return filtered, nil
}

func (this *SaxClient) Search(req *SaxSearchRequest) ([]*SaxCandidate, error) {
	resp, err := this.PostJSON("/api/face/clustersearch", req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result SaxSearchResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := validateSaxResponseContext(result.Context); err != nil {
		return nil, errors.WithStack(err)
	}
	return result.Candidates, nil
}

func (this *SaxClient) SearchVidCombines(req *SaxSearchVidCombinesRequest) ([]*SaxVidCombine, error) {
	resp, err := this.PostJSON("/api/face/vidcombines", req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result *SaxSearchVidCombinesResponse
	if err := json.Unmarshal(resp, result); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := validateSaxResponseContext(result.Context); err != nil {
		return nil, errors.WithStack(err)
	}
	return result.VidCombines, nil
}

func (this *SaxClient) SearchAllVidCombinesInGivenTime(startTimestamp int64, endTimestamp int64) ([]*SaxVidCombine, error) {

	vidCombinesMap := map[string]bool{}
	saxVidCombines := []*SaxVidCombine{}

	req := SaxSearchVidCombinesRequest{
		StartTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
		Offset:         0,
		Limit:          100,
	}

	for {

		resp, err := this.PostJSON("/api/face/vidcombines", req)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		result := &SaxSearchVidCombinesResponse{}
		if err := json.Unmarshal(resp, result); err != nil {
			return nil, errors.WithStack(err)
		}
		if err := validateSaxResponseContext(result.Context); err != nil {
			return nil, errors.WithStack(err)
		}

		addCount := 0
		for _, vidCombine := range result.VidCombines {
			if _, isExist := vidCombinesMap[vidCombine.SourceVid]; !isExist {
				vidCombinesMap[vidCombine.SourceVid] = true
				saxVidCombines = append(saxVidCombines, vidCombine)
				addCount++
			}
		}

		if addCount > 0 {
			req.Offset += 100
			continue
		}

		return saxVidCombines, nil

	}

}

func newSaxSearchReq(feature string) (*SaxSearchRequest, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &SaxSearchRequest{
		Context: &SaxRequestContext{
			SessionId: id.String(),
		},
		Feature:      feature,
		Manufacturer: SAX_MANUFACTURER_DEEPGLINT,
	}, nil
}

func validateSaxResponseContext(respCtx *SaxResponseContext) error {
	if respCtx == nil {
		return fmt.Errorf("Nil response context")
	}
	if respCtx.Status != SaxResponseContextStatusOK {
		return fmt.Errorf("Sax error %v: %v", respCtx.Status, respCtx.Message)
	}
	return nil
}
