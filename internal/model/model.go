package model

import (
	"fmt"
	"trading-system/pkg/setting"
	"trading-system/global"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	return db, nil
}