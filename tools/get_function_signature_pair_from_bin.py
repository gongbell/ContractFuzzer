import  os
import  sys
import  argparse
# fun_sig: <sig,sig_line_number>
# Code_lines = ["JUMPI"]
Code_lines = []
Jump_table = {}
# funSigs = [["0xadfaf1",10]]
def readFunSigs():
    funSigs_line_no = 0
    fun_sigs_line_no = []
    funSigs = []
    while funSigs_line_no<len(Code_lines):
        if Code_lines[funSigs_line_no]=="STOP":
            break
        if Code_lines[funSigs_line_no].startswith("PUSH4") and Code_lines[funSigs_line_no+1]=="EQ":
            funSigs.append([Code_lines[funSigs_line_no].split()[1],int(Code_lines[funSigs_line_no+2].split()[1],16)])
        funSigs_line_no += 1
    return  funSigs
    pass
def readFunBody(fun_sig_jump_line_no):
    fun_start_line_no = Jump_table[fun_sig_jump_line_no]
    fun_stop_line_no = fun_start_line_no+1
    while fun_stop_line_no<len(Code_lines):
        if Code_lines[fun_stop_line_no]=="STOP":
            break
        else:
            fun_stop_line_no = fun_stop_line_no + 1
    fun_body = Code_lines[fun_start_line_no:fun_stop_line_no+1]
    return  fun_body
def readSegs(fun_body):
    code_segs_line_no = []
    n = 0
    while n<len(fun_body):
        if fun_body[n].startswith("PUSH2") and (fun_body[n+1]=="JUMPI" or fun_body[n+1]=="JUMP"):
            code_segs_line_no.append(int(fun_body[n].split()[1],16))
        n = n + 1
    code_segs_bodys = []
    for seg_line_no in code_segs_line_no:
        code_segs_bodys.append(readCodeSeg(seg_line_no))
    return  code_segs_bodys
def readCodeSeg(code_jump_line_no):
    seg_start_line_no = Jump_table[code_jump_line_no]
    seg_JUMP_line_no = seg_start_line_no + 1
    while True:
        if Code_lines[seg_JUMP_line_no]=="JUMP":
            break
        else:
            seg_JUMP_line_no = seg_JUMP_line_no + 1
    return Code_lines[seg_start_line_no:seg_JUMP_line_no+1]
    pass
def read_innercall_sigs_from_codeseg(codeseg):
    valid = False
    Sigs = set()
    for line in codeseg:
        if line=="CALL":
             valid = True
        if line.startswith("PUSH4") and line.split()[1]!="0xffffffff":
            Sigs.add(line.split()[1])
    if valid:
        return Sigs
    else:
        return set()
def clearLines(lines):
    global  Code_lines
    nlines = []
    for line in lines:
        Code_lines.append(line.split(":")[1].strip())
    runtime_part_line_no = 0
    for line_no in range(len(Code_lines)):
        line = Code_lines[line_no]
        if line == "CODECOPY" and Code_lines[line_no + 1] == "PUSH1 0x00" and Code_lines[line_no + 2] == "RETURN":
            if Code_lines[line_no+3]=="STOP":
                runtime_part_line_no = line_no + 4
            else:
                runtime_part_line_no = line_no+3
            break
    if runtime_part_line_no > 0:
        delta = int(lines[runtime_part_line_no].split(":")[0])
        Code_lines = Code_lines[runtime_part_line_no:]
        lines = lines[runtime_part_line_no:]
    else:
        delta = 0
    print("delta:",delta)
    for line_no in range(len(Code_lines)):
        Jump_table[int(lines[line_no].split(":")[0].strip())-delta] = line_no
    return  Code_lines
def readFile(bin_file):
    with open(bin_file) as f:
        return clearLines(f.readlines()[1:])
    pass
def solve_file(bin_dir,bin_item):
    global  args
    if not os.path.exists("./sig"):
        os.mkdir("./sig")
    disam_data_lines = os.popen('evm disasm '+bin_dir+"/"+bin_item).readlines()
    Code_lines.clear()
    Jump_table.clear()
    try:
        lines = clearLines(disam_data_lines[1:])
    except IndexError:
        return
    fun_sigs = readFunSigs()
    D = dict()
    for item in fun_sigs:
        fun_sig = item[0]
        fun_jump_line_no = item[1]
        try:
            segs = readSegs(readFunBody(fun_jump_line_no))
        except IndexError:
            continue
        sigs = set()
        for seg in segs:
            sigs = sigs.union(read_innercall_sigs_from_codeseg(seg))
        if len(sigs) > 0:
            D[fun_sig] = sigs
        if len(D)!=0:
            with open("./sig/"+bin_item+".sig", "w+") as f:
                for fun_sig in D:
                    print(fun_sig)
                    print(D[fun_sig])
                    f.write(fun_sig + ": " + " ".join(D[fun_sig]))
                    f.write("\n")
                f.close()
    pass
def solve_dir(dir):
    items = os.listdir(dir)
    for item in items:
        try:
            solve_file(dir,item)
        except KeyError or IndexError:
            continue
def main():
    global  args
    parser = argparse.ArgumentParser()
    group = parser.add_argument_group('Model 1')
    groupex = group.add_mutually_exclusive_group(required=True)

    groupex.add_argument("-c", "--contract", type=str,dest="contract",
                       help="set the contract name whose function signature pair will be calculated")
    groupex.add_argument("-a", "--all", help="handle all contracts in directory specified by option '--bin_dir'",action="store_true")
    groupex2 = group.add_mutually_exclusive_group(required=True)
    groupex2.add_argument("-bd", "--bin_dir", type=str,dest="bin_dir",
                       help="set the contracts' bin directory where to get function signature pair")
    args = parser.parse_args()
    if args.contract:
       if args.contract.find("."):
           args.contract = args.contract.split(".")[0]+".bin"
    if args.bin_dir:
        if args.bin_dir[-1] == "/":
            args.bin_dir = args.bin_dir[:len(args.bin_dir)-1]
    if not args.all:
        solve_file(args.bin_dir,args.contract)
    else:
        solve_dir(args.bin_dir)
    pass

if __name__=="__main__":
    # test_file("./verified_contract_bins","TriWallet.bin")
    main()