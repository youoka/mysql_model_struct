package mysql_model_struct

type Kf struct {
	Id  int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	URL string `gorm:"column:url;type:varchar(255)" json:"url"`
}

func (Kf) TableName() string {
	return "kf"
}
