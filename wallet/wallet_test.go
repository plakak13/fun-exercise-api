package wallet

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type MockWallet struct {
	wallets      []Wallet
	updateResult sql.Result
	err          error
}

// UpdateUserWallet implements Storer.
func (s MockWallet) UpdateUserWallet(wallet *Wallet) (sql.Result, error) {
	return s.updateResult, s.err
}

func (s MockWallet) Wallets(filter Filter) ([]Wallet, error) {
	return s.wallets, s.err
}

func (s MockWallet) Create(*Wallet) error {
	return s.err
}

func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := MockWallet{err: echo.ErrInternalServerError}
		w := New(stubError)
		w.WalletsHandler(c)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		createdAt := time.Date(2024, time.March, 25, 14, 19, 0, 729237000, time.UTC)
		stubList := MockWallet{
			wallets: []Wallet{
				{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Wallet", WalletType: "Credit Card", Balance: 100.00, CreatedAt: createdAt},
				{ID: 2, UserID: 1, UserName: "John Doe", WalletName: "John's Wallet", WalletType: "Credit Card", Balance: 200.00, CreatedAt: createdAt},
			},
			err: nil,
		}

		w := New(stubList)
		w.WalletsHandler(c)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
		}
		expectedJson := `[{"id":1,"user_id":1,"user_name":"John Doe","wallet_name":"John's Wallet","wallet_type":"Credit Card","balance":100,"created_at":"2024-03-25T14:19:00.729237Z"},{"id":2,"user_id":1,"user_name":"John Doe","wallet_name":"John's Wallet","wallet_type":"Credit Card","balance":200,"created_at":"2024-03-25T14:19:00.729237Z"}]`
		gotJson := rec.Body.String()
		fmt.Println(gotJson)
		fmt.Println(expectedJson)
		if !areJSONEqual(expectedJson, gotJson) {
			t.Errorf("Expected response body %s but got %s", expectedJson, gotJson)
		}

	})
}

func areJSONEqual(json1, json2 string) bool {
	var obj1, obj2 interface{}
	if err := json.Unmarshal([]byte(json1), &obj1); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(json2), &obj2); err != nil {
		return false
	}
	return reflect.DeepEqual(obj1, obj2)
}
