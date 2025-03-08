package entities

type Role struct {
	ID    uint      `gorm:"primaryKey;autoincrement"`
	Role  string    `gorm:"unique;not null"`
	Users []Account `gorm:"many2many:account_roles"`
}
