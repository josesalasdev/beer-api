// clients/db.go

package clients

import (
	"fmt"

	"github.com/josesalasdev/beer-api/cmd/api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// ConnectDataBase return connection
func ConnectDataBase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		config.PostgresHostName,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDataBase,
		config.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(dsn)
	}

	//db.AutoMigrate(&Model{})     // nolint

	DB = db
}
