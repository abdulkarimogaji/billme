package db_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/abdulkarimogaji/billme/db"
	"github.com/abdulkarimogaji/billme/utils"
)

func TestGetUsers(t *testing.T) {
	// ctx := context.Background()
	err := db.TestStorage.Ping()
	if err != nil {
		t.Error("failed to ping test database")
	}

	// createRandomUsers(ctx, 5, t)

	// users, err := db.TestStorage.GetUsers(ctx, db.PaginationParams{}, db.GetUserFilters{})

	// if err != nil {
	// 	t.Errorf("Expected no error but got %s", err)
	// }

	// if len(users) < 5 {
	// 	t.Error("Expected at least 5 users but got 0")
	// }
}

func createRandomUsers(ctx context.Context, num int, t *testing.T) {
	var wg sync.WaitGroup
	var err error
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			_, err = db.TestStorage.CreateUser(ctx, db.CreateUserArgs{
				Username:    utils.GenerateRandomString(10),
				Email:       fmt.Sprintf("%s@%s.%s", utils.GenerateRandomString(6), utils.GenerateRandomString(5), utils.GenerateRandomString(3)),
				PhoneNumber: utils.GenerateRandomString(5),
				Role:        "test",
				Password:    utils.GenerateRandomString(10),
				Status:      1,
			})
			if err != nil {
				t.Errorf("Error creating random user %s", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()

}
