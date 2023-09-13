package utils

import (
	"context"
	"erp/api_errors"
	"errors"

	uuid "github.com/satori/go.uuid"
)

func GetUserStringIDFromContext(ctx context.Context) string {
	return ctx.Value("x-user-id").(string)
}

func GetUserUUIDFromContext(ctx context.Context) (uuid.UUID, error) {
	sid := ctx.Value("x-user-id").(string)

	u, err := uuid.FromString(sid)
	if err != nil {
		return uuid.Nil, errors.New(api_errors.ErrInvalidUserID)
	}

	return u, nil
}

func GetStoreIDFromContext(ctx context.Context) string {
	return ctx.Value("x-store-id").(string)
}

func GetPageCount(total int64, limit int64) int64 {
	if total == 0 {
		return 0
	}

	if total%limit != 0 {
		return total/limit + 1
	}

	return total / limit
}
