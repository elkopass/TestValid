package main

import (
	"context"
	"database/sql"
	"testing"
)

type Requests struct {
	Braces string
	Result string
}

var getRequestsQuery = `SELECT * FROM requests`
var insertRequestQuery = `INSERT INTO requests VALUES ($1, $2)`

func GetRequests(ctx context.Context, db *sql.DB) ([]Requests, error) {
	rows, err := db.QueryContext(ctx, getRequestsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	requestss := []Requests{}
	for rows.Next() {
		requests := Requests{}
		err := rows.Scan(&requests.Braces, &requests.Result)
		if err != nil {
			return requestss, err
		}
		requestss = append(requestss, requests)
	}

	return requestss, nil
}

func debugTest(ctx context.Context, db *sql.DB, t *testing.T) {
	_, err := GetRequests(ctx, db)
	if err != nil {
		t.Errorf(err.Error())
	}
}
