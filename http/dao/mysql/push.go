package mysql

import (
	"github.com/jinzhu/gorm"
)

type DBNameDao struct {
	DB *gorm.DB
}

type UserDevice struct {
	ID 				uint 	`gorm:"column:id;primary_key;not null;unsigned data type;AUTO_INCREMENT"`
	UserID			uint	`gorm:"column:userId;not null;unsigned data type"`
	DeviceNum		string	`gorm:"column:deviceNum;type:varchar(64);not null;DEFAULT:''"`
	InsertTime		uint64	`gorm:"column:insertTime;not null;unsigned data type"`
	ModifyTime		uint64	`gorm:"column:modifyTime;not null;unsigned data type"`
	Platform		string	`gorm:"column:platform;type:char(1);not null;DEFAULT:''"`
	Source 			uint8  	`gorm:"column:source;type:tinyint(3);not null;unsigned data type;DEFAULT:1"`
	IsOff			bool	`gorm:"column:isOff;type:tinyint(1);not null;DEFAULT:0"`
	ExtId 			string	`gorm:"column:extId;type:varchar(128);not null;DEFAULT:''"`
	PushPlatform	string	`gorm:"column:pushPlatform;type varchar(20);not null"`
	Version        	int64 	`gorm:"column:version;not null"`
}

func(p *DBNameDao) SelectDevicenuByUseridSource(userID int, source int) (userDevice UserDevice, err error){
	p.DB.Table("userDevice").Where("userId = ?", userID).Where("source = ?", source).Find(&userDevice)
	return
}
