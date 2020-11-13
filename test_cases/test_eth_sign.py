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


def test_hello():
    r=requests.request('post',"http://123.56.170.101:8545",
                       headers ={"Content-Type": "application/json" },
                       data = { "json": 2,"method": "eth_sign","params": ["0xe2043a5B808bf2215D94e0934182E224a38A47D7","0xdeadbeaf"]
        } )
    response = r.json()  # 返回字典格式
    return response
    # assert response['id'] == 1, "响应结果为：{}".format(response['result'])

a = test_hello()
print(a)



# if __name__ == '__main__':
#     pytest.main(['-v'])