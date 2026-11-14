package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Ride struct {
	ID       int
	Time     time.Time
	Duration time.Duration
	Distance float64
	Price    float64
}

type Decoder struct {
	decoder *json.Decoder
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{decoder: json.NewDecoder(r)}
}

func (d *Decoder) DecodeRide(ride *Ride) error {
	if err := d.decoder.Decode(ride); err != nil {
		return fmt.Errorf("decode ride: %w", err)
	}

	return nil
}

type Encoder struct {
	encoder *json.Encoder
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{encoder: json.NewEncoder(w)}
}

func (e *Encoder) EncodeRide(ride *Ride) error {
	if err := e.encoder.Encode(ride); err != nil {
		return fmt.Errorf("encode ride: %w", err)
	}

	return nil
}

func UnmarshalRide(data []byte, ride *Ride) error {
	r := bytes.NewReader(data)
	return NewDecoder(r).DecodeRide(ride)
}

func MarshalRide(ride *Ride) ([]byte, error) {
	var buf bytes.Buffer
	if err := NewEncoder(&buf).EncodeRide(ride); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
