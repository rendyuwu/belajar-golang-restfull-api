package repository

import (
	"context"
	"database/sql"

	"github.com/rendyuwu/belajar-golang-restfull-api/model/domain"
)

type ProductCategory interface {
	Save(ctx context.Context, tx sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx sql.Tx, productId int) domain.Product
	Find(ctx context.Context, tx sql.Tx) []domain.Product
}
