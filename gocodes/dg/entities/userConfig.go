package entities

type UserConfig struct {
	UserID string `xorm:"user_id not null VARCHAR(1024)"`
	Key    string `xorm:"key not null VARCHAR(1024)"`
	Config string `xorm:"config not null VARCHAR(1024)"`
}

func (UserConfig) TableName() string {
	return TableNameUserConfig
}
