package json

// http://jsoniter.com/go-tips.cn.html
import (
	"bytes"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// 临时忽略struct字段
func TestJsonIter1(t *testing.T) {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		// many more fields…
	}
	user := &User{
		Email:    "yc@dg.com",
		Password: "123",
	}
	b, _ := json.Marshal(struct {
		*User
		Password bool `json:"password,omitempty"`
	}{
		User: user,
		//Password: true,
	})
	t.Log(string(b))
}

// 临时添加额外的字段
func TestJsonIter2(t *testing.T) {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		// many more fields…
	}
	user := &User{
		Email:    "yc@dg.com",
		Password: "123",
	}
	b, _ := json.Marshal(struct {
		*User
		Token    string `json:"token"`
		Password bool   `json:"password,omitempty"`
	}{
		User:  user,
		Token: "i am token",
	})
	t.Log(string(b))
}

// 临时粘合两个struct
func TestJsonIter3(t *testing.T) {
	type BlogPost struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	}
	type Analytics struct {
		Visitors  int `json:"visitors"`
		PageViews int `json:"page_views"`
	}

	b, _ := json.Marshal(struct {
		*BlogPost
		*Analytics
	}{&BlogPost{
		URL:   "1",
		Title: "2",
	}, &Analytics{
		Visitors:  3,
		PageViews: 4,
	}})
	t.Log(string(b))
}

// 一个json切分成两个struct
func TestJsonIter4(t *testing.T) {
	type BlogPost struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	}
	type Analytics struct {
		Visitors  int `json:"visitors"`
		PageViews int `json:"page_views"`
	}

	post := &BlogPost{}
	analytics := &Analytics{}

	_ = json.Unmarshal([]byte(`{
  "url": "attila@attilaolah.eu",
  "title": "Attila's Blog",
  "visitors": 6,
  "page_views": 14
}`), &struct {
		*BlogPost
		*Analytics
	}{post, analytics})

	t.Log(post)
	t.Log(analytics)
}

func TestJsonIter5(t *testing.T) {
	//type CacheItem struct {
	//	Key    string `json:"key"`
	//	MaxAge int    `json:"cacheAge"`
	//	Value  Value  `json:"cacheValue"`
	//}
	//
	//json.Marshal(struct {
	//	*CacheItem
	//
	//	// Omit bad keys
	//	OmitMaxAge omit `json:"cacheAge,omitempty"`
	//	OmitValue  omit `json:"cacheValue,omitempty"`
	//
	//	// Add nice keys
	//	MaxAge int    `json:"max_age"`
	//	Value  *Value `json:"value"`
	//}{
	//	CacheItem: item,
	//
	//	// Set the int by value:
	//	MaxAge: item.MaxAge,
	//
	//	// Set the nested struct by reference, avoid making a copy:
	//	Value: &item.Value,
	//})
}

// 用字符串传递数字
func TestJsonIter6(t *testing.T) {
	type TestObject struct {
		Field1 int `json:",string"`
	}
	val := &TestObject{}
	jsoniter.Unmarshal([]byte(`{
"Field1": "123""
}`), val)
	t.Log(val)
	assert.Equal(t, 123, val.Field1)
}

// 容忍字符串和数字互转
func TestJsonIter7(t *testing.T) {
	extra.RegisterFuzzyDecoders()
	var val1 string
	jsoniter.UnmarshalFromString(`100`, &val1)
	var val2 float32
	jsoniter.UnmarshalFromString(`"1.23"`, &val2)
	t.Log(val1, val2)
}

// 容忍空数组作为对象
func TestJsonIter8(t *testing.T) {
	extra.RegisterFuzzyDecoders()
	var val map[string]interface{}
	jsoniter.UnmarshalFromString(`[]`, &val)
	t.Log(val)
}

// 使用 MarshalJSON支持time.Time
type timeImplementedMarshaler time.Time

