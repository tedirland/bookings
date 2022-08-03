package repository

import (
	"time"

	"github.com/tedirland/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriciton) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomId int) (bool, error)
	SearchAvailabilityByDatesForAllRooms(start, end time.Time) ([]models.Room, error)
}
