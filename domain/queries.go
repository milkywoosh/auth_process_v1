package domain

import "context"

type Queries interface {
	// StockTransferQueries(ctx context.Context) (ItemTransfer, error)
	AuthQueries(ctx context.Context) (Authentication, error)
}
