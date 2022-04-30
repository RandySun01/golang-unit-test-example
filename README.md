# golang-unit-test-example
Go单元测试

执行单元测试: go test

查看测试函数名称和运行时间: go test -v

正则匹配测试函数: go test -v -run="More"

正则匹配子测试函数: go test -v -run="ChildrenSplit/leading_sep"

跳过某些测试用例: go test -short

查看测试覆盖率： go test -cover

测试覆盖率生成文件: go test -cover -coverprofile=test_cover.out // 文件名

测试覆盖率生成html报告: go tool cover -html=test_cover.out // 生成的文件名

基准测试函数执行: go test -bench=Split  // 基准测试用例名字匹配

基准测试获取内存分配的统计数据: go test -bench=Split -benchmem

基准测试指定执行时间: go test -bench=Fib40 -benchtime=20s

指定cup执行基准测试: go test -bench=Para -cpu 8

示例函数执行: go test -run ExampleS // 实例函数名称
