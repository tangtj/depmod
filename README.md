# Depmod
基于干预CI/CD流程，做到对go微服务互相依赖的版本管控。可以看这个文章，我的一点小想法，
[Go微服务之间模块依赖管理的思考](https://tangtj.cn/post/2021-07-18/Go%E5%BE%AE%E6%9C%8D%E5%8A%A1%E4%B9%8B%E9%97%B4%E6%A8%A1%E5%9D%97%E4%BE%9D%E8%B5%96%E7%AE%A1%E7%90%86%E7%9A%84%E6%80%9D%E8%80%83)

## 使用
1. 下载
```shell
go install github.com/tangtj/depmod
```

2. 校验
```shell
depmod -c config.json -m go.mod
```
-c 配置文件路径
-m 需要校验的 go.mod 文件路径