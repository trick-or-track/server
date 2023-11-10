package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/trick-or-track/server/model"
)

type DataStore struct {
	db *sql.DB
}

func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{
		db: db,
	}
}

func (ds *DataStore) Create(data *model.Data) error {
	query, err := ds.db.Prepare(`
	INSERT INTO data (
		user_id, year,
		one, two, three,
		four, five, six,
		seven, eight, nine,
		ten, created_at
	) VALUES (
		$1, $2,
		$3, $4, $5,
		$6, $7, $8,
		$9, $10, $11,
		$12, $13
	) RETURNING id;`)
	if err != nil {
		return err
	}
	var insertedId int
	if err := query.QueryRow(
		data.UserID,
		data.Year,
		data.One,
		data.Two,
		data.Three,
		data.Four,
		data.Five,
		data.Six,
		data.Seven,
		data.Eight,
		data.Nine,
		data.Ten,
		time.Now().UTC(),
	).Scan(&insertedId); err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"data_unique_year_user_id\"" {
			return fmt.Errorf("Data all ready collected for year %d", data.Year)
		}
		return err
	}
	data.ID = insertedId
	return nil
}

func (ds *DataStore) GetByUserID(userID int) ([]*model.Data, error) {
	var result []*model.Data
	rows, err := ds.db.Query(`
	SELECT 
		id, user_id, year, 
		one, two, three, 
		four, five, six, 
		seven, eight, nine, ten
	FROM data 
	WHERE user_id = $1 ORDER BY year DESC;`, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var m model.Data
		if err := rows.Scan(
			&m.ID,
			&m.UserID,
			&m.Year,
			&m.One,
			&m.Two,
			&m.Three,
			&m.Four,
			&m.Five,
			&m.Six,
			&m.Seven,
			&m.Eight,
			&m.Nine,
			&m.Ten,
		); err != nil {
			return nil, err
		}
		result = append(result, &m)
	}
	return result, nil
}
