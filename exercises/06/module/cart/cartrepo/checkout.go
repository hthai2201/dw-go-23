package cartrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartmodel"
)

type CheckoutCartStorage interface {
	Checkout(ctx context.Context, total *cartmodel.Checkout) error
}

type checkoutCartRepo struct {
	store CheckoutCartStorage
}

func NewCheckoutCartRepo(store CheckoutCartStorage) *checkoutCartRepo {
	return &checkoutCartRepo{store: store}
}
func (repo *checkoutCartRepo) Checkout(ctx context.Context) (cartmodel.Checkout, error) {
	var c cartmodel.Checkout
	err := repo.store.Checkout(ctx, &c)
	if err != nil {
		return c, common.ErrCannotUpdateEntity(cartmodel.EntityName, err)
	}

	return c, nil
}
