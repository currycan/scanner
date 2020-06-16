# scanner

> 扫描特定服务端口是否联通的小工具。

## 文档

```text
A tool that check the connectivity of specific server or servers,
you can set a yaml config file for multiple services or use command line for a single server.
For example:
	scanner -c /path/to/example/scanner.yaml
	scanner --host 10.0.0.1 --port 6379 --name redis --project test
If you want to open the debug mode,you can add '--debug true in the end of command

Usage:
  scanner [flags]
  scanner [command]

Available Commands:
  help        Help about any command
  version     show scanner version

Flags:
  -c, --config string    config file path
      --debug            debug mode
  -h, --help             help for scanner
      --host string      the IP or domain name of the server
      --name string      the name of the server
      --port int         the port of the server
      --project string   the project of the server

Use "scanner [command] --help" for more information about a command.
```

### 说明

支持两种模式:
- 通过配置文件批量检测是否联通，配置文件参见[example](example/scanner.yaml)
- 命令行模式：scanner --host 10.0.0.1 --port 6379 --name redis --project test

其他说明：
+ 使用配置文件批量扫描端口支持热更新， 修改完配置文件，会自动重新扫描

### 部署
使用 k8s daemonSet 方式部署，配置清单参见[deploy](deploy/deploy.yaml)
 