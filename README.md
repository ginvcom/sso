# sso
基于rbac的单点系统

这是一套后台单独管理系统，用于解决1个以上的后台账号登录和授权管理。ginv默认使用前后端分离技术。服务端采用golang开发，基于go-zero框架，前端才用vue3开发，基于ant design vue。

sso具有auth、ssoms、和gateway三个系统，分别具有以下功能：
```
auth： 负责系统登录验证、页面权限验证、接口授权验证
ssoms: 负责角色、用户管理、系统和菜单管理、角色授权
gateway：登录登出页面、负责接入请求和转发请求
```

Ginv SSO的多个后台系统需是在同一个域名下的子域名。这样做的目的是使用安全的httpOnly的Cookie在多个系统之间进行用户身份识别。

## 特色

根据go-zero的goctl创建的api规则, 将请求按resetful风格进行权限校验，前端可以通过goctl的api文件生成ts，然后请求gateway。

### 关于前端请求
前端请求的入口，接收的请求前端请求必须具有如下信息：
headers
| key | 说明 |
|:--|:--|
| cookie | 登录的后生成的cookie, 由gateway的登录后的跳转页面提供 |
| Refere | 告诉服务端由哪个页面发起的请求，这边使用unsafe-url配置 |
| x-client-env | 环境，可选值有online\|gray |
| x-client-service | 请求的服务端的服务名 |
| x-client-system | 后台系统的系统名 |
| x-client-uri | 后台系统的系统名 |


### 关于gateway：
gateway主要负责转发功能，为了保证生成Cookie由服务端下发的页面创建，网关还提供了登录、登出和登录跳转页(跳转到相应的系统)。

gateway通过请求的headers[x-client-service], 知道转发给哪个服务，在转发请求之前，会请求auth服务，根据cookie, 会解析到用户的信息，根据x-client-system和x-client-uri
则能可以知道是否具有权限。









 
 

