package repositories

import (
	"context"

	models "github.com/kekubhai/zyntic/cmd/api"
)
type EventRepository struct {
	db any
}
func (r *EventRepository) GetMany(ctx context.Context)([]*models.Event,error){
	return nil, nil
}
func (r *EventRepository) GetOne(ctx context.Context,eventId string)(*models.Event,error){
	return nil, nil
}
func (r *EventRepository) CreateOne(ctx context.Context,event models.Event)(*models.Event,error){
	return nil, nil
}


func NewEventRepositories(db any) models.EventRepository{
	return &EventRepository{
		db:db,
	}

}