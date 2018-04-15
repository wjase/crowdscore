package db_test

import (
	"os"
	"path"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/wjase/crowdscore/db"
	"github.com/wjase/crowdscore/db/sqlite"
)

func setupTest(t *testing.T) func(*testing.T) {
	dbpath := path.Join(os.TempDir(), "testdb"+".db")
	dbToUse, err := sqlite.InitDB(dbpath)
	if err != nil || dbToUse == nil {
		t.Log(err)
		t.FailNow()
	}

	db.Init(dbToUse)

	return func(t *testing.T) {
		defer os.Remove(dbpath)
		dbToUse.Close()

	}

}

func TestCreateUsers(t *testing.T) {
	atEnd := setupTest(t)
	defer atEnd(t)

	results, err := db.DB.Raw("select * from users").Rows()
	assert.Nil(t, err)
	var count int
	db.DB.Model(&db.User{}).Count(&count)
	assert.Equal(t, 0, count)

	userToAdd := db.User{DisplayName: "bob",
		DepartmentID: 1,
		Username:     "bob@bob.com",
	}
	bstream, _ := json.Marshal(&userToAdd)
	t.Log(string(bstream))
	db.DB.Create(&userToAdd)
	assert.Nil(t, db.DB.Error)

	//list tables
	results, err = db.DB.Raw("select * from users").Rows()
	assert.Nil(t, nil)
	defer results.Close()

	db.DB.Model(&db.User{}).Count(&count)
	assert.Equal(t, 1, count)

}

// create user
// set password
// create team
// add self to team
// create heat
// add team to heat
// vote for team in heat
// get leaders in heat

//db.InitDB
