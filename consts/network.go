package consts

// 网络类型常量
const (
	NetworkTypeETH   = "TROM"  // 以太坊网络
	NetworkTypeTRC20 = "TRC20" // 波场网络
	NetworkTypeTRX   = "ETH"   // 波场网络
	NetworkTypeERC20 = "ERC20" // 以太坊网络

	CoinTypeETH  = "ETH"  // 以太坊
	CoinTypeTRX  = "TRX"  // 波场
	CoinTypeUSDT = "USDT" // USDT

	ContractAddressERC20USDT = "0xdAC17F958D2ee523a2206206994597C13D831ec7" // 以太坊合约地址
	ContractAddressTRC20USDT = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"         // 波场合约地址

	AddressTypeWhitelist = "whitelist" // 白名单地址（用户手动添加）
	AddressTypeDeposit   = "deposit"   // 系统充值地址（系统分配）
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
