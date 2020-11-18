import requests
import pytest
import yaml
from confest import date_path,url,get_test_data
import json

cases,list_params = get_test_data(date_path + 'test_eth_getbalance')

#请求返回响应结果
class Test_eth(object):
    @pytest.mark.parametrize("case,http,expected",list(list_params),ids=cases)
    def test_sign(self,case,http,expected):
        r= requests.request(
            http["method"],
            url=url,
            headers=http["headers"],
            data=json.dumps(http['data'])
        )
        response = r.json() #返回字典格式
        print(response)
        assert response['result'] != None ,"响应结果为：{}".format(response)