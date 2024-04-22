package wallet

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Create(wallet *Wallet) error
	Wallets(filter Filter) ([]Wallet, error)
	UpdateUserWallet(wallet *Wallet) (sql.Result, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//
//	@Param			wallet_type	query		string	false	"wallet type"	Enums(Savings, Credit Card, Crypto Wallet)
//
//	@Success		200			{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletsHandler(c echo.Context) error {
	wType := c.QueryParam("wallet_type")

	filter := Filter{
		WalletType: wType,
	}
	wallets, err := h.store.Wallets(filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// WalletsUser
//
//	@Summary		Get user wallets
//	@Description	Get all wallets of user
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		int	false	"user id"
//
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/{id}/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletsUserHandler(c echo.Context) error {
	uid := c.Param("id")

	if uid == "" {
		return c.JSON(http.StatusBadRequest, Err{Message: "user_id is required!"})
	}

	var userID int

	userID, err := strconv.Atoi(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	filter := Filter{
		UserID: userID,
	}

	wallet, err := h.store.Wallets(filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, wallet)
}

// CreateWallet
//
//	@Summary		Create wallet
//	@Description	Create wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		Wallet	true	"wallet request"
//
//	@Success		200		{object}	Wallet
//	@Router			/api/v1/wallets [post]
//	@Failure		500	{object}	Err
func (h *Handler) CreateUserWalletHandler(c echo.Context) error {

	wallet := new(Wallet)
	if err := c.Bind(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "failed to parse wallet data"})
	}

	if err := h.store.Create(wallet); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, wallet)
}

// UpdateWalletByID
//
//	@Summary		Create wallet
//	@Description	Create wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		Wallet	true	"wallet request"
//
//	@Success		200		{object}	Wallet
//	@Router			/api/v1/wallets [patch]
//	@Failure		500	{object}	Err
func (h *Handler) UpdateUserWalletHabdler(c echo.Context) error {
	wallet := new(Wallet)
	_, err := h.store.UpdateUserWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallet)
}
