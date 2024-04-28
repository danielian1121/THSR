package keyManager

import "context"

type Service interface {
	GetKeys(ctx context.Context, keys ...string) (map[string]string, error)
}
