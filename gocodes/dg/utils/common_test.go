package utils

import "testing"

func TestAppendSlice(t *testing.T) {
	type Data struct {
		A int
		B string
	}
	//var s = []Data{{A: 1, B: "1"}}
	//s := []Data{{A: 1, B: "1"}}
	var s []Data
	//s := []Data{}
	s = append(s, Data{A: 2, B: "2"})
	t.Log(s)
}

func TestCopySlice1(t *testing.T) {
	type Data struct {
		A int
		B string
	}
	s := []Data{{A: 1, B: "1"}}
	s = append(s, Data{A: 2, B: "2"})

	var f = func(datas []Data) []Data {
		datas[0] = Data{A: 3, B: "3"}
		return datas
	}
	rets := f(s)
	t.Log(rets)
	t.Log(s)

	t.Log(s)
}

func TestCopySlice2(t *testing.T) {
	type Data struct {
		A int
		B string
	}
	s := []Data{{A: 1, B: "1"}}
	s = append(s, Data{A: 2, B: "2"})

	var f = func(datas []Data) []Data {
		datas[0] = Data{A: 3, B: "3"}
		datas = append(datas, Data{A: 4, B: "4"})
		datas = append(datas, Data{A: 5, B: "5"})
		datas = append(datas, Data{A: 6, B: "6"})
		datas = append(datas, Data{A: 7, B: "7"})
		return datas
	}
	rets := f(s)
	t.Log(rets)

	t.Log(s)
}

func TestInitMap(t *testing.T) {
	type Data struct {
		A int
		B string
	}
	//var m map[string]Data
	m := make(map[string]Data, 10)
	m["1"] = Data{A: 1, B: "2"}
	t.Log(m)
}

func TestSlice(t *testing.T) {
	type Data struct {
		A int
		B string
		C []Data
	}
	d := Data{A: 1}
	d.C = []Data{{A: 2, B: "2", C: []Data{{A: 3, B: "3"}}}}
	t.Log(d)
}
