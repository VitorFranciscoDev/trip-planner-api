package usecases

import (
	"context"

	"github.com/VitorFranciscoDev/trip-planner-api/domain/entities"
	"github.com/VitorFranciscoDev/trip-planner-api/domain/repositories"
)

type TripUseCases struct {
	r repositories.TripRepository
}

func NewTripUseCases(r repositories.TripRepository) *TripUseCases {
	return &TripUseCases{r: r}
}

func (t TripUseCases) AddTrip(ctx context.Context, trip *entities.Trip) error {
	err := t.r.AddTrip(ctx, trip)
	if err != nil {
		return err
	}

	return nil
}

func (t TripUseCases) DeleteTrip(ctx context.Context, id int) error {
	err := t.r.DeleteTrip(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (t TripUseCases) UpdateTrip(ctx context.Context, trip *entities.Trip) error {
	err := t.r.UpdateTrip(ctx, trip)
	if err != nil {
		return err
	}

	return nil
}

func (t TripUseCases) GetAllTrips(ctx context.Context, userID int) ([]entities.Trip, error) {
	trips, err := t.r.GetAllTrips(ctx, userID)
	if err != nil {
		return nil, err
	}

	return trips, nil
}
