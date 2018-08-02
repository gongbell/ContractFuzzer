/**
* @HackerUtils
* provide simple util function to compare,check eq,or hash for call info.
 */
package vm

import
(
	 "math/big"
     "strings"
	 "github.com/ethereum/go-ethereum/common"
)
type HackerUtils struct{}
type HackerStateDiff struct{
	addr common.Address
	isDiff bool
	storageDelta map[common.Hash][]common.Hash
	balanceDelta [3]big.Int
}
func(HackerUtils) Cmp(a,b common.Address) int{
	if strings.Compare(a.String(),b.String()) == 0{
		return 0
	}
	return -1
}
func (HackerUtils) equalStr(a,b string) bool{
	if strings.Compare(a,b) ==0{
		return true
	}else{
		return false
	}
}
func (HackerUtils) Hash(call *HackerContractCall) common.Hash{
	var (
		caller = call.caller
		callee = call.callee
		input = call.input
	)
	hash1 := caller.Hash()
	hash2 := callee.Hash()
	//hash3 := common.BytesToHash(input)
	var hash3 common.Hash
	if len(input)>20 {
		hash3 = common.BytesToHash(input[:20])
	} else{
		empty := make([]byte, 20-len(input))
		for i, _ := range empty{
			empty[i] = 0
		}
		hash3 = common.BytesToHash(append(input,empty...))
	}
	var h common.Hash
	h.SetBytes(new(big.Int).Add(new(big.Int).Add(hash1.Big(),hash2.Big()),hash3.Big()).Bytes())
	return h
}
func (HackerUtils) showDiffStorage(a map[common.Hash]common.Hash,b map[common.Hash]common.Hash)(bool,map[common.Hash][]common.Hash){
	var isDiff bool
	var delta map[common.Hash][]common.Hash
	var util HackerUtils
	isDiff= false
	delta = make(map[common.Hash][]common.Hash,0)
	for key,_ := range a{
		if util.equalStr(a[key].Hex(),b[key].Hex())!=true{
			isDiff = true
			pair := make([]common.Hash,0,2)
			pair = append(append(pair,a[key]),b[key])
			delta[key]=pair
		}
	}
	return isDiff,delta
}
func (HackerUtils) ShowDiffBalance(a big.Int,b big.Int) (bool, [3]big.Int){
	var isDiff bool
	var delta big.Int
	if a.Cmp(&b) != 0{
		delta = *new(big.Int).Sub(&a,&b)
		isDiff = true
	}else{
		isDiff = false
		delta = *new(big.Int).SetInt64(0)
	}
	bDelta := [...]big.Int{a,b,delta}
	return isDiff,bDelta
}

func (HackerUtils) ShowDiffState(a *HackerState,b *HackerState)(bool,*HackerStateDiff){
	var isDiff bool
	var diff  *HackerStateDiff
	var util HackerUtils
	isDiff = false
	for _,contractA := range a.contracts{
		for _,contractB := range  b.contracts{
			if util.Cmp(contractA.addr,contractB.addr) == 0{
				var balanceDelta [3]big.Int
				var storageDelta map[common.Hash][]common.Hash
				var tmp bool
				tmp,balanceDelta = util.ShowDiffBalance(contractA.balance,contractB.balance)
				if tmp== true{
					isDiff = true
				}
				tmp,storageDelta = util.showDiffStorage(contractA.storage,contractB.storage)
				if tmp == true{
					isDiff = true
				}
				if isDiff == true {
					diff = &HackerStateDiff{addr:contractA.addr,isDiff: isDiff, balanceDelta: balanceDelta, storageDelta: storageDelta}
					return isDiff,diff
				}
			}
		}
	}
	return 	isDiff,nil
}
func (HackerUtils) isSameCall(a ,b *HackerContractCall) bool{
	if a!=b&&a.callee.Big().Cmp(b.callee.Big())==0&&a.caller.Big().Cmp(b.caller.Big())==0&&strings.Compare(string(a.input),string(b.input))==0{
		return true
	}
	return false
}