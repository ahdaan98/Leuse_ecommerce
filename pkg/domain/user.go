package domain

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=20"`
	Phone    string `json:"phone" validate:"required"`
	Blocked  bool   `json:"blocked" gorm:"default:false"`
}

type Address struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	Users     User   `json:"-" gorm:"foreignkey:UserID"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}
