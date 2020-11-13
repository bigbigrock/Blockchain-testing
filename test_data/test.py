import requests
def test_hello():
    r=requests.request('post',"http://123.56.170.101:8545",
                       headers ={"Content-Type": "application/json" },
                       data = { "json": 2,"method": "eth_sign","params": ["0xe2043a5B808bf2215D94e0934182E224a38A47D7","0xdeadbeaf"]
        } )
    response = r.json()  # 返回字典格式
    return response



a = test_hello()
print(a)
