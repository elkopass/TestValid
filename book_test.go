//go:build integration
// +build integration

package main

import (
	"context"
	"testing"

	"github.com/Hendra-Huang/databaseintegrationtest/database"
	"github.com/Hendra-Huang/databaseintegrationtest/testingutil"
)

func TestInsertBook(t *testing.T) {
	t.Parallel()
	testDb, schemaName, cleanup := database.CreateTestDatabase(t)
	defer cleanup()

	loadTestData(t, testDb, schemaName, "book")
	insertBookQuery = addSchemaPrefix(schemaName, insertBookQuery)

	title := "New title"
	author := "New author"
	err := InsertBook(context.Background(), testDb, title, author)
	testingutil.Ok(t, err)
}

func TestGetBooks(t *testing.T) {
	t.Parallel()
	testDb, schemaName, cleanup := database.CreateTestDatabase(t)
	defer cleanup()

	loadTestData(t, testDb, schemaName, "book")
	getBooksQuery = addSchemaPrefix(schemaName, getBooksQuery)

	books, err := GetBooks(context.Background(), testDb)
	testingutil.Ok(t, err)
	testingutil.Equals(t, 2, len(books))
}
