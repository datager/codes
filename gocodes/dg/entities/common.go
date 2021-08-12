package entities

import (
	"bytes"
	"codes/gocodes/dg/utils/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type IntArray []int

func (s *IntArray) FromDB(bts []byte) error {
	if len(bts) == 0 {
		return nil
	}

	str := string(bts)
	if strings.HasPrefix(str, "{") {
		str = "[" + str[1:]
	}

	if strings.HasSuffix(str, "}") {
		str = str[0:len(str)-1] + "]"
	}

	ia := &[]int{}
	err := json.Unmarshal([]byte(str), ia)
	if err != nil {
		return errors.WithStack(err)
	}

	*s = IntArray(*ia)
	return nil
}

func (s *IntArray) ToDB() ([]byte, error) {
	return serializeIntArray(*s, "{", "}"), nil
}

func (s *IntArray) MarshalJSON() ([]byte, error) {
	return serializeIntArrayAsString(*s, "[", "]"), nil
}

func (s *IntArray) UnmarshalJSON(b []byte) error {
	var strarr []string
	var intarr []int
	err := json.Unmarshal(b, &strarr)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, v := range strarr {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		intarr = append(intarr, int(i))
	}

	*s = intarr
	return nil
}

func serializeIntArray(s []int, prefix string, suffix string) []byte {
	var buffer bytes.Buffer
	_, _ = buffer.WriteString(prefix)
	for idx, val := range s {
		if idx > 0 {
			_, _ = buffer.WriteString(",")
		}

		_, _ = buffer.WriteString(strconv.FormatInt(int64(val), 10))
	}
	_, _ = buffer.WriteString(suffix)

	return buffer.Bytes()
}

func serializeIntArrayAsString(s []int, prefix string, suffix string) []byte {
	var buffer bytes.Buffer
	_, _ = buffer.WriteString(prefix)

	for idx, val := range s {
		if idx > 0 {
			_, _ = buffer.WriteString(",")
		}
		_, _ = buffer.WriteString("\"")
		_, _ = buffer.WriteString(strconv.FormatInt(int64(val), 10))
		_, _ = buffer.WriteString("\"")
	}

	_, _ = buffer.WriteString(suffix)
	return buffer.Bytes()
}

type StringArray []string

func (s *StringArray) FromDB(bts []byte) error {
	if len(bts) == 0 {
		return nil
	}

	str := string(bts)
	if strings.HasPrefix(str, "{") {
		str = "[" + str[1:]
	}

	if strings.HasSuffix(str, "}") {
		str = str[0:len(str)-1] + "]"
	}

	ia := &[]string{}
	err := json.Unmarshal([]byte(str), ia)
	if err != nil {
		return errors.WithStack(err)
	}

	*s = StringArray(*ia)
	return nil
}

func (s *StringArray) ToDB() ([]byte, error) {
	return serializeStringArray(*s, "{", "}"), nil
}

func (s *StringArray) MarshalJSON() ([]byte, error) {
	return serializeStringArrayAsString(*s, "[", "]"), nil
}

func (s *StringArray) UnmarshalJSON(b []byte) error {
	var strarr []string
	err := json.Unmarshal(b, &strarr)
	if err != nil {
		return errors.WithStack(err)
	}

	*s = strarr
	return nil
}

func serializeStringArray(s []string, prefix string, suffix string) []byte {
	var buffer bytes.Buffer
	_, _ = buffer.WriteString(prefix)
	for idx, val := range s {
		if idx > 0 {
			_, _ = buffer.WriteString(",")
		}

		_, _ = buffer.WriteString(val)
	}
	_, _ = buffer.WriteString(suffix)

	return buffer.Bytes()
}

func serializeStringArrayAsString(s []string, prefix string, suffix string) []byte {
	var buffer bytes.Buffer
	_, _ = buffer.WriteString(prefix)

	for idx, val := range s {
		if idx > 0 {
			_, _ = buffer.WriteString(",")
		}
		_, _ = buffer.WriteString("\"")
		_, _ = buffer.WriteString(val)
		_, _ = buffer.WriteString("\"")
	}

	_, _ = buffer.WriteString(suffix)
	return buffer.Bytes()
}
