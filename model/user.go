package mysql

type User struct {
	CreatedAt null.Time   `gorm:"column:created_at" json:"created_at"`
	ID        int         `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name      null.String `gorm:"column:name" json:"name"`
	Nickname  null.String `gorm:"column:nickname" json:"nickname"`
	Phone     null.String `gorm:"column:phone" json:"phone"`
	UpdatedAt null.Time   `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
