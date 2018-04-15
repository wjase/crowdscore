package db

import (
	"fmt"

	"time"

	"github.com/jinzhu/gorm"
)

//BootstrapData sets up initial data
func BootstrapData(db *gorm.DB) {
	bootstrapTeams(db)
	bootstrapHeats(db)
}

func bootstrapTeams(db *gorm.DB) {
	teams := []Team{
		{Slug: "team1",
			Description: "The first team1",
		},
		{Slug: "team2",
			Description: "The first team2",
		},
		{Slug: "team3",
			Description: "The first team3",
		},
		{Slug: "team4",
			Description: "The first team4",
		},
		{Slug: "team5",
			Description: "The first team5",
		},
		{Slug: "team6",
			Description: "The first team6",
		},
		{Slug: "team7",
			Description: "The first team7",
		},
		{Slug: "team8",
			Description: "The first team8",
		},
		{Slug: "team9",
			Description: "The first team9",
		},
		{Slug: "team10",
			Description: "The first team10",
		},
	}
	var count int
	DB.Model(&Team{}).Count(&count)
	if count > 0 {
		fmt.Println("Skipping init of data - already present")
	}
	if count == 0 {
		for _, team := range teams {
			DB.Create(&team)
		}

	}

}

func bootstrapHeats(db *gorm.DB) {
	teams := []Heat{
		{Slug: "heat1",
			Description: "Daintree Level 3",
			StartTime:   time.Date(2018, 3, 20, 13, 0, 0, 0, time.Local).Unix(),
			EndTime:     time.Date(2018, 3, 20, 14, 0, 0, 0, time.Local).Unix(),
		},
		{Slug: "heat2",
			Description: "Uluru Level 4",
			StartTime:   time.Date(2018, 3, 20, 13, 0, 0, 0, time.Local).Unix(),
			EndTime:     time.Date(2018, 3, 20, 14, 0, 0, 0, time.Local).Unix(),
		},
		{Slug: "heat3",
			Description: "Boardroom Level 1",
			StartTime:   time.Date(2018, 3, 20, 13, 0, 0, 0, time.Local).Unix(),
			EndTime:     time.Date(2018, 3, 20, 14, 0, 0, 0, time.Local).Unix(),
		},
	}

	var count int
	DB.Model(&Heat{}).Count(&count)
	if count > 0 {
		fmt.Println("Skipping init of heat data - already present")
	}
	if count == 0 {
		for _, team := range teams {
			DB.Create(&team)
		}

	}

}
