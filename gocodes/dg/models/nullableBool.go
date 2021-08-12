package models

const (
	nullableBool_Int_True  = 1
	nullableBool_Int_False = 2
)

type Nullable interface {
	Null() bool
}

type NullableBool interface {
	Nullable
	True() bool
}

func NewNullableBoolByInt(val int) NullableBool {
	return &nullableBool_Int{val}
}

type nullableBool_Int struct {
	Value int
}

func (this *nullableBool_Int) Null() bool {
	return this.Value != nullableBool_Int_True && this.Value != nullableBool_Int_False
}

func (this *nullableBool_Int) True() bool {
	return this.Value == nullableBool_Int_True
}
