package entity

//Book represents book table in database
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"desc"`
	UserId      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"user"`
}
