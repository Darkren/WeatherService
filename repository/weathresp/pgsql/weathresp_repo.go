// Package pgsql provides WeatherResponseRepository PgSQL implementation
package pgsql

import (
	"database/sql"
	"fmt"

	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/repository"
)

const tableName = "weather_responses"

// WeatherResponseRepository is a PgSQL WeatherResponseRepository
type WeatherResponseRepository struct {
	db *sql.DB
}

// New constructs and returns pointer to repository connected to the PgSQL DB
func New(user, password, dbName, host string, port int) (repository.WeatherResponseRepository, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
		user, password, dbName, host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &WeatherResponseRepository{db: db}, nil
}

// Add stores response to repository
func (r *WeatherResponseRepository) Add(resp *models.WeatherResponse) (int64, error) {
	sql := fmt.Sprintf(`INSERT INTO %s
							(request_id, temperature, humidity, pressure, is_succeeded) 
						VALUES
							($1, 0, 0, 0, false)
						RETURNING id;`, tableName)

	err := r.db.QueryRow(sql, resp.RequestID).Scan(&resp.ID)
	if err != nil {
		return 0, err
	}

	return resp.ID, nil
}

// GetByRequestID returns response with the specified request ID
func (r *WeatherResponseRepository) GetByRequestID(requestID int64) (*models.WeatherResponse, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE request_id = $1;", tableName)

	rows, err := r.db.Query(sql, requestID)
	if err != nil {
		return nil, err
	}

	response := models.WeatherResponse{}
	if rows.Next() {
		err = rows.Scan(&response.ID, &response.RequestID, &response.Temperature,
			&response.Humidity, &response.Pressure, &response.IsSucceeded)
		if err != nil {
			return nil, err
		}
	}

	return &response, nil
}
