package consts

// 网络类型常量
const (
	NetworkTypeETH = "ETH"  // 以太坊网络
	NetworkTypeTRX = "TRON" // 波场网络
)

// 代币类型
const (
	CoinTypeETH  = "ETH"  // 以太坊
	CoinTypeTRX  = "TRX"  // 波场
	CoinTypeUSDT = "USDT" // USDT
)

// 合约地址
const (
	ContractAddressERC20USDT = "0xdAC17F958D2ee523a2206206994597C13D831ec7" // 以太坊合约地址
	ContractAddressTRC20USDT = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"         // 波场合约地址
)

// 判断是否在TypeList中
func IsCoinType(coinType string) bool {
	for _, v := range []string{CoinTypeETH, CoinTypeTRX, CoinTypeUSDT} {
		if v == coinType {
			return true
		}
	}
	return false
}

// 判断是否在NetworkTypeList中
func IsNetworkType(networkType string) bool {
	for _, v := range []string{NetworkTypeETH, NetworkTypeTRX} {
		if v == networkType {
			return true
		}
	}
	return false
}
