#-*-coding:utf8-*-
import os
path="./"
with open("README.md","w") as readme:
    content='''# mock 
    模拟各种数据及界面\n'''
    content+='''## 已完成列表\n|名称 | 代号 |\n| ---- | ---- |\n'''
    rank=1
    for item in sorted(os.listdir(path)):
        if item in ["README.md",__file__]:
            continue
        content+='''| {} | {} | \n'''.format(item,rank)
        rank+=1
    readme.write(content)    