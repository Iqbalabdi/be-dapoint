package user

type User struct {
	ID         uint64 `gorm:"primaryKey"`
	Name       string `json:"name" form:"name"`
	Email      string `gorm:"unique" json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Photo      string
	Role       string `gorm:"default:user"`
	TotalPoint uint64 `json:"total_point" form:"total_point" gorm:"default:0"`
}
