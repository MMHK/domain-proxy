# Domain Proxy

增加内容代理域名工具，使内网服务拥有一个公共的外网2级域名，需要使用泛域名指向到宿主主机。

具体实现的功能如下:

- 自定义二/三级域名绑定到内网IP


## 编译
----
- 安装Golang环境, Go >= 1.13
- checkout 源码
- 在源码目录 执行` go mod vendor `签出所有的依赖库
- ` go build -o domain-proxy . ` 编译成二进制可执行文件
- 执行文件 ` domain-proxy -c ./config.json`

## 配置文件
----
该项目使用json文件进行配置，具体例子如下

```JS
{
  "listen": "{$HOST}", //服务绑定端口及地址
  "reload-cmd": "{$RELOAD}", //重启 webservice 服务的命令，可以是apache 或者 nginx 等服务
  "web_root": "{$ROOT}", //http 服务静态服务根目录
  "config-template-path": "{$TEMPLATE_PATH}", //配置的模板的路径
  "save-path": "{$SAVE_PATH}", // 生成的配置放置的路径
  "filename-format": "{{.Domain}}.conf" //配置命名格式
}
```


## 生成 `swagger` 文档

- 安装 [swagger-go](https://github.com/go-swagger/go-swagger)
- 在项目目录执行
```bash
swagger generate spec -o ./webroot/swagger/swagger.json
```
