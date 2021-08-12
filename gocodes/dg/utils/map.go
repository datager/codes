package utils

type StringIntMap map[string]int

func (m StringIntMap) GetKeys() []string {
	s := make([]string, len(m))
	if len(m) == 0 {
		return s
	}
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	return s
}

type StringBoolMap map[string]bool

func (m StringBoolMap) GetKeys() []string {
	s := make([]string, len(m))
	if len(m) == 0 {
		return s
	}
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	return s
}

func AppendResultMapErrors(map1 map[string]error, map2 map[string]error) map[string]error {
	m := make(map[string]error)
	for k, m1 := range map1 {
		m[k] = m1
	}
	for k, m2 := range map2 {
		m[k] = m2
	}
	return m
}
