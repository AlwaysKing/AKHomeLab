# 展示
![Example 1](/doc/1.png)
![Example 2](/doc/2.png)
![Example 3](/doc/3.png)

# 编译与部署
1. 编译网页
```
cd akhome
npm install
npm run build
```
2. 编译server
```
cd Server
./make.sh
```
3. 编译docker镜像
```
cd docker
./copy.sh
./make.sh
```
4. 创建容器  
    先将配置文件准备好放到指定目录，然后运行如下命令，注意替换 **"/host/path"** 为你实际放置配置的目录
```
docker run -d -p 9886:9886 -v /host/path:/data akhomelab
```
# 配置文件
以 **/data** 为根目录做演示
### 身份验证配置文件：/data/config.json
```
[
    {
        "Name":"Demo", 
        "Addr":"htttp://www.yourbarkapi.com/yourbarkkey"
    }
]
```
### 用户配置文件: **/data/对应用户名(Demo)/config.json**   
image 可以是文件名或者完整的url，如果是文件名则需要将图片放在 **"/data/对应用户名/icon"** 目录下
```
{
"Title":"Demo's Lab",
"Recent":[
  "google", "ChatGpt"
],  
"Apps":[
  {
    "Title": "搜索引擎",
    "Image": "search.png",
    "List": [
      { "AppName": "baidu", "URL": "https://www.baidu.com", "Image": "baidu.jpg" },
      { "AppName": "google", "URL": "https://www.google.com",  "Image": "google.png" },
      { "AppName": "Yandex", "URL": "https://yandex.com/",  "Image": "yandex.png" }
    ]
  },
  {
    "Title": "AI助手",
    "Image": "ai.png",
    "List": [
      { "AppName": "ChatGpt", "URL": "https://chat.openai.com/",  "Image": "chatgpt.png" },
      { "AppName": "Claude", "URL": "https://claude.ai/",  "Image": "claude.png" },
      { "AppName": "Kimi", "URL": "https://tongyi.aliyun.com/qianwen/", "Image": "tongyiqianwen.png" }
    ]
  }
]
}
```

完整的路径如下所示
```
-/data
    - config.json
    - Demo
      - config.json
      - icon
        - baidu.png
        - gogle.png
        - yandex.png
        ....
```