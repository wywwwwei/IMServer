# A Simple IM back-end

[项目地址:https://github.com/wywwwwei/IMServer](https://github.com/wywwwwei/IMServer)

executable文件夹中是后端的可执行文件（测试时选一使用）

- cmd.exe 运行于Windows
- httpserver运行于Mac

## API文档

> 具体的服务端口可通过Service目录中config.go修改

### HTTP （端口6666）

- 获取完整个人信息

  请求

  ```json
  GET http://host:6666/profile/{userid}
  ```

  返回

  - userid不存在

    ```json
    Status Code:400
    {
        "msg":"Out of range"
    }
    ```

  - userid存在

    ```json
    Status Code:200
    {
        "user":-,
        "password":-,
        "name":-,
        "sex":-,
        "email":-,
        "signature":-
    }
    ```

- 获取好友列表（用户ID+用户名）

  请求：

  ```json
  GET http://host:6666/list/{userid}
  ```

  返回

  - userid不存在

    ```json
    Status Code:400
    {
        "msg":"Out of range"
    }
    ```

  - userid存在

    ```json
    Status Code:200
    [
        {"user":"-","name":"-"},
        {"user":"-","name":"-"}
    ]
    //user表示用户id
    //name表示对应的用户名
    ```

- 登陆验证

  请求：

  ```json
  POST http://host:6666/login
  
  //表单数据
  {
      "user":"-"	//用户id
      "password":"-" //登陆密码
  }
  ```

  返回

  - 密码正确

    ```json
    Status Code:200
    {
        "user":-,
        "password":-,
        "name":-,
        "sex":-,
        "email":-,
        "signature":-
    }
    ```

  - 密码错误

    ```json
    Status Code:400
    {
        "msg":"Out of range"	//用户不存在
    }
    {
        "msg":"Wrong userID or password" //密码错误
    }
    ```

- 用户注册

  请求

  ```json
  POST http://host:6666/regist
  
  //表单数据
  {
      "password":-,
      "name":-,
      "sex":-,
      "email":-,
      "signature":-
  }
  ```

  返回

  ```json
  Status Code:200
  {
      "user":-,			//需要根据返回的用户id进行登陆
      "password":-,
      "name":-,
      "sex":-,
      "email":-,
      "signature":-
  }
  ```

### TCP（端口8888）

TCP消息包格式

| Data Length |       Data        |
| :---------: | :---------------: |
|   4 Bytes   | Data Length Bytes |

```json
//具体的Data格式
{
    "type":"Message" or "Regist"
    "message":"-"
    "sender":"-"
    "receiver":"-"
    "createTime":double
}
```

- Regist包

  当客户端连接成功时发送至服务端，为当前用户注册tcp连接

  必填字段

  ```json
  {
      "type":"Regist",
      "sender":"current user id"
  }
  ```

- Message包

  这就是用户之间发送的消息

  所有字段都必须填