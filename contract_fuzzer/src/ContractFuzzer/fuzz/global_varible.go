package fuzz


const fuzz_Scale = 5
var (
	G_current_contract interface{}
	G_current_fun interface{}
	G_current_bin_fun_sigs = make(map[string][]string,0)
	G_current_abi_sigs = make(map[string]string,0)
)
