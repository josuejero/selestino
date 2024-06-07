// internal/repository/mock_db.go

package repository

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
	*sql.DB
}

func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	arguments := m.Called(append([]interface{}{query}, args...)...)
	return arguments.Get(0).(sql.Result), arguments.Error(1)
}

func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	arguments := m.Called(append([]interface{}{query}, args...)...)
	return arguments.Get(0).(*sql.Row)
}

func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	arguments := m.Called(append([]interface{}{query}, args...)...)
	return arguments.Get(0).(*sql.Rows), arguments.Error(1)
}
