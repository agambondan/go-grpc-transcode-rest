//go:build !production
// +build !production

package repo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewRepositoriesTest(database ...string) (*Repositories, error) {
	dbPath := "file::memory:"
	if len(database) > 0 {
		dbPath = database[0]
		if dbPath == "" {
			dbPath = "file::memory:"
		}
	}
	gormConfig := gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gormConfig)
	if nil != err {
		panic(err)
	}
	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}
