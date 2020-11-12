import logging
import os
import time
import pytest
from confest import cases_path,report_path,rerun,max_fail

logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s-%(levelname)s-%(message)s',
                    )
logger = logging.getLogger(__name__)


def init_env(now_time):
    '''
    初始化测试报告目录
    '''
    os.mkdir(report_path + now_time)
    os.mkdir(report_path + now_time + "/image")

def run():
    logger.info("测试开始执行》》》》》")
    now_time = time.strftime("%Y-%m-%d-%H-%M-%S")
    init_env(now_time)
    html_report = os.path.join(report_path,now_time,'report.html')
    pytest.main(['-v',
                 cases_path,
                 "--html=" + html_report,
                 "--maxfail",max_fail,
                 "--reruns",rerun])
    logger.info("测试用例运行结束")

if __name__ == '__main__':
    run()