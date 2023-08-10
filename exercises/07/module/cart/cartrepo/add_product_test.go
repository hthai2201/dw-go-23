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

// MockAddCartProductStorage is a mock implementation of AddCartProductStorage interface
type MockAddCartProductStorage struct {
	FindProductFunc   func(ctx context.Context, cond map[string]interface{}, moreInfos ...string) (*cartmodel.Cart, error)
	AddProductFunc    func(ctx context.Context, data *cartmodel.CartProductAdd) error
	UpdateProductFunc func(ctx context.Context, data *cartmodel.CartProductUpdate) error
	RemoveProductFunc func(ctx context.Context, productId int) error
	Cart              *cartmodel.Cart
}

func (m *MockAddCartProductStorage) FindProduct(ctx context.Context, cond map[string]interface{}, moreInfos ...string) (*cartmodel.Cart, error) {
	return m.FindProductFunc(ctx, cond, moreInfos...)
}

func (m *MockAddCartProductStorage) AddProduct(ctx context.Context, data *cartmodel.CartProductAdd) error {
	return m.AddProductFunc(ctx, data)
}

func (m *MockAddCartProductStorage) UpdateProduct(ctx context.Context, data *cartmodel.CartProductUpdate) error {
	return m.UpdateProductFunc(ctx, data)
}

func (m *MockAddCartProductStorage) RemoveProduct(ctx context.Context, productId int) error {
	return m.RemoveProductFunc(ctx, productId)
}

func TestAddCartProductRepo_AddCartProduct_Success(t *testing.T) {

	type args struct {
		cartProduct cartmodel.CartProductAdd
	}
	mockStore := &MockAddCartProductStorage{
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
	}{{
		name: "Should create new cart product",
		args: args{
			cartProduct: cartmodel.CartProductAdd{
				ProductId: 2,
				Quantity:  1,
			},
		},
		wantErr: nil,
	}, {
		name: "Should update cart product quantity",
		args: args{
			cartProduct: cartmodel.CartProductAdd{
				ProductId: 1,
				Quantity:  1,
			},
		},
		wantErr: nil,
	},
		{
			name: "Should remove cart product",
			args: args{
				cartProduct: cartmodel.CartProductAdd{
					ProductId: 1,
					Quantity:  -5,
				},
			},
			wantErr: nil,
		},
		{
			name: "Should throw error when new quantity > product quantity",
			args: args{
				cartProduct: cartmodel.CartProductAdd{
					ProductId: 1,
					Quantity:  5,
				},
			},
			wantErr: common.ErrCannotUpdateEntity("Quantity", nil),
		},
	}

	for _, tc := range testcases {

		mockStore.FindProductFunc = func(ctx context.Context, cond map[string]interface{}, moreInfos ...string) (*cartmodel.Cart, error) {
			if cond["product_id"] == mockStore.Cart.ProductId {
				return mockStore.Cart, nil
			}
			return nil, common.ErrCannotGetEntity("test", errors.New("test")) // Return nil cart to simulate product not found in cart
		}

		mockStore.AddProductFunc = func(ctx context.Context, data *cartmodel.CartProductAdd) error {
			return nil
		}
		mockStore.UpdateProductFunc = func(ctx context.Context, data *cartmodel.CartProductUpdate) error {
			return nil
		}
		mockStore.RemoveProductFunc = func(ctx context.Context, productId int) error {
			return nil
		}

		repo := cartrepo.NewAddCartProductRepo(mockStore)
		err := repo.AddCartProduct(context.Background(), &tc.args.cartProduct)

		if !errors.Is(err, tc.wantErr) {
			t.Errorf("add_product error = %v, wantErr %v", err, tc.wantErr)
		}
	}

}
