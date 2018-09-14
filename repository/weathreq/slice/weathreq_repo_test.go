package slice

import (
	"testing"

	"github.com/Darkren/weatherservice/models"
)

func TestNew(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("Got err on creating repo")
	}
}

func TestAdd(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []*models.WeatherRequest{}

	id, err := repo.Add()
}

func TestGetNotComplete(t *testing.T) {

}
