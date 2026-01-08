package repositories

import (
	"context"

	"github.com/VitorFranciscoDev/trip-planner-api/domain/entities"
)

type TripRepository interface {
	AddTrip(ctx context.Context, trip *entities.Trip) error
	DeleteTrip(ctx context.Context, id int) error
	UpdateTrip(ctx context.Context, trip *entities.Trip) error
	GetAllTrips(ctx context.Context, userID int) ([]entities.Trip, error)
}
