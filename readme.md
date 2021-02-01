# COMMON FOR GO

1. 压力测试

``` zsh
# 压力测试 -count 表示执行5次 -v 打印信息 -test.timeout 设置测试超时时间
# -cpu=2 指定cpu核数 默认不指定即最大 也可以指定多种 -cpu=2,4,6
# -benchtime 可以指定时间 也可以指定次数  -benchtime=5s  -benchtime=100x 100次
go test -test.bench=".*" -count=5 -v
# 测试指定压力方法
go test -test.bench testfuncname
# 测试指定通过方法
go test -test.run testfuncname
```