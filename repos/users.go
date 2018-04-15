package repos

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User basic defintion of a User
type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	// Password     string `json:"password" binding:"required"`
	DisplayName  string `json:"displayName"`
	DepartmentID int    `json:"department"`
	Created      int64  `json:"created"`
}

// UserRepository - db access for Users
type UserRepository struct {
	DB *gorm.DB
}

// UserPassword set the password for a user
type UserPassword struct {
	gorm.Model
	UserID   int
	Password string
}

// NewUserRepository takes a gorm DB and returns a user repo
func NewUserRepository(dbToUse *gorm.DB) UserRepository {
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// panic("failed to connect database")
	// }
	//defer db.Close()
	repository := UserRepository{DB: dbToUse}
	return repository
}

// GetPassword - gets the password for the user
func (r *UserRepository) GetPassword(ForUser *User) (string, error) {
	passwd := UserPassword{}
	r.DB.Model(&passwd).Where("user_id = ?", ForUser.ID).Find(&passwd)
	return "", nil
}

// DeleteUser removes the  pecified user
func (r *UserRepository) DeleteUser(UserToDelete *User) error {
	r.DB.Delete(&UserToDelete)
	return nil
}

// CreateUser creates a new DB record
func (r *UserRepository) CreateUser(UserToAdd *User) error {

	if r.DB.NewRecord(UserToAdd) {
		return r.DB.Create(&UserToAdd).Error
	}
	return errors.New("Already has a key set")
}

// UpdateUser updates the specified user
func (r *UserRepository) UpdateUser(UserToUpdate *User) error {
	err := r.DB.Update(UserToUpdate).Error
	return err
}

// Find use the specified criteria to locate users
func (r *UserRepository) Find(criteriaClause func(*gorm.DB) *gorm.DB) ([]User, error) {
	var users = []User{}
	r.withCriteria(criteriaClause).Find(&users)
	return users, r.DB.Error
}

// FindAll use the specified criteria to locate users
func (r *UserRepository) FindAll() ([]User, error) {
	var users = []User{}
	r.DB.Find(&users)
	return users, r.DB.Error
}

func (r *UserRepository) withCriteria(criteriaClause func(*gorm.DB) *gorm.DB) *gorm.DB {
	if criteriaClause == nil {
		criteriaClause = func(db *gorm.DB) *gorm.DB { return db }
	}
	return criteriaClause(r.DB)
}
