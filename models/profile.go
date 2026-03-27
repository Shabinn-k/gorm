package models
type Profile struct{
	ID uint `gorm:"primaryKey"`
	UserID uint 
	Bio string `json:"bio"`
	Age int `json:"age"`
}
