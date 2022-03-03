package repository

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	mysql := NewDB().Mysql
	err := mysql.Ping()
	assert.Nil(t, err)
}
