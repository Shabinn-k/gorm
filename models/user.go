package models
type User struct{
	ID uint `gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}