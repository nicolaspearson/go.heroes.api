package models

import (
	"strings"

	u "go-hero/utils"

	"github.com/jinzhu/gorm"
)

// TableName : Override the default table name
func (Hero) TableName() string {
	return "heroes"
}

// Hero : A hero object struct
type Hero struct {
	gorm.Model
	Name     string
	Identity string
	Hometown string
	Age      uint
}

// Validate : Validates a hero's details
func (hero *Hero) Validate() (map[string]interface{}, bool) {
	if strings.TrimSpace(hero.Name) == "" {
		return u.Message(false, "A name is required"), false
	}

	if strings.TrimSpace(hero.Identity) == "" {
		return u.Message(false, "An identity is required"), false
	}

	if strings.TrimSpace(hero.Hometown) == "" {
		return u.Message(false, "A hometown is required"), false
	}

	if hero.Age < 1 {
		return u.Message(false, "An age is required"), false
	}

	// Hero name must be unique
	temp := &Hero{}

	// Check for duplicates
	err := GetDB().Table("heroes").Where("name = ?", hero.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Name != "" {
		return u.Message(false, "This hero has already been created."), false
	}

	return u.Message(false, "Validation passed"), true
}

// Create : Creates a new hero
func (hero *Hero) Create() map[string]interface{} {
	if resp, ok := hero.Validate(); !ok {
		return resp
	}

	GetDB().Create(hero)

	if hero.ID <= 0 {
		return u.Message(false, "Failed to create hero, connection error.")
	}

	response := u.Message(true, "Hero has been created")
	response["hero"] = hero
	return response
}

// GetHero : Fetches a hero from the database
func GetHero(id uint) *Hero {
	hero := &Hero{}
	GetDB().Table("heroes").Where("id = ?", id).First(hero)
	if hero.Name == "" {
		// Hero not found
		return nil
	}

	return hero
}
