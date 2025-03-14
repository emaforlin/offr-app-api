package entities

type Role struct {
	ID   uint   `gorm:"primaryKey;autoincrement" json:"-"`
	Role string `gorm:"unique;not null" json:"role"`
}

type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}
