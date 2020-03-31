#### 1. app1
```bash
cd app1/config
go test -v
cd app1/api
go test -v
# 性能测试
go test -bench="." -v
```

#### 2. combin
```bash
cd combin
go test combination_test.go combination.go           # 单元测试
go test --cover combination_test.go combination.go   # 单元测试覆盖率
go test -bench=. combination_test.go combination.go  # 性能测试
```

#### 3. Gomock
```bash
mockgen -source=./person/male.go -destination=./mock/male_mock.go -package=mock 
```

#### tools
1. [goconvey](github.com/smartystreets/goconvey)