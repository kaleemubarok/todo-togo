package repository

import (
	"errors"
	"regexp"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestSRepo_SelectAllStatus(t2 *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t2.Fatalf("an error '%s' was not expected when opening a stub"+
			"database connection", err)
	}
	defer db.Close()
	rows := sqlxmock.NewRows([]string{"id", "desc"}).
		AddRow(1, "New").
		AddRow(2, "OnGoing").
		AddRow(3, "Done").
		AddRow(4, "Deleted")

	query := regexp.QuoteMeta("SELECT * FROM status")
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewStatusRepo(db)
	res, err := a.SelectAllStatus()
	assert.NoError(t2, err)
	assert.NotNil(t2, res)

	err = errors.New("An error occured")
	mock.ExpectQuery(query).WillReturnError(err)
	a = NewStatusRepo(db)
	res, err = a.SelectAllStatus()
	assert.Error(t2, err)
	assert.Nil(t2, res)

}
