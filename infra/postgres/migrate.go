package postgres

import (
	"log"

	"gorm.io/gorm"
)

var dbAutoMigrate = func(db *gorm.DB, dst ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if n, ok := r.(error); ok {
				err = n
			}
		}
	}()
	return db.AutoMigrate(dst...)
}

func Migrate(db *gorm.DB, models ...interface{}) error {
	log.Println("migrating database ... ", db.Migrator().CurrentDatabase())
	// auto migrate
	if err := dbAutoMigrate(
		db,
		models...,
	); err != nil {
		return err
	}

	log.Println("database migrated!")
	return nil
}
