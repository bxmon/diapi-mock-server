package storage

import (
	"testing"
	"time"

	"github.com/bxmon/diapi-mock-server/model"
	"github.com/google/go-cmp/cmp"
)

func TestStorage(t *testing.T) {
	s := NewStorage("account.db", "accountbucket")
	defer s.BoltDB.Close()

	user := model.User{
		ID:         1,
		FirstName:  "Test 01",
		LastName:   "model.User",
		FullName:   "model.User Test 01",
		Email:      "test.user@example.com",
		ProfilePic: "https://example.com/test/profile.png",
		CreateAt:   time.Now(),
		IsActive:   true,
	}

	t.Run("Add/Get user", func(t *testing.T) {
		err := s.AddNewUser(&user)
		if err != nil {
			t.Errorf("Error occurs while insert new user : %v", err)
		}

		insertedUser, err := s.GetUserByID(1)
		if err != nil {
			t.Errorf("Error occurs while query new inserted user : %v", err)
		}

		if !cmp.Equal(&user, insertedUser) {
			t.Errorf("Expect new inserted user %v equal to input user %v", insertedUser, user)
		}
	})

	t.Run("Replace/Get user", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			if err := s.ReplaceUser(&user); err != nil {
				t.Errorf("Error occurs while replace user : %v", err)
			}
		}

		insertedUser, err := s.GetUserByID(1)
		if err != nil {
			t.Errorf("Error occurs while query new inserted user : %v", err)
		}

		if !cmp.Equal(&user, insertedUser) {
			t.Errorf("Expect new inserted user %v equal to input user %v", insertedUser, user)
		}
	})

	t.Run("Update/Get user", func(t *testing.T) {
		user.FullName = "Thao Nguyen"
		if err := s.UpdateUser(&user); err != nil {
			t.Errorf("Error occurs while update user : %v", err)
		}

		insertedUser, err := s.GetUserByID(1)
		if err != nil {
			t.Errorf("Error occurs while query new inserted user : %v", err)
		}

		if insertedUser.FullName != "Thao Nguyen" {
			t.Errorf("Expect new fullname %v = Thao Nguyen", insertedUser.FullName)
		}
	})

	t.Run("Delete/Get user", func(t *testing.T) {
		err := s.DeleteUserByID(1)
		if err != nil {
			t.Errorf("Error occurs while delete user with id 1 : %v", err)
		}

		searchUser, err := s.GetUserByID(1)
		if err.Error() != "No user with id 1" {
			t.Errorf("Error occurs while query user with id 1 : %v", err)
		}

		if searchUser != nil {
			t.Errorf("Error occurs delete user hasn't done")
		}
	})

	t.Run("Add/Get multiple user", func(t *testing.T) {
		for i := 1; i < 5; i++ {
			user.ID = i
			if err := s.AddNewUser(&user); err != nil {
				t.Errorf("Error occurs while insert new user : %v", err)
			}
		}

		users, err := s.GetAllUsers()
		if err != nil {
			t.Errorf("Error occurs while get all user detail : %v", err)
		}

		if len(users) != 4 {
			t.Errorf("Expect to have 4 user but get %d", len(users))
		}

		for i, u := range users {
			if u.ID != i+1 {
				t.Errorf("Expect to have ID %d but get %d", i+1, u.ID)
			}
		}
	})
}
