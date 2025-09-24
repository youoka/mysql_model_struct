package mysql_model_struct

type ShopConfig struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigType int    `gorm:"column:type;type:int" json:"configType"`
	Values     string `gorm:"column:values;type:text" json:"values"`
}

func (c ShopConfig) TableName() string {
	return "shop_config"
}
