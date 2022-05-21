package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"database/sql"

	_ "github.com/lib/pq" // postgresql driver
)

const (
	dbPort     = 5439
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "test"
)

func CreateTestDatabase(t *testing.T) (*sql.DB, string, func()) {
	connectionString := fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=disable", dbPort, dbUser, dbPassword, dbName)
	db, dbErr := sql.Open("postgres", connectionString)
	if dbErr != nil {
		t.Fatalf("Fail to create database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())
	tableName := "TestDataBase" + strconv.FormatInt(rand.Int63(), 10)

	_, err := db.Exec("CREATE SCHEMA " + tableName)
	if err != nil {
		t.Fatalf("Fail to create schema. %s", err.Error())
	}

	return db, tableName, func() {
		_, err := db.Exec("DROP SCHEMA " + tableName + " CASCADE")
		if err != nil {
			t.Fatalf("Fail to drop database. %s", err.Error())
		}
	}
}
