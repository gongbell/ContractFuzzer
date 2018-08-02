import os
import re

import  requests
from pyquery import PyQuery as pq
import  excelUtil
url = r'https://etherscan.io/contractsVerified'
code_url = r'https://etherscan.io/address/'
def getContracts(table):
    pass
def getContractSourceCode(dom,contractName):
    source_code = dom("#dividcode")(".js-sourcecopyarea").text()
    # print(type(source_code))
    source_code.replace("&gt;",">").replace("&lt;","<")
    file_path = r'verified_contracts/' + contractName + ".sol"
    # if os.path.exists(file_path):
    #     return None
    print(file_path)
    # if os.path.exists(file_path):
    out = open(file_path, "w+")
    out.write(source_code)
    out.close()
    return source_code
    pass
def getContractAbi(dom,contractName):
    abi = dom("#dividcode")("#js-copytextarea2")
    if abi.text() is None or len(abi.text())==0:
        return None
    file_path = r'verified_contract_abis/' + contractName + ".abi"
    # if os.path.exists(file_path):
    #     return None
    print(file_path)
    out = open(file_path,"w+")
    out.write(abi.text())
    out.close()
    # print(abi.text())
    return abi.text()
    pass
def getContractBin(dom,contractName):
    try:
        bin = dom("#dividcode")("#verifiedbytecode2")
        if bin.text() is None or len(bin.text())==0:
            return None
        file_path = r'verified_contract_bins/' + contractName + ".bin"
        # if os.path.exists(file_path):
        #     return None
        print(file_path)
        out = open(file_path, "w+")
        out.write(bin.text())
        out.close()
        # print(bin.text())
        return bin.text()
    except IOError as e:
        print(e)
        return None
    pass
def getContractConstructorParams(dom,contractName):
    try:
        wordwraps  = dom("pre.wordwrap")
        if len(wordwraps)<4:
            return None
        constructor = wordwraps.eq(2)
        # print(dom("pre.wordwrap").eq(2).text())
        # print(constructor.text())
        if constructor.text() is None or len(constructor.text())==0:
             return None
        file_path = r'verified_contract_constructorparams/' + contractName + ".constructorparams"
        print(file_path)
        out = open(file_path, "w+")
        txt = [constructor.text().split("-----Decoded View---------------")[0]]
        txt.extend(constructor.text().split("-----Decoded View---------------")[1].split("Arg"))
        txt = "\n".join(txt)
        print(txt)
        out.write(txt)
        out.close()
        # print(bin.text())
        return constructor.text()
    except IndexError as e:
        print(e)
        return None
    pass
def getCode(contract):
    address = contract["Address"]
    name = contract["ContractName"]
    # address = "0xe9203c8c8eeaeffefecffad6d883c3ea205a832b"
    # name = "MultiSigWallet"
    surl = code_url + address + "#code"
    print(surl)
    page = requests.get(surl)
    # print(page.text)
    dom = pq(page.text)
    getContractSourceCode(dom, name)
    getContractAbi(dom,name)
    getContractBin(dom,name)
    getContractConstructorParams(dom,name)

    pass
def getPage(index):
    hurl = url
    if index != 1:
        hurl = url +r'/'+str(index)
    page = requests.get(hurl)
    dom  = pq(page.text)
    tableNode = dom(".container")(".row").eq(2)("table")
    return tableNode
    pass
def getTable(tableNode):
    thead = tableNode("thead")
    print(thead)
    heads = thead("th")
    ls = list()
    for i in range(len(heads)):
        ls.append(heads.eq(i).text())
    print(ls)
    tbody = tableNode("tbody")
    trs = tbody("tr")
    table = list()
    table.append(ls)
    for i in range(len(trs)):
        tr = trs.eq(i)
        tds = tr("td")
        item = list()
        for j in range(len(tds)):
            # if j == 0:
            #     address = tds.eq(j)("a").text()
            #     item.append(address)
            #     #print(tds.eq(j)("a").text())
            ele = tds.eq(j).text()
            item.append(ele)
        print(item)
        table.append(item)
    return table
    pass
def getPageInfo(url):
    hurl = url
    print(hurl)
    page = requests.get(hurl)
    dom  = pq(page.text)
    profile = dom(".container")(".row").eq(1)
    regex = re.compile(r'A Total Of (\d+) verified contract source codes found')
    totalInfo = profile.children(".col-md-6").eq(0).children("span").eq(1)
    m = regex.match(totalInfo.text())
    totalCount = 0
    if m:
        totalCount = m.group(1)
        print(m.group(1))
    lastHref = profile.children(".col-md-6").eq(1)("#ContentPlaceHolder1_HyperLinkLast").eq(0).attr("href")
    print(lastHref)
    lastregex = re.compile(r'contractsVerified/(\d+)')
    m = lastregex.match(lastHref)
    pageSize = 0
    if m:
        pageSize = m.group(1)
        print(pageSize)
    d = {"totalCount":totalCount,"pageSize":pageSize}
    return d
    pass


def start():
    workbook = excelUtil.createWorkBook()
    sheet = excelUtil.createSheet(workbook,"Verified Contracts")
    excel_file = "智能合约.xls"
    sheet_row_index = 0
    d = getPageInfo(url=url)
    print(d)
    for i in range(int(d["pageSize"])):
        tableNode = getPage(i)
        table = getTable(tableNode)
        if sheet_row_index == 0:
            for j in range(len(table)):
                excelUtil.write(sheet,sheet_row_index,table[j])
                sheet_row_index = sheet_row_index +1
        else:
            for j in range(1,len(table)):
                excelUtil.write(sheet, sheet_row_index, table[j])
                sheet_row_index = sheet_row_index + 1
    excelUtil.saveWorkBook(workbook,excel_file)
    getContracts()
def start2():
    print(url)
    d = getPageInfo(url=url)
    print(d)
    for i in range(int(d["pageSize"])):
        print("page",i)
        tableNode = getPage(i)
        table = getTable(tableNode)
        contract = {}
        heads = table[0]
        for j in range(1,len(table)):
            for k in range(0,len(table[0])):
                contract[heads[k]] = table[j][k]
            getCode(contract)
            break
        break
    pass
def clearEmptyFile(dir):
    items = os.listdir(dir)
    for item in items:
        print(item)
        file_path = dir + r'/' + item
        with open(file_path, "r") as f:
            if len(f.read()) == 0:
                print(file_path)
                f.close()
                os.remove(file_path)
    pass

# def start3():
def clearEmpty():
    dirs  = [r'./verified_contracts',r'./verified_contract_abis',r'./verified_contract_bins',r'./verified_contract_constructorparams']
    for dir in dirs:
        clearEmptyFile(dir)
    pass
def getContracts():
    contracts = excelUtil.excel_table_byname("智能合约.xls",by_name='Verified Contracts')

    # for contract in contracts:
    #     print(contract)
    #     getCode(contract)
    f = open("log.txt","a+")
    start = 3561
    # start 2 =1010
    for i in range(start,len(contracts)):
        getCode(contracts[i])
        print(i,contracts[i]["ContractName"])
        str = "%s %s\n" %(i,contracts[i]["ContractName"])
        f.write(str)
        f.flush()
    f.close()
    pass
def start4():
    getContracts()
    pass
def getFiles(path):
    dirs = os.listdir(path)
    return dirs
    pass
def start5():
    dir = r'./verified_contracts'
    files = getFiles(dir)
    print(len(files))
    pass
def main():
    # start()
    # start2()
    # start3()
    start4()
    # start5()
    pass
if __name__ == "__main__":
    main()

