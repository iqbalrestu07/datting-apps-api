package server

import (
	"fmt"
	"log"

	"github.com/iqbalrestu07/datting-apps-api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	conf := GetAPPConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBPort,
	)
	fmt.Println("dsn =-=-=-=-=-=-=-=-=-=", dsn)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	migrate(dbConn)
	return dbConn, err
}

func migrate(db *gorm.DB) {
	installUUIDOsspExtension(db)
	_ = db.AutoMigrate(
		domain.Interest{},
		domain.User{},
		domain.Match{},
		domain.Photo{},
		domain.UserInterest{},
	)
	seedInterests(db)

}

// seedInterests seeds the interests table with initial data.
func seedInterests(db *gorm.DB) {
	interests := []domain.Interest{
		{Name: "Sports"},
		{Name: "Music"},
		{Name: "Movies"},
		{Name: "Technology"},
		{Name: "Food"},
	}
	result := db.Create(&interests)
	if result.Error != nil {
		fmt.Printf("Error seeding interest : %v\n", result.Error)
	}

	fmt.Println("Interests seeded successfully.")
}

func installUUIDOsspExtension(db *gorm.DB) {
	// SQL command to create the uuid-ossp extension if not exists
	createExtensionSQL := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

	result := db.Exec(createExtensionSQL)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("uuid-ossp extension installed successfully.")
}
