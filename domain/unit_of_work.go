package domain

import "context"

type UnitOfWork interface {
	BeginStockTransfer(ctx context.Context) (WarehouseSrv, error)
	BeginStockReceive(ctx context.Context) (WarehouseSrv, error)
}
