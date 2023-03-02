package test

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	HOST     string = "icd-cluster-test-mmfxtblz.postgres.database.azure.com"
	DATABASE string = "recordchangetracker"
	USER     string = "psqladminun@icd-cluster-test-mmfxtblz"
	PASSWORD string = "ucg06tmB3u6wqWTCxPW6sMcjULQuithC"
)

func TestPostgreSQL_test(t *testing.T) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", HOST, USER,
		PASSWORD, DATABASE)

	db, err := sql.Open("postgres", connectionString)
	assert.NoError(t, err)
	err = db.Ping()
	assert.NoError(t, err)

	sqlStatement := "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'"
	rows, err := db.Query(sqlStatement)
	assert.NoError(t, err)
	defer rows.Close()

	var tableName string
	var tables []string
	for rows.Next() {
		err := rows.Scan(&tableName)
		tables = append(tables, tableName)
		assert.NoError(t, err)
	}
	assert.Equal(t, 3, len(tables))

	expectedTables := []string{"partnervatid_previous", "partnervatid_current", "record_tracker_properties"}
	assert.Equal(t, expectedTables, tables)
}
