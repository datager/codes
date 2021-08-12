package entities

type Kv struct {
	Key   string `xorm:"not null pk VARCHAR(8192)"`
	Value string `xorm:"VARCHAR(8192)"`
}
