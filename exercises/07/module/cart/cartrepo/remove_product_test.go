package cartrepo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartrepo"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
)

// MockRemoveCartProductStorage is a mock implementation of AddCartProductStorage interface
type MockRemoveCartProductStorage struct {
	FindProductFunc   func(ctx context.Context, cond map[string]interface{}, moreInfos ...string) (*cartmodel.Cart, error)
	RemoveProductFunc func(ctx context.Context, productId int) error
	Cart              *cartmodel.Cart
}

func (m *MockRemoveCartProductStorage) FindProduct(ctx context.Context, cond map[string]interface{}, moreInfos ...string) (*cartmodel.Cart, error) {
	return m.FindProductFunc(ctx, cond, moreInfos...)
}

func (m *MockRemoveCartProductStorage) RemoveProduct(ctx context.Context, productId int) error {
	return m.RemoveProductFunc(ctx, productId)
}

func TestCartRepo_RemoveCartProduct_Success(t *testing.T) {

	type args struct {
		cartProduct cartmodel.CartProductRemove
	}
	mockStore := &MockRemoveCartProductStorage{
		Cart: &cartmodel.Cart{
			ProductId: 1,
			Quantity:  5,
			Product: productmodel.Product{
				Quantity: 6,
			},
		},
	}

	testcases := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Should remove cart product",
			args: args{
				cartProduct: cartmodel.CartProductRemove{
					ProductId: 1,
					Quantity:  -5,
				},
			},
			wantErr: nil,
		},
		{
			name: "Should throw error when new quantity > product quantity",
			args: args{
				cartProduct: cartmodel.CartProductRemove{
					ProductId: 2,
					Quantity:  5,
				},
			},
			wantErr: common.ErrCannotGetEntity("Cart", nil),
		},
	}

	for _, tc := range testcases {

		mockStore.FindProductFunc = func(ctx context.Context, cond map[string]interface{}, moreInfos ...string) (*cartmodel.Cart, error) {
			if cond["product_id"] == mockStore.Cart.ProductId {
				return mockStore.Cart, nil
			}
			return nil, common.ErrCannotGetEntity("Cart", nil) // Return nil cart to simulate product not found in cart
		}

		mockStore.RemoveProductFunc = func(ctx context.Context, productId int) error {
			return nil
		}

		repo := cartrepo.NewRemoveCardProductRepo(mockStore)
		_, err := repo.RemoveCartProduct(context.Background(), &tc.args.cartProduct)

		if !errors.Is(err, tc.wantErr) {
			t.Errorf("remove_product error = %v, wantErr %v", err, tc.wantErr)
		}
	}

}