//func (obj timeImplementedMarshaler) MarshalJSON() ([]byte, error) {
//	seconds := time.Time(obj).Unix()
//	return []byte(strconv.FormatInt(seconds, 10)), nil
//}
func TestJsonIter9(t *testing.T) {
	type TestObject struct {
		Field timeImplementedMarshaler
	}
	should := require.New(t)
	val := timeImplementedMarshaler(time.Unix(123, 0))
	obj := TestObject{val}
	bytes, err := jsoniter.Marshal(obj)
	should.Nil(err)
	t.Log(string(bytes))
	should.Equal(`{"Field":123}`, string(bytes))
}

// 使用 RegisterTypeEncoder支持time.Time
func TestJsonIter10(t *testing.T) {
	should := require.New(t)
	extra.RegisterTimeAsInt64Codec(time.Microsecond)
	output, err := jsoniter.Marshal(time.Unix(1, 1002))
	should.Nil(err)
	should.Equal("1000001", string(output))
}

// 使用 MarshalText支持非字符串作为key的map
func TestJsonIter11(t *testing.T) {
	//f, _, _ := big.ParseFloat("1", 10, 64, big.ToZero)
	//val := map[*big.Float]string{f: "2"}
	//str, err := MarshalToString(val)
	//should.Equal(`{"1":"2"}`, str)
}

// 使用 json.RawMessage
func TestJsonIter12(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 string
		Field2 json.RawMessage
	}
	var data TestObject
	json.Unmarshal([]byte(`{"field1": "hello", "field2": [1,2,3]}`), &data)
	should.Equal(`[1,2,3]`, string(data.Field2))
}

func TestJsonIter13(t *testing.T) {
	should := require.New(t)
	type TestObject struct {
		Field1 string
		Field2 []int64
	}
	var data TestObject
	json.Unmarshal([]byte(`{"field1": "hello", "field2": [1,2,3]}`), &data)
	should.Equal(`[1,2,3]`, data.Field2)
}

// 使用 json.Number
func TestJsonIter14(t *testing.T) {
	should := require.New(t)
	decoder1 := json.NewDecoder(bytes.NewBufferString(`123`))
	decoder1.UseNumber()
	var obj1 interface{}
	decoder1.Decode(&obj1)
	should.Equal(json.Number("123"), obj1)
}

func TestJsonIter15(t *testing.T) {
	should := require.New(t)
	json := jsoniter.Config{
		IndentionStep:                 0,
		MarshalFloatWith6Digits:       false,
		EscapeHTML:                    false,
		SortMapKeys:                   false,
		UseNumber:                     true,
		DisallowUnknownFields:         false,
		TagKey:                        "",
		OnlyTaggedField:               false,
		ValidateJsonRawMessage:        false,
		ObjectFieldMustBeSimpleString: false,
		CaseSensitive:                 false,
	}.Froze()
	var obj interface{}
	json.UnmarshalFromString("123", &obj)
	should.Equal(jsoniter.Number("123"), obj)
}

// 统一更改字段的命名风格
func TestJsonIter16(t *testing.T) {
	should := require.New(t)
	output, _ := jsoniter.Marshal(struct {
		UserName      string `json:"user_name"`
		FirstLanguage string `json:"first_language"`
	}{
		UserName:      "taowen",
		FirstLanguage: "Chinese",
	})
	should.Equal(`{"user_name":"taowen","first_language":"Chinese"}`, string(output))
}

func TestJsonIter17(t *testing.T) {
	//extra.SetNamingStrategy(jsoniter.LowerCaseWithUnderscores)
	//output, err := jsoniter.Marshal(struct {
	//	UserName      string
	//	FirstLanguage string
	//}{
	//	UserName:      "taowen",
	//	FirstLanguage: "Chinese",
	//})
	//should.Nil(err)
	//should.Equal(`{"user_name":"taowen","first_language":"Chinese"}`, string(output))
}

// 使用私有的字段
func TestJsonIter18(t *testing.T) {
	should := require.New(t)
	extra.SupportPrivateFields()
	type TestObject struct {
		field1 string
	}
	obj := TestObject{}
	jsoniter.UnmarshalFromString(`{"field1":"Hello"}`, &obj)
	should.Equal("Hello", obj.field1)
}
