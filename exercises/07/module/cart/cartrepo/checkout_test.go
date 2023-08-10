package cartrepo_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartrepo"
)

// MockCheckoutCartProductStorage is a mock implementation of AddCartProductStorage interface
type MockCheckoutCartProductStorage struct {
	Products []cartmodel.Cart
}

func (*MockCheckoutCartProductStorage) Checkout(ctx context.Context, total *cartmodel.Checkout) error {
	total.Total = 10
	return nil
}

func TestCartRepo_Checkout_Success(t *testing.T) {

	type args struct {
		wantTotal float64
	}
	mockStore := &MockCheckoutCartProductStorage{
		Products: []cartmodel.Cart{{
			ProductId: 1,
			Quantity:  2,
		},
			{
				ProductId: 2,
				Quantity:  3,
			},
		},
	}

	testcases := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Should checkout cart successfully",
			args: args{
				wantTotal: 10,
			},
			wantErr: nil,
		},
		{
			name: "Should update total value",
			args: args{
				wantTotal: 10,
			},
			wantErr: nil,
		},
	}

	for _, tc := range testcases {

		repo := cartrepo.NewCheckoutCartRepo(mockStore)
		total, err := repo.Checkout(context.Background())
		if total.Total != tc.args.wantTotal {
			t.Errorf("checkout_cart total = %.2f, wantTotal %.2f", total.Total, tc.args.wantTotal)

		}
		if !errors.Is(err, tc.wantErr) {
			t.Errorf("remove_product error = %.2f, wantErr %.2f", err, tc.args.wantTotal)
		}
	}

}
