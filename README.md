# goid 发号器平台

单调递增分段发号服务`leaf`, 使用主备模式, 避免单点故障

提供发号器平台后台管理配置, 指定机器节点提供相应发号业务

## layout

[servant](https://github.com/ByronLiang/servant)

### Use

配置环境与文件路径

```yaml
environment:
      - CONFIG_PATH=/var/platform_id
      - CONFIG_ENV=dev
```

使用挂载目录, 将宿主机文件目录挂载进入容器目录

```yaml
volumes:
 - /var/platform_id:/var/platform_id
```


