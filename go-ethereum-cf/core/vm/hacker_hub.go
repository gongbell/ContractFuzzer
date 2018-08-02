/**
*  @hub.go   define data structure for recording infos
*  @Note: most parts of this file has been useless other than last function 'isRelOracle()'.
 */
package vm

import (
	"github.com/ethereum/go-ethereum/common"
)

var relOracle = common.HexToAddress("0xfa7b9770ca4cb04296cac84f37736d4041251cdf")

func isRelOracle(addr common.Address) bool {
	if addr.Big().Cmp(relOracle.Big()) == 0 {
		return true
	} else {
		return false
	}
}
