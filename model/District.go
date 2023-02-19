package model

// 省市区
type District struct {
	ID    uint   `gorm:"primarykey"`
	Pid   uint   `gorm:"type:uint;not null;comment:父id"`
	Name  string `gorm:"type:varchar(120);not null;comment:名称"`
	Level uint8  `gorm:"type:tinyint(1);not null;comment:层级"`
}

func (district *District) TableName() string {
	return "district"
}
