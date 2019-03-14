package models

import (
	"fmt"
	"strings"

	u "go-hero/utils"

	"github.com/jinzhu/gorm"
)

// TableName : Override the default table name
func (Hero) TableName() string {
	return "hero"
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

	return u.Message(false, "Validation passed"), true
}

// GetHero : Fetches a hero from the database
func GetHero(id uint) *Hero {
	if id < 1 {
		return nil
	}

	hero := &Hero{}
	GetDB().Table("heroes").Where("id = ?", id).First(hero)
	if hero.Name == "" {
		// Hero not found
		return nil
	}

	return hero
}

// GetHeroes : Fetches all of the heroes in the database
func GetHeroes() []*Hero {
	heroes := make([]*Hero, 0)
	err := GetDB().Table("heroes").Order("id DESC").Find(&heroes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return heroes
}

// Create : Creates a new hero
func (hero *Hero) Create() map[string]interface{} {
	if resp, ok := hero.Validate(); !ok {
		return resp
	}

	// Check for duplicates
	temp := &Hero{}
	err := GetDB().Table("heroes").Where("name = ?", hero.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		return u.Message(false, "Insert operation failed")
	}
	// Hero name must be unique
	if temp.Name != "" {
		return u.Message(false, "This hero has already been created.")
	}

	GetDB().Create(hero)

	if hero.ID <= 0 {
		return u.Message(false, "Unable to create hero!")
	}

	response := u.Message(true, "Hero has been created")
	response["hero"] = hero
	return response
}

// Update : Updates an existing hero
func (hero *Hero) Update(id uint) map[string]interface{} {
	if id < 1 {
		return u.Message(false, "Invalid id!")
	}

	if resp, ok := hero.Validate(); !ok {
		return resp
	}

	// Ensure record exists
	temp := &Hero{}
	err := GetDB().Table("heroes").Where("id = ?", id).First(temp).Error
	if err != nil {
		fmt.Println(err)
		return u.Message(false, "Update operation failed, invalid id!")
	}

	hero.ID = temp.ID
	hero.CreatedAt = temp.CreatedAt
	err = GetDB().Save(hero).Error
	if err != nil {
		fmt.Println(err)
		return u.Message(false, "Update operation failed!")
	}

	response := u.Message(false, "Hero has been updated")
	response["hero"] = hero
	return response
}

// DeleteHero : Deletes an existing hero
func DeleteHero(id uint) *Hero {
	if id < 1 {
		return nil
	}

	hero := &Hero{}
	err := GetDB().Where("id = ?", id).Find(&hero).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = GetDB().Where("id = ?", id).Delete(Hero{}).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return hero
}
