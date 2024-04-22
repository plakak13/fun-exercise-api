package postgres

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

func (p *Postgres) Create(wallet *wallet.Wallet) error {
	row := p.Db.QueryRow("INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance) values($1, $2, $3, $4, $5) RETURNING id, created_at",
		wallet.UserID,
		wallet.UserName,
		wallet.WalletName,
		wallet.WalletType,
		wallet.Balance)
	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(&wallet.ID, &wallet.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
