package database

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/puvadon-artmit/gofiber-template/config"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/puvadon-artmit/gofiber-template/seeder/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.GetEnvConfig("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Wrong port!")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnvConfig("DB_HOST"),
		port, config.GetEnvConfig("DB_USER"),
		config.GetEnvConfig("DB_PASSWORD"),
		config.GetEnvConfig("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	if err = DB.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Permission{},
		&model.PermissionComponent{},
		&model.PermissionGroup{},
		&model.Category{},
		&model.Branch{},
		&model.Group{},
		&model.Group_Story{},
		&model.Status{},
		&model.Type_things{},
		&model.Manufacturer{},
		&model.Item_model{},
		&model.Story{},
		&model.Ground{},
		&model.Ground_Story{},
		&model.Responsible{},
		&model.Responsible_Story{},
		&model.Assets{},
		&model.Assets_Story{},
		&model.Typeplan{},
		&model.Typeplan_story{},
		&model.Asset_count{},
		&model.Inspection_story{},
		&model.Asset_check{},
		&model.Location{},
		&model.Location_story{},
		&model.Signature{},
		&model.Main_branch{},
		&model.Branch_Story{},
		&model.Main_Branch_Story{},
		&model.Round_Count{},
		&model.Asset_Round{},
		&model.Round_Count_Story{},
		&model.Asset_Round{},
		&model.Main_Category{},
		&model.Branch_Autoclik{},
		&model.Count_Autoclik{},
		&model.Asset_count_Main_Category{},
		&model.Asset_count_Category{},
		&model.Category_Story{},
	); err == nil && DB.Migrator().HasTable(&model.Role{}) {
		if err := DB.First(&model.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			for _, seed := range seeds.All() {
				if err := seed.Run(DB); err != nil {
					log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
				}
			}
		}
	}
	fmt.Println("Database Migrated")
}
