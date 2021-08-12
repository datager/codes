package models

import (
	"fmt"
	"codes/gocodes/dg/utils/config"
	"codes/gocodes/dg/utils/math"
	"codes/gocodes/dg/utils/slice"
)

const (
	RootOrgLevel = 0

	// root org code: 0; not null default org code in sql: -1; invalid org code in code: -2;
	RootOrgCode    = 0
	InValidOrgCode = -2
)

var (
	MaxOrgNumPerLevel int64
	MaxOrgLevel       int64
	MaxOrgNum         int64
)

func LoadOrgConfVal() {
	MaxOrgNumPerLevel = int64(config.GetConfig().GetIntOrDefault("api.org.max_org_num_per_level", 500))
	MaxOrgLevel = int64(config.GetConfig().GetIntOrDefault("api.org.max_org_level", 6))
	MaxOrgNum = math.PowInt64(MaxOrgNumPerLevel, MaxOrgLevel)
}

//---------------allocate new one---------------
func GetOrgCodesToBeAllocate(superOrgLevel, superOrgCode int64) []int64 {
	var orgCodes []int64
	for i := 1; i < int(MaxOrgNumPerLevel); i++ {
		xx := math.PowInt64(MaxOrgNumPerLevel, superOrgLevel)
		yy := MaxOrgNum / xx
		zz := int64(i) * yy
		ww := superOrgCode + zz
		_ = ww
		orgCodes = append(orgCodes, superOrgCode+int64(i)*(MaxOrgNum/(math.PowInt64(MaxOrgNumPerLevel, superOrgLevel))))
	}
	return orgCodes
}

func CheckAndPickNewOrgCode(tobeAllocate, allocated []int64) (int64, error) {
	remainedOrgCode := slice.GetDiffIntSet(tobeAllocate, allocated)
	if len(remainedOrgCode) == 0 {
		fmt.Print("err org code")
		return InValidOrgCode, fmt.Errorf("[OrgCode] no org code remaining, len(allocated):%v, len(tobeAllocate):%v, allocated:%v\n, tobeAllocate:%v", len(allocated), len(tobeAllocate), allocated, tobeAllocate)
	}
	return remainedOrgCode[0], nil // random pick one
}

//---------------capture search---------------
// level1 not need org code condition
type OrgCodeRange struct {
	IsNeedInSQL bool  // if root org(level 0): false
	MinOrgCode  int64 // >=
	MaxOrgCode  int64 // <
}
