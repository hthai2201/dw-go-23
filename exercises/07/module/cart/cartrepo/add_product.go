package cartrepo

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
)

type AddCartProductStorage interface {
	FindProduct(
		ctx context.Context,
		cond map[string]interface{},
		moreInfos ...string,
	) (*cartmodel.Cart, error)
	AddProduct(ctx context.Context, data *cartmodel.CartProductAdd) error
	UpdateProduct(ctx context.Context, data *cartmodel.CartProductUpdate) error
	RemoveProduct(ctx context.Context, productId int) error
}

type addCartProductRepo struct {
	store AddCartProductStorage
}

func NewAddCartProductRepo(store AddCartProductStorage) *addCartProductRepo {
	return &addCartProductRepo{store: store}
}
func (repo *addCartProductRepo) AddCartProduct(ctx context.Context, data *cartmodel.CartProductAdd) error {
	cart, err := repo.store.FindProduct(ctx, map[string]interface{}{"product_id": data.ProductId, "status": 1})
	if err != nil {
		err = repo.store.AddProduct(ctx, data)
		if err != nil {
			return common.ErrCannotUpdateEntity(cartmodel.EntityName, err)
		}
		return nil
	}

	if cart == nil {
		if cart.Product.Quantity < data.Quantity {
			return common.ErrCannotUpdateEntity("Quantity", err)
		}
		err = repo.store.AddProduct(ctx, data)
	} else if data.Quantity > 0 {
		newQuantity := data.Quantity + cart.Quantity
		if cart.Product.Quantity < newQuantity {
			return common.ErrCannotUpdateEntity("Quantity", err)
		} else if newQuantity < 1 {
			err = repo.store.RemoveProduct(ctx, data.ID)
		} else {
			updateData := &cartmodel.CartProductUpdate{

				ProductId: data.ProductId,
				Quantity:  newQuantity,
			}
			updateData.ID = cart.ID
			err = repo.store.UpdateProduct(ctx, updateData)
		}

	} else {
		err = repo.store.RemoveProduct(ctx, data.ID)
	}

	if err != nil {
		return common.ErrCannotUpdateEntity(cartmodel.EntityName, err)
	}

	return nil
}
