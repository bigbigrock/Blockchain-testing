import time
import os


BASE_DIR = os.path.dirname(os.path.abspath(__file__))    #项目目录配置
cases_path = BASE_DIR + "/test_cases"   #要运行的测试用例的目录
report_path = BASE_DIR + "/test_report/"    #测试报告目录
date_path = BASE_DIR + "/test_data/"    #接口数据
rerun = 3   #失败重跑次数
max_fail = "5"  #当达到最大失败次数停止执行
url = "http://123.56.170.101:8545"  #请求地址