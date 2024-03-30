package wallet

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

type StubWallet struct {
	wallets []Wallet
	err     error
}

func (s StubWallet) Wallets() ([]Wallet, error) {
	return s.wallets, s.err
}
func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := StubWallet{err: echo.ErrInternalServerError}
		w := New(stubError)
		w.WalletHandler(c)
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

		stubList := StubWallet{
			wallets: []Wallet{},
			err:     nil,
		}

		w := New(stubList)
		w.WalletHandler(c)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}
	})
}
