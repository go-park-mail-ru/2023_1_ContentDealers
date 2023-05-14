package payment

import "context"

type Gateway interface {
	Subscribe(ctx context.Context, userID uint64) error
}
