package entities

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
