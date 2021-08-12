package models

type DictResponse struct {
	Code     int
	Message  string
	Redirect string
	Data     interface{}
}

type VseResponse struct {
	Context string `json:"Context"`
}

type VseAttribute struct {
	Name  VseAttributeName             `json:"name"`
	Value map[string]VseAttributeValue `json:"value"`
}

type VseAttributeName struct {
	Chinese string `json:"chs"`
	English string `json:"eng"`
}

type VseAttributeValue struct {
	Name VseAttributeName `json:"name"`
}

type VseBrandResponse struct {
	Brand map[string]VseBrandAttribute `json:"brand"`
}

type VseBrandAttribute struct {
	Name     VseAttributeName       `json:"name"`
	Subbrand map[string]VseSubbrand `json:"subbrand"`
}

type VseSubbrand struct {
	Name VseAttributeName             `json:"name"`
	Year map[string]VseAttributeValue `json:"year"`
}
