import time
import os
import yaml
import pytest
import requests
import json

BASE_DIR = os.path.dirname(os.path.abspath(__file__))    #项目目录配置
cases_path = BASE_DIR + "/test_cases"   #要运行的测试用例的目录
report_path = BASE_DIR + "/test_report/"    #测试报告目录
date_path = BASE_DIR + "/test_data/"    #接口数据
rerun = 3   #失败重跑次数
max_fail = "5"  #当达到最大失败次数停止执行
url = "http://123.56.170.101:8545"  #请求地址


#获取测试数据
def get_test_data(test_data_path):
    case = [] #存储测试用例名称
    http = [] #存储请求对象
    expected = [] #存储预期结果
    with open(test_data_path,encoding='utf-8') as f:
        date = yaml.load(f.read(),Loader=yaml.SafeLoader)
        test = date['tests']
        for td in test:
            case.append(td.get('case'))
            http.append(td.get('http'))
            expected.append('expected')
    parameters = zip(case,http,expected)
    return case,parameters


