package entity

type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func (User) TableName() string {
	return "users"
}
