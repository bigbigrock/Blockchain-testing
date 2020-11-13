import requests
import pytest
import yaml
from confest import date_path,url


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

cases,list_params = get_test_data(date_path + 'test_eth_sign')

# class Test_Sign(object):
#     @pytest.fixture()
#     @pytest.mark.parametrize("case,http,expected",list(list_params),ids=cases)
#     def test_sign(self,url,case,http,expected):
#         r= requests.request(
#             http["method"],
#             url=url,
#             headers=http["headers"],
#             data=http['params']
#         )
#         response = r.json() #返回字典格式
#         assert response['id'] == expected['id'],"响应结果为：{}".format(response['result'])





# if __name__ == '__main__':
#     pytest.main(['-v'])