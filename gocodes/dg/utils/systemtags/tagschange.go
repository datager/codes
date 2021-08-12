package systemtags

import (
	"fmt"
	"codes/gocodes/dg/utils/config"
	"strconv"
	"strings"
)

func GetTags() (map[string]string, error) {

	tagsMap := config.GetConfig().GetStringMapString("tags.person")
	if len(tagsMap) == 0 {
		return nil, fmt.Errorf("config read err")
	}
	return tagsMap, nil
}

func TagsChange(tag []string) (int, error) {
	tagsMap, err := GetTags()
	if err != nil {
		return 0, err
	}
	var sum int
	for _, value := range tag {
		if v, ok := tagsMap[value]; ok {
			In, _ := strconv.Atoi(v)
			sum += In
		} else {
			return 0, fmt.Errorf("system input err")
		}
	}
	return sum, nil
}

func TagsMerge(tag1, tag2 []string) string {
	var s string
	tag := append(tag1, tag2...)
	for i := 0; i < len(tag); i++ {
		if i == len(tag)-1 {
			s += tag[i]
			break
		}
		s += tag[i] + ","
	}
	return s
}

func TagsSplit(tags string) ([]string, []string) {
	allTags := strings.Split(tags, ",")
	tagsMap, err := GetTags()
	if err != nil {
		return nil, nil
	}
	var systemTags, customTags []string
	for _, value := range allTags {
		if value == "" {
			continue
		} else if _, ok := tagsMap[value]; ok {
			systemTags = append(systemTags, value)
		} else {
			customTags = append(customTags, value)
		}

	}
	return systemTags, customTags
}
