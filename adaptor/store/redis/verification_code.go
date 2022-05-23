package redis

import (
	"context"
	"time"

	"gitlab.com/gocastsian/writino/entity"
)

func (r RedisStore) CreateVerCode(ctx context.Context, verCode entity.VerificationCode) error {

	err := r.client.Set(ctx, verCode.Email, verCode.Code, 5*time.Minute)
	return err.Err()
}

func (r RedisStore) FindVerCode(ctx context.Context, email string) (string, error) {

	res := r.client.Get(ctx, email)
	return res.Result()
}
