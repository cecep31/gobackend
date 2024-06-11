package entities

type Tags struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
