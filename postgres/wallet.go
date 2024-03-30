package postgres

import (
	"log"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets() ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) WalletById(id int) (Wallet, error) {
	var wallet Wallet
	rows, err := p.Db.Query("SELECT * FROM wallet WHERE id = $1", id)
	if err != nil {
		return Wallet{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var w Wallet
		err = rows.Scan(&w)
		if err != nil {
			log.Fatal(err)
			return w, err
		}
		wallet = w
	}

	return wallet, nil
}
