package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

// Storage defines storage structure
type Storage struct {
	boltDB *bolt.DB
}

// NewStorage inits new storage
func NewStorage(name string) *Storage {
	db, err := bolt.Open(name, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("AccountBucket"))
		if err != nil {
			log.Fatalf("Create bucket failed: %s", err)
		}
		return nil
	})

	return &Storage{boltDB: db}
}

// AddNewUser add new user to database
func (s *Storage) AddNewUser(user *User) error {
	// Serialize the struct to JSON
	jsonBytes, _ := json.Marshal(user)

	// Generate a key
	key := strconv.Itoa(user.ID)

	// Write the data to the AccountBucket
	err := s.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))
		err := b.Put([]byte(key), jsonBytes)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID gets user by userid from database
func (s *Storage) GetUserByID(userID int) (*User, error) {
	// Allocate an empty user instance
	var user User

	// Generate a key
	key := strconv.Itoa(userID)

	// Read an object from the bucket using boltDB.View
	err := s.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("AccountBucket"))

		// Read the value identified by our userId supplied as []byte
		accountBytes := b.Get([]byte(key))
		if accountBytes == nil {
			return fmt.Errorf("No user with id %d", userID)
		}

		// Unmarshal the returned bytes into the user struct we created at
		// the top of the function
		json.Unmarshal(accountBytes, &user)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})

	// If there were an error, return the error
	if err != nil {
		return nil, err
	}
	// Return the user struct and nil as error.
	return &user, nil
}

// GetAllUsers gets all user from database
func (s *Storage) GetAllUsers() ([]*User, error) {

	var users []*User

	// Read an object from the bucket using boltDB.View
	err := s.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("AccountBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var user User
			json.Unmarshal(v, &user)
			users = append(users, &user)
		}

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})

	// If there were an error, return the error
	if err != nil {
		return nil, err
	}

	// Return the user struct and nil as error.
	return users, nil
}

// UpdateUser updates user in database
func (s *Storage) UpdateUser(user *User) error {
	return s.ReplaceUser(user)
}

// ReplaceUser replaces user in database
func (s *Storage) ReplaceUser(user *User) error {
	// Serialize the struct to JSON
	jsonBytes, _ := json.Marshal(user)

	// Generate a key
	key := strconv.Itoa(user.ID)

	// Write the data to the AccountBucket
	err := s.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))
		err := b.Put([]byte(key), jsonBytes)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserByID deletes user in database
func (s *Storage) DeleteUserByID(userID int) error {
	// Generate a key
	key := strconv.Itoa(userID)

	return s.boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("AccountBucket"))
		return bucket.Delete([]byte(key))
	})
}
