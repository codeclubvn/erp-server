package utils

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

func ValidInt(in *int) int {
	if in == nil {
		return 0
	}
	return *in
}

func ValidFloat64(in *float64) float64 {
	if in == nil {
		return 0
	}
	return *in
}

func ValidString(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

func ValidTime(in *time.Time) time.Time {
	if in == nil {
		return time.Time{}
	}
	return *in
}

func ValidUUID(in *uuid.UUID) uuid.UUID {
	if in == nil {
		return uuid.UUID{}
	}
	return *in
}

//

func IntPointer(i int) *int {
	return &i
}

func UUIDPointer(i uuid.UUID) *uuid.UUID {
	return &i
}

func StringPointer(i string) *string {
	return &i
}
