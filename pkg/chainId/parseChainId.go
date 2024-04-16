package chainId

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

func ParseChainId(numberString string) string {
	number, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		fmt.Println("Error converting number string to int64:", err)
		return numberString
	}
	return hexutil.EncodeBig(big.NewInt(number))
}
