package org

import (
	"fmt"
	"testing"
)

func TestGetOrgCodesToBeAllocate(t *testing.T) {
	xx := GetFullOrgCodesToBeAllocate(1, 0)
	fmt.Println(xx)
}
