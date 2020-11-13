import time
import os

#项目目录配置
BASE_DIR = os.path.dirname(os.path.abspath(__file__))

#要运行的测试用例的目录
cases_path = BASE_DIR + "/test_cases"

#测试报告目录
report_path = BASE_DIR + "/test_report/"

#失败重跑次数
rerun = 3

#当达到最大失败次数停止执行
max_fail = "5"

#请求地址
host = "http://123.56.170.101:8545"