# Bifrost

Take you to the land of light, the city of freedom(A unified external service management system for NAS).

# demo
globalConfig demo:
```yaml
kind: config
name: globalConfig
description: xxx
data:
  ssh-password: xxx
  ssh-user: xxx
```

action element demo:
```yaml
kind: action
name: remoteInstallNginx
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


full workflow demo:
```yaml
name: exposeLocalService
kind: workflow
description: xxx
on: dispatch #dispatch/schedule
schedule: xxx
- use: nginxInstall
  name: remoteInstallNginx
  with:
    - name: ip
      value: 192.168.1.6
    - name: sshUser
      value: root
    - name: sshPassword
      valueFrom:
        kind: config
        name: globalConfig
        key: ssh-password
- use: configNginx
  para:
    - name: configPath
      valueFrom:
        kind: action
        name: remoteInstallNginx
        key: configPath
- use: action/frpServerInstall
  with:
    - name: ip
      value: 192.168.1.6
    - name: sshUser
      value: root
    - name: sshPassword
      value: root
- use: frpClientInstall
  with:
    - xxx
- use: action/exportLocalService
  with:
    - xxx
```
