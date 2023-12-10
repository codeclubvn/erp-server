package valid_pointer

import (
	uuid "github.com/satori/go.uuid"
)

func String(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

func Int(in *int) int {
	if in == nil {
		return 0
	}
	return *in
}

func UUID(in *uuid.UUID) uuid.UUID {
	if in == nil {
		return uuid.UUID{}
	}
	return *in
}

// pointer

func UUIDPointer(in uuid.UUID) *uuid.UUID {
	return &in
}
