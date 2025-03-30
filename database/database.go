package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("task.db"), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
			os.Exit(1)
		}
	})
	return db
}
