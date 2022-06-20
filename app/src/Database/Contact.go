package Database

import "time"

const ContactTableName = string("contacts")

type Contact struct {
	Id        uint      `json:"id"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (b *Contact) TableName() string {
	return ContactTableName
}
