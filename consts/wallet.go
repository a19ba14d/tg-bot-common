package consts

// 钱包结构体
type Wallet struct {
	Code           string `json:"code"`             // 类型编码 (存储值)
	CoinType       string `json:"coin_type"`        // 类型名称 (显示值)
	DecimalPlaces  int    `json:"decimal_places"`   // 小数位数
	IsCreateWallet bool   `json:"is_create_wallet"` // 是否创建钱包
	WalletType     string `json:"wallet_type"`      // 钱包类型
}

// 法币类型
const (
	FiatTypeUSD = "USD" // 美元
	FiatTypeCNY = "CNY" // 人民币
)

// 钱包代币类型
var AssetWalletList = []Wallet{
	{Code: CoinTypeETH, CoinType: "coin", DecimalPlaces: 8, IsCreateWallet: true, WalletType: "ASSET"},
	{Code: CoinTypeTRX, CoinType: "coin", DecimalPlaces: 8, IsCreateWallet: true, WalletType: "ASSET"},
	{Code: CoinTypeUSDT, CoinType: "coin", DecimalPlaces: 8, IsCreateWallet: true, WalletType: "ASSET"},

	// {Code: FiatTypeUSD, CoinType: "fiat", DecimalPlaces: 4, IsCreateWallet: true, WalletType: "ASSET"},
	// {Code: FiatTypeCNY, CoinType: "fiat", DecimalPlaces: 4, IsCreateWallet: true, WalletType: "ASSET"},
}
