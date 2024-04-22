package postgres

import (
	"database/sql"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type UpdateResponse struct {
}

func (p *Postgres) UpdateUserWallet(wallet *wallet.Wallet) (sql.Result, error) {
	queryUpdate := "UPDATE user_wallet SET wallet_name = $3, wallet_type = $4, balance = $5 WHERE id = $1 AND user_id = $2"
	stmt, err := p.Db.Prepare(queryUpdate)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(wallet.ID, wallet.UserID, wallet.WalletName, wallet.WalletType)
	if err != nil {
		return nil, err
	}

	return res, nil
}
