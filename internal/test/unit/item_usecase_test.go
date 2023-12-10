// item_handler_test.go
package unit

import (
	"context"
	"errors"
	"invoice-system/internal/adpaters/repository"
	"invoice-system/internal/core/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setUpDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/invoice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	db.AutoMigrate(
		&models.Item{},
		&models.Customer{},
	)

	return db
}

func TestCreateItem(t *testing.T) {

	testCases := []struct {
		name          string
		itemRequest   models.Item
		expectedError error
	}{
		{
			name: "ValidItem",
			itemRequest: models.Item{
				Name: "Design",
				Type: "Service",
			},
			expectedError: nil,
		},
		{
			name: "RequiredItemName",
			itemRequest: models.Item{
				Type: "Service",
				// Provide invalid values or incomplete data to test error handling
			},
			expectedError: errors.New(""), // Any error will match
		},
		// Add more test cases as needed
	}
	// Set up the database for testing
	testDB := setUpDB()
	defer func() {
		// Close the Gorm database connection
		sqlDB, err := testDB.DB()
		if err != nil {
			t.Fatalf("Error getting underlying database connection: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			t.Fatalf("Error closing database connection: %v", err)
		}
	}()

	// Create a new repository instance using the test database
	repo := repository.NewRepository(testDB)

	// Call your CreateItem function (assuming it's implemented in your repository)
	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createdItem, err := repo.CreateItemRepository(context.Background(), tc.itemRequest)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Nil(t, createdItem)
				// Add additional assertions for specific error cases
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, createdItem)
				assert.NotEmpty(t, createdItem.ID)
				assert.Equal(t, tc.itemRequest.Name, createdItem.Name)
				assert.Equal(t, tc.itemRequest.Type, createdItem.Type)
			}
		})
	}
}
