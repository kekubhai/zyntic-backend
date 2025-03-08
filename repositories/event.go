package repositories

import (
	"context"
	"time"

	"github.com/kekubhai/zyntic/models"
)

type EventRepository struct {
	db any
}


func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}
	events = append(events, &models.Event{
		ID:        "j48489893838334484",
		Name:      "Anirban",
		Location:  "Kolkata",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}
func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}
func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepositories(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}

}
