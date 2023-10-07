# Bifrost

Take you to the land of light, the city of freedom(A unified external service management system for NAS).

## 场景抽象

对于公网服务暴露，存在如下场景：
无公网：
1. 本地安装frpc，本地frpc配置/长期运行
2. 服务器安装frps，或者再添加nginx，服务端配置/长期运行
3. 域名配置/长期运行


其他本地和服务端的业务配置和安装


有公网：
1. 本地ngin反向代理配置
2. 无
3. 域名配置

从上面可以发现，都是一个一个步骤的抽象，上一个步骤完成，然后执行下一个步骤。所以，本质上是一个工作流的处理，携带着各种参数。

所以，本质上两种资源：步骤，和网站的服务流，服务类由一个一个步骤抽象组成。定义如下：

# demo
步骤1
```yaml
name: FrpcServiceUp
description: xxx
parameter:
  in:
    server:
      type: string
      description: xxx
      default: xxx
      required: true
    ports:
      type: array
      description: xxx
      required: true
      items:
        localPort:
          type: string
        remotePort:
          type: string
  out:
    configPath:
      type: string
image: xxxx
```

步骤2
```yaml
name: FrpsServiceUp
description: xxx
parameter:
  in:
    ip: x1
    sshUser: x2
    sshPasword: x3
    os: ubuntu
  out:
    configPath: xxx
image: xxxx
```

步骤3
```yaml
name: DDNS
parameter:
  in:
    xxx
  out:
    xxx
image:xxx
```

full workflow demo:
```yaml
name: exposeLocalService
kind: ServiceExpose
description: xxx
configurationSteps:
- use: xxx
  name: frpInstall
  in:
    ip: xxx
    xxx: xxx
- use: xxx
  in: |
    ip: frpInstall.out.xxx
- use: xxx
  name: xxx
  with:
    - xxx: xxx.out.xxx
- use: xxx
  with:
    -name: xxx
```
