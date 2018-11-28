# -*- coding: utf-8 -*-
import codecs
import csv
import os
import types
import  xdrlib ,sys

import re
from datetime import datetime

import xlrd
import xlwt
from xlrd import xldate_as_tuple
def createWorkBook():
    return xlwt.Workbook()
    pass
def createSheet(workbook,sheet_name = "Sheet1"):
    return workbook.add_sheet(sheet_name)
    pass
def save2Excel(excel_file,Heads,rows):
    wb = createWorkBook()
    sheet = createSheet(wb)
    writeRows(sheet,Heads,rows)
    saveWorkBook(wb,excel_file)
    pass
def saveWorkBook(work_book,excel_file):
    work_book.save(excel_file)
    pass
def write(excel_sheet,row_index,row):
    for col in range(len(row)):
        excel_sheet.write(row_index,col,row[col])
    pass
def writeRows(excel_sheet,Heads,rows):
    row_index = 0
    write(excel_sheet,row_index,Heads)
    for row in rows:
        row_index = row_index + 1
        write(excel_sheet,row_index,forlist(Heads,row))
    pass
def forlist(Heads,row):
    nrow = list()
    print(row)
    print(Heads)
    for index in range(len(Heads)):
        Head = Heads[index]
        nrow.append(row[Head])
    return nrow
    pass
def open_excel(file= 'file.xls'):
    try:
        data = xlrd.open_workbook(file)
        return data
    except Exception as e:
        print(str(e))
#根据索引获取Excel表格中的数据   参数:file：Excel文件路径     colnameindex：表头列名所在行的所以  ，by_index：表的索引
def getRowValues(sheet,row_index,cols):
    row_content = []
    i = row_index
    for j in range(cols):
            ctype = sheet.cell(i, j).ctype  # 表格的数据类型
            cell = sheet.cell_value(i, j)
            if ctype == 2 and cell % 1 == 0:  # 如果是整形
                cell = int(cell)
            elif ctype == 3:
                # 转成datetime对象
                date = datetime(*xldate_as_tuple(cell, 0))
                cell = date.strftime('%Y.%m.%d')
            elif ctype == 4:
                cell = True if cell == 1 else False
            row_content.append(cell)
    return row_content
def excel_table_byindex(file= 'file.xls',colnameindex=0,by_index=0):
    data = open_excel(file)
    table = data.sheets()[by_index]
    nrows = table.nrows #行数
    ncols = table.ncols #列数
    colnames =  getRowValues(table,colnameindex,ncols) #某一行数据
    list =[]
    for rownum in range(1,nrows):
         row = getRowValues(table,rownum,ncols)
         # row = table.row_values(rownum)
         if row:
             app = {}
             for i in range(len(colnames)):
                 if type(row[i])==str:
                    app[colnames[i].strip()] = row[i].strip().replace("\r\n","").replace("\n"," ").replace("\r","").replace(" ","")
                 else:
                     app[colnames[i].strip()] = str(row[i])
             list.append(app)
    return list
def excel_table_header_byindex(file= 'file.xls',header_index=0,by_index=0):
    data = open_excel(file)
    table = data.sheets()[by_index]
    nrows = table.nrows #行数
    colnames =  table.row_values(header_index) #某一行数据
    colnames = replace(colnames, "\r\n", "")
    colnames = replace(colnames, "\n", "")
    colnames = replace(colnames, " ", "")
    return colnames
#根据名称获取Excel表格中的数据   参数:file：Excel文件路径     colnameindex：表头列名所在行的所以  ，by_name：Sheet1名称
def excel_table_byname(file= 'file.xls',colnameindex=0,by_name=u'Sheet1'):
    data = open_excel(file)
    table = data.sheet_by_name(by_name)
    nrows = table.nrows #行数
    colnames =  table.row_values(colnameindex) #某一行数据
    list =[]
    for rownum in range(colnameindex+1,nrows):
         row = table.row_values(rownum)
         if row:
             app = {}
             for i in range(len(colnames)):
                 if type(row[i]) == str:
                     app[colnames[i].strip()] = row[i].strip().replace("\r\n","").replace("\n"," ").replace("\r","").replace(" ","")
                 else:
                     app[colnames[i].strip()] = str(row[i])
             list.append(app)
    return list
def excel_table_header_byname(file= 'file.xls',header_index=0,by_name=u'Sheet1'):
    data = open_excel(file)
    table = data.sheet_by_name(by_name)
    nrows = table.nrows #行数
    colnames =  table.row_values(header_index) #某一行数据
    colnames = replace(colnames,"\r\n","")
    colnames = replace(colnames, "\n", "")
    colnames = replace(colnames, "\r", "")
    colnames = replace(colnames, "", "")
    return colnames
def replace(list,str_replace,str_target):
    for i in range(len(list)):
        list[i] = list[i].replace(str_replace, str_target)
     

    return list
