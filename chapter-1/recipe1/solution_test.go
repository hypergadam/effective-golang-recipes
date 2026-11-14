package recipe1

import (
	"errors"
	"io"
	"math"
	"testing"
	"time"
)

func TestRideRoundTrip(t *testing.T) {
	t.Parallel()
	want := Ride{
		ID:       1,
		Time:     time.Date(2026, time.November, 14, 12, 0, 0, 0, time.UTC),
		Duration: 15 * time.Minute,
		Distance: 8.5,
		Price:    12.75,
	}
	data, err := MarshalRide(&want)
	if err != nil {
		t.Fatal(err)
	}
	var got Ride
	if err := UnmarshalRide(data, &got); err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("round trip = %+v, want %+v", got, want)
	}
}

func TestUnmarshalRideInvalidInput(t *testing.T) {
	t.Parallel()

	for _, input := range []struct {
		name string
		data string
	}{
		{name: "empty", data: ""},
		{name: "malformed", data: "{"},
		{name: "wrong field type", data: `{"ID":"invalid"}`},
	} {
		t.Run(input.name, func(t *testing.T) {
			t.Parallel()
			var ride Ride
			if err := UnmarshalRide([]byte(input.data), &ride); err == nil {
				t.Fatal("expected a decoding error")
			} else if input.data == "" && !errors.Is(err, io.EOF) {
				t.Errorf("empty input error = %v, want EOF", err)
			}
		})
	}
}

func TestMarshalRideUnsupportedValue(t *testing.T) {
	t.Parallel()

	data, err := MarshalRide(&Ride{Price: math.NaN()})
	if err == nil {
		t.Fatal("expected an encoding error for NaN")
	}

	if data != nil {
		t.Errorf("data = %q, want nil on error", data)
	}
}
