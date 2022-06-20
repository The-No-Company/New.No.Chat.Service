package Database

import "time"

const AdvertisementTableName = string("advertisements")

type Advertisement struct {
	Id        uint      `json:"id"`
	Post      string    `json:"post"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Advertisement) TableName() string {
	return AdvertisementTableName
}
