#!/usr/bin/python3
# -*- coding: utf-8 -*-
import  argparse
import subprocess

LOG = "log.txt"
BalanceGtZero_LOG = "list/BalanceGtZero.log"
SIGN_CALLEE="callee:"
SIGN_INPUT = "input:"
ADDR_NAME = dict()
Label = "HackerBalanceGtZero"
Delegate_fun_map = dict()
def importCsvMap(csvfile):
    lines = open(csvfile).readlines()
    for line in lines:
        str2 = line.strip().split(",")
        addr = str2[0]
        name = str2[1]
        ADDR_NAME[addr] = name.strip()
    return  ADDR_NAME
    pass
def solve(dir,BalanceGtZero_LOG_writer):
    global  ADDR_NAME
    lines = open(dir+"/"+LOG).readlines()
    S = set()
    for line in lines:
        # print(line)
        if line.find(Label)!=-1:
            callee = line[line.index(SIGN_CALLEE):].split(",")[0].split(SIGN_CALLEE)[1]
            print(callee[1:len(callee)-1])
            S.add(ADDR_NAME[callee[1:len(callee)-1].lower()])
            callee = callee[1:len(callee)-1].lower()
            fun  = line[line.index(SIGN_INPUT):].split(",")[0].split(SIGN_INPUT)[1][:8]
            if callee not in Delegate_fun_map:
                Delegate_fun_map[callee] = set()
                Delegate_fun_map[callee].add(fun)
            else:
                Delegate_fun_map[callee].add(fun)
    BalanceGtZero_LOG_writer.write("\n".join(S))
    

def parseLog2BalanceGtZeroList(BalanceGtZero_LOG):
    global args
    importCsvMap("./addrmap.csv")
   # BalanceGtZero_LOG = "list/BalanceGtZero.log"
    BalanceGtZero_LOG_writer = open(BalanceGtZero_LOG,"w+")
    parser = argparse.ArgumentParser()
    group = parser.add_argument_group('Model 1')
    groupex = group.add_mutually_exclusive_group(required=True)

    groupex.add_argument("-d", "--dir", type=str, dest="dir",
                         help="set directory where test_data in")
    args = parser.parse_args()
    if args.dir:
        if args.dir[-1] == "/":
            args.dir = args.dir[:len(args.dir) - 1]
        solve(args.dir,BalanceGtZero_LOG_writer)
    BalanceGtZero_LOG_writer.close()
    pass
def parseBin2No_ether_refundList(o_file):
   # o_file = "no_ether_refund.list"
    o_writer = open(o_file,"w+")
    bin_dir = "./verfied_contract_bins"
    bins = os.listdir(bin_dir)
    ls = list()
    for bin in bins:
        name = bin.split(".")[0]
        bin_path = "./verfied_contract_bins/%s" % bin
        cmd = "grep -E 'CALL|DELEGATECALL|SUICIDE|DESTRUCT' %s" % bin_path
        output = subprocess.getoutput(cmd)
        if len(output) == 0:
            ls.append(name)
    o_writer.write("TotalNo:"+str(len(ls)))
    o_writer("\n".join(ls))
    o_writer.close()  

def cmp(log_no_refund,log_balancegtzero,o_file):
     s1 = set()
     s2 = set()
     lines = open(log_no_refund).readlines()[1:]
     for line in lines:
         line = line.strip()
         s1.add(line)
     lines = open(log_balancegtzero).readlines()[1:]
     for line in lines:
         line = line.strip()
         s2.add(line)
     s3 = s1.intersection(s2)
     s1_2 = s1.difference(s2)
     s2_1 = s2.difference(s1)
     print(len(s3))
     print(len(s1_2))
     print(len(s2_1))
     with open(o_file,"w+") as f:
         f.write("TotalNo:"+str(len(s3))+"\n")
         f.write("\n".join(s3))
         f.close()
     pass
def main():
    o_file = "./list/freezingether.list"
    log_no_refund = "no_ether_refund.list"
    log_balancegtzero = "BalanceGtZero.log"
    parseBin2No_ether_refundList("no_ether_refund.list")
    parseLog2BalanceGtZeroList("BalanceGtZero.log")           
    cmp(fuzz_log,oyente_log,o_file)  
    pass

if __name__=="__main__":
    main()
