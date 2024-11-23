package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	GormDB                 *gorm.DB
	SqlDB                  *sql.DB
	DefaultPaginationLimit int
}

var (
	gormOpen    = gorm.Open
	singletonDB *DB
	once        sync.Once
)

func New() *DB {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			"Asia/Jakarta",
		)

		log.Println("initializing database ... ", os.Getenv("DB_NAME"))

		// set gorm debug level
		debugMode := logger.Info

		db, err := gormOpen(postgres.New(postgres.Config{
			DSN: dsn,
		}), &gorm.Config{
			SkipDefaultTransaction: true, // disable default transaction for write operations
			Logger:                 logger.Default.LogMode(debugMode),
			PrepareStmt:            true, // creates a prepared statement when executing any SQL and caches them to speed up future calls
		})
		if err != nil {
			panic(err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(10)

		err = sqlDB.Ping()
		if err != nil {
			panic(err)
		}

		log.Println("database initialized!")

		singletonDB = &DB{
			GormDB:                 db,
			SqlDB:                  sqlDB,
			DefaultPaginationLimit: 10,
		}
	})

	return singletonDB
}
