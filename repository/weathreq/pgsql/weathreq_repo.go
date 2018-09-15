package pgsql

import (
	"database/sql"
	"fmt"

	"github.com/Darkren/weatherservice/models"

	"github.com/Darkren/weatherservice/repository"
)

const tableName = "weather_requests"

type WeatherRequestRepository struct {
	db *sql.DB
}

func New(user, password, dbName, host string, port int) (repository.WeatherRequestRepository, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
		user, password, dbName, host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &WeatherRequestRepository{db: db}, nil
}

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

func (r *WeatherRequestRepository) GetForProcessing() (*models.WeatherRequest, error) {
	sql := fmt.Sprintf(`SELECT * FROM %s 
						WHERE is_complete = false AND is_in_progress = false 
						LIMIT 1 OFFSET 0;`,
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
	}

	return &request, nil
}

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
