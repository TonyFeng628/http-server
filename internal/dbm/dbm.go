package dbm

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "fmt"
  "log"
)

type DBManager struct {
	db  *gorm.DB
}

func NewDBManager(user string,pass string,addr string,dbname string) *DBManager {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,pass,addr,dbname)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  	if err != nil {
  		log.Fatal(err)
  	}
  	db.AutoMigrate(&ContactInfo{})
  	return &DBManager{
  		db:db,
  	}

}

func (dbm *DBManager)Write(input ContactInfo) error {
	dbm.db.Create(&input)
	return nil
}  

func (dbm *DBManager)Read(condition ContactInfo) []ContactInfo {
	contactInfo := make([]ContactInfo,0)
	dbm.db.Where(&condition).Find(&contactInfo)
	return contactInfo
}

func (dbm *DBManager)Update(input ContactInfo) error {
	log.Println(input)
	dbm.db.Debug().Where("id",input.ID).Updates(&input)
	return nil
}
	
func (dbm *DBManager)Delete(input ContactInfo) error {
	dbm.db.Delete(&input)
	return nil
}