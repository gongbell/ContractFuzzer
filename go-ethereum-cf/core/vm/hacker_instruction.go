package vm

import (
	"encoding/hex"
	"fmt"
	"os"
)

var times = 0

/**
  hacker_instruction.go, list the instruction opXXX which will be recorded and analyzed later
*/
type opFunc func(pc *uint64, evm *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error)

func Hacker_record(op OpCode, fun opFunc, pc *uint64, evm *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	f, err := os.OpenFile("/tmp/text.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	if times == 0 {
		if _, err = f.WriteString(hex.EncodeToString(contract.Code)); err != nil {
			panic(err)
		}
		times++
	}
	oper := contract.GetOp(*pc)
	defer f.Close()
	if _, err = f.WriteString(fmt.Sprintf("0x%X %d %s\n", *pc, uint64(oper), oper.String())); err != nil {
		panic(err)
	}

	if hacker_call_stack != nil {
		call := hacker_call_stack.peek()
		if call != nil {
			switch op {
			case DIV:
				call.OnDiv()
			case SDIV:
				call.OnDiv()
			case NOT:
				call.OnRelationOp(NOT)
			case LT:
				call.OnRelationOp(LT)
			case GT:
				call.OnRelationOp(GT)
			case SLT:
				call.OnRelationOp(SLT)
			case SGT:
				call.OnRelationOp(SGT)
			case EQ:
				call.OnRelationOp(EQ)
			case ISZERO:
				call.OnRelationOp(ISZERO)
			case AND:
				call.OnRelationOp(AND)
			case OR:
				call.OnRelationOp(OR)
			case XOR:
				call.OnRelationOp(XOR)
			case SHA3:
				call.OnSha3()
			case CALLER:
				call.OnCaller()
			case ORIGIN:
				call.OnOrigin()
			case CALLVALUE:
				call.OnCallValue()
			case CALLDATALOAD:
				call.OnCalldataLoad()
			case BLOCKHASH:
				call.OnBlockHash()
			case TIMESTAMP:
				call.OnTimestamp()
			case BALANCE:
				call.OnBalance()
			case NUMBER:
				call.OnNumber()
			case MLOAD:
				call.OnMload()
			case MSTORE:
				call.OnMstore()
			case SLOAD:
				call.OnSload()
			case SSTORE:
				call.OnSstore()
			case JUMP:
				call.OnJump()
			case JUMPI:
				call.OnJumpi()
			case GAS:
				call.OnGas()
			case CREATE:
				call.OnCreate()
			case CALL:
				break
			case CALLCODE:
				break
			case DELEGATECALL:
				break
			case SELFDESTRUCT:
				call.OnSuicide()
			case RETURN:
				call.OnReturn()
			default:
				break
			}
		}
	}
	return fun(pc, evm, contract, memory, stack)
}
