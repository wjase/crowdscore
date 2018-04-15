package db

import (
	"github.com/jinzhu/gorm"
)

var (
	// DB is the gorm db obj
	DB *gorm.DB
)

// User represents a logged in user and a competitor
type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	// Password     string `json:"password" binding:"required"`
	DisplayName  string `json:"displayName"`
	DepartmentID int    `json:"department"`
}

type TeamMember struct {
	gorm.Model
	User   User
	UserID int
	Team   Team
	TeamID int
}

// UserPassword set the password for a user
type UserPassword struct {
	gorm.Model
	Password string
	UserID   int
}

type Team struct {
	gorm.Model
	Slug        string `gorm:"type:varchar(20);unique_index"`
	Description string
	URL         string
	TeamMembers []TeamMember
}

type Heat struct {
	gorm.Model
	Slug            string `gorm:"type:varchar(20);unique_index"`
	Description     string
	StartTime       int64
	EndTime         int64
	ShowLeaderboard bool
	Teams           []Team
}

type TeamsInHeat struct {
	gorm.Model
	TeamID  int
	RoundID int
}

type Department struct {
	gorm.Model
	Description string
}

type ScoringCategory struct {
	gorm.Model
	Order       int
	DisplayText string
	Description string
	URL         string
	Min         int
	Max         int
}

type SingleVoteRegister struct {
	gorm.Model
	UserID    int
	HeatID    int
	ProjectID int
}

type Score struct {
	gorm.Model
	UserID    int
	HeatID    int
	ProjectID int
	Rating    int
}

func Init(LoadedDB *gorm.DB) {

	DB = LoadedDB
	DB.CreateTable(&User{})
	DB.CreateTable(&TeamMember{})
	DB.CreateTable(&UserPassword{})
	DB.CreateTable(&Team{})
	DB.CreateTable(&Heat{})
	DB.CreateTable(&TeamsInHeat{})
	DB.CreateTable(&Department{})
	DB.CreateTable(&ScoringCategory{})
	DB.CreateTable(&SingleVoteRegister{})
	DB.CreateTable(&Score{})
}
