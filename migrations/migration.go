package migrations

import (
	"fmt"
	"invoice-system/internal/core/domain/models"
	"invoice-system/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Item{},
		&models.Customer{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
