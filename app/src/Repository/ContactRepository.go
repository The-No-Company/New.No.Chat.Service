package Repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"nochat/app/config"
	"nochat/app/src/Database"
)

func GetAllContacts(contact *[]Database.Contact) (err error) {
	if err = config.DB.Find(contact).Error; err != nil {
		return err
	}
	return nil
}

func CreateContact(contact *Database.Contact) (err error) {
	if err = config.DB.Create(contact).Error; err != nil {
		return err
	}
	return nil
}

func GetContactByID(contact *Database.Contact, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(contact).Error; err != nil {
		return err
	}
	return nil
}

func UpdateContact(contact *Database.Contact) (err error) {
	fmt.Println(contact)
	config.DB.Save(contact)
	return nil
}

func DeleteContact(contact *Database.Contact, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(contact)
	return nil
}

func GetLimitedAmount(contact *[]Database.Contact, count int) (err error) {
	if err = config.DB.Limit(count).Find(contact).Error; err != nil {
		return err
	}

	return nil
}

func GetCountRecords() (int, error) {
	var count int64
	var err error

	err = config.DB.Table(Database.ContactTableName).Count(&count).Error

	if err != nil {
		return -1, err
	}

	return int(count), nil
}
