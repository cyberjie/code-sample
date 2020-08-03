#-*-coding:utf8-*-
import os
path="../lib/"
with open("../README.md","w") as readme:
    content='''# code-sample
    代码样本，编写和收集相关代码 (code sample,write and collect code.)。
    \n'''
    content+='''## 已完成列表\n|名称 | 代号 |\n| ---- | ---- |\n'''
    rank=1
    for item in sorted(os.listdir(path)):
        content+='''| {} | {} | \n'''.format(item,rank)
        rank+=1
    readme.write(content)    