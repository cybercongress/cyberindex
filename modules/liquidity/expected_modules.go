package liquidity

type BankModule interface {
	RefreshBalances(height int64, addresses []string) error
}

type AuthModule interface {
	RefreshAccounts(height int64, addresses []string) error
}
