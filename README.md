# Depmod
基于干预CI/CD流程，做到对go微服务互相依赖的版本管控。可以看这个文章，我的一点小想法，
[Go微服务之间模块依赖管理的思考](https://tangtj.cn/post/2021-07-18/Go%E5%BE%AE%E6%9C%8D%E5%8A%A1%E4%B9%8B%E9%97%B4%E6%A8%A1%E5%9D%97%E4%BE%9D%E8%B5%96%E7%AE%A1%E7%90%86%E7%9A%84%E6%80%9D%E8%80%83)

## 思路
集成在CI的流程中。预先解析项目的go.mod文件，检查是否有存在废弃的依赖、是否有引用低于标准的依赖。发现非预期依赖直接报错，停止编译流程。促使团队内成员，去升级依赖或移除依赖。保证团队微服务互相引用中没有过时的依赖。

### 依赖配置
使用一个外置json格式的配置文件用作，对依赖的管理的配置。[schema参考](/config.schema.json)

人为维护一个config，用于配置每个依赖的最低版本，是否被弃用。

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