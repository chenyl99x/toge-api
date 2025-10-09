package model

// Nation 国家模型
// @Description 国家信息

type Nation struct {
	ID   uint   `json:"id" gorm:"primaryKey" example:"1"`
	Name string `json:"name" gorm:"not null;size:100" example:"蒙德"`
}

func (Nation) TableName() string {
	return "nation"
}
