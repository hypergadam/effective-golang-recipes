package main

import (
	"context"
	"log/slog"
	"time"
)

func main() {
	// The trip service prepares a completed ride.
	ride := Ride{
		ID:       101,
		Time:     time.Now().UTC(),
		Duration: 45 * time.Minute,
		Distance: 8.5,
		Price:    12.75,
	}

	// Convert the ride into bytes for transport.
	data, err := MarshalRide(&ride)
	if err != nil {
		slog.ErrorContext(
			context.Background(),
			"failed to marshal ride",
			"err", err.Error(),
		)
	}

	slog.Info("Encoded ride", "ride", data)

	// simulate another service decoding the received bytes.
	var received Ride
	if err := UnmarshalRide(data, &received); err != nil {
		slog.ErrorContext(
			context.Background(),
			"failed to unmarshal ride",
			"err", err.Error(),
		)
	}

	slog.Info(
		"ride bytes received successfully",
		"id", received.ID,
		"distance", received.Distance,
		"price", received.Price,
	)
}
