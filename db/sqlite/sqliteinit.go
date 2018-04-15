package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // reasons
)

// InitDB creates a new DB connection
func InitDB(dbname string, params ...string) (SQLDB *gorm.DB, err error) {
	SQLDB, err = gorm.Open("sqlite3", dbname)
	return
}
