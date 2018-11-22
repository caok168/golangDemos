package app

import (
	"github.com/jinzhu/gorm"
	"time"
)

func ConnectPostgres(url string, extlog bool) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	if extlog {
		db.SetLogger(&gorm.Logger{})
	}
	db.SingularTable(false)
	db.Callback().Create().Replace("gorm:update_time_stamp", createCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, err
}

// createCallback will set `CreatedAt`, `ModifiedAt` when creating
func createCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}
