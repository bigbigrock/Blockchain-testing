import logging
import os
import time
import pytest
from confest import cases_path,report_path,rerun,max_fail

logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s-%(levelname)s-%(message)s'
                    )
logger = logging.getLogger(__name__)


def run():
    logger.info("测试开始执行》》》》》")
    now_time = time.strftime("%Y-%m-%d-%H_%M_%S")
    # init_env(now_time)
    html_report = os.path.join(report_path,now_time,'report.html')
    pytest.main(['-v',cases_path,'--html='+html_report])
    logger.info("测试用例运行结束")

if __name__ == '__main__':
    run()