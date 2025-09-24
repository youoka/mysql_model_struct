package mysql_model_struct

type ShopConfig struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigType int    `gorm:"column:type;type:int" json:"configType"` // 1:系统配置 2:支付配置
	Values     string `gorm:"column:values;type:text" json:"values"`
}

func (c ShopConfig) TableName() string {
	return "shop_config"
}
