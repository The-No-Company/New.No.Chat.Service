package Repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"nochat/app/config"
	"nochat/app/src/Database"
)

func GetAllAdvertisements(advertisement *[]Database.Advertisement) (err error) {
	if err = config.DB.Find(advertisement).Error; err != nil {
		return err
	}
	return nil
}

func CreateAdvertisement(advertisement *Database.Advertisement) (err error) {
	if err = config.DB.Create(advertisement).Error; err != nil {
		return err
	}
	return nil
}

func GetAdvertisementByID(advertisement *Database.Advertisement, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(advertisement).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAdvertisement(advertisement *Database.Advertisement) (err error) {
	fmt.Println(advertisement)
	config.DB.Save(advertisement)
	return nil
}

func DeleteAdvertisement(advertisement *Database.Advertisement, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(advertisement)
	return nil
}
