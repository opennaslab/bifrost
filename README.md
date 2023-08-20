# nasone

Take you to the land of light, the city of freedom(A unified external service management system for NAS).

# demo

action element demo:
```yaml
kind: action
name: remoteInstallNginx
description: xxx
parameter:
    ip: x1
    sshUser: x2
    sshPasword: x3
    os: ubuntu
image: xxxx
```


full workflow demo:
```yaml
name: exposeLocalService
kind: workflow
on: dispatch #dispatch/period/trigger
- use: action/nginxInstall
  with:
    ip: 192.168.1.6
    sshUser: root
    sshPassword: root
    os: ubuntu
- use: action/frpServerInstall
  with:
    ip: 192.168.1.6
    sshUser: root
    sshPassword: root
    os: ubuntu
- use: action/frpClientInstall
  with:
    ip: 192.168.1.2
    sshUser: root
    sshPassword: root
    os: ubuntu
- use: action/exportLocalService
  with:
    localPort: 45
    remotePort: 80
    dns: video.opennaslab.com
```
