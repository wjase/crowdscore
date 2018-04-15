package repos

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/au/com/cybernostics/crowdscore/server/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB *gorm.DB
)

func init() {
	user := services.User{}
	var err error
	DB, err = gorm.Open("sqlite3", filepath.Join(os.TempDir(), "gorm.db"))
	if err != nil {
		panic(err.Error())
	}
	DB.CreateTable(&user)
}
func TestCanSetPassword(t *testing.T) {

	user := services.User{
		Username: "barry@barry.com",
		Created:  time.Now().Unix(),
	}

	repo := services.NewUserRepository(DB)

	repo.CreateUser(&user)

	users, _ := repo.FindAll()

	for i, u := range users {
		fmt.Printf("%s - %d\n", u.Username, i)
	}

	user = services.User{
		Username: "john@barry.com",
		Created:  time.Now().Unix(),
	}

	repo.CreateUser(&user)

	users, _ = repo.Find(func(DB *gorm.DB) *gorm.DB { return DB.Where("username = ?", "john@barry.com") })

	for i, u := range users {
		fmt.Printf("%s - %d\n", u.Username, i)
	}

}
