// Package pgsql provides WeatherRequestRepository PgSQL implementation
package pgsql

import (
	"database/sql"
	"fmt"

	"github.com/Darkren/weatherservice/models"

	"github.com/Darkren/weatherservice/repository"
)

const tableName = "weather_requests"

// WeatherRequestRepository is a PgSQL WeatherReuestRepository
type WeatherRequestRepository struct {
	db *sql.DB
}

// New constructs and returns pointer to repository connected to the PgSQL DB
func New(db *sql.DB) repository.WeatherRequestRepository {
	return &WeatherRequestRepository{db: db}
}

// Add stores request in DB
func (r *WeatherRequestRepository) Add(req *models.WeatherRequest) (int64, error) {
	sql := fmt.Sprintf(`INSERT INTO %s
							(lat, lon, created, is_complete, is_in_progress) 
						VALUES
							($1, $2, NOW(), false, false)
						RETURNING id;`, tableName)

	err := r.db.QueryRow(sql, req.Lat, req.Lon).Scan(&req.ID)
	if err != nil {
		return 0, err
	}

	return req.ID, nil
}

// GetForProcessing returns first available not complete request and
// as a side effect marks it as IsInProgress
func (r *WeatherRequestRepository) GetForProcessing() (*models.WeatherRequest, error) {
	sql := fmt.Sprintf(`UPDATE %[1]s SET 
							is_in_progress = true 
						WHERE 
							id = ANY(
								SELECT id FROM %[1]s 
								WHERE is_complete = false AND is_in_progress = false 
								LIMIT 1 OFFSET 0
							)
						RETURNING *;`,
		tableName)

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	request := models.WeatherRequest{}

	if rows.Next() {
		err = rows.Scan(&request.ID, &request.Lat, &request.Lon, &request.Created,
			&request.IsComplete, &request.IsInProgress)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	return &request, nil
}

// ProcessingFinished marks the request as complete, also setting IsInProgress to false
// which is not neccessary but actually right
func (r *WeatherRequestRepository) ProcessingFinished(id int64) error {
	sql := fmt.Sprintf(`UPDATE %s SET 
							is_complete = true, 
							is_in_progress = false 
						WHERE 
							id = $1;`, tableName)

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
