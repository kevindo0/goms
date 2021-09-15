#### [jwtauth](https://stackoverflow.com/questions/31195469/golang-jwt-no-signature-crypto-rsa-verification-error)
```bash
openssl genrsa -out private.pem 1024
openssl rsa -in private.pem -pubout -out public.pem
go test ./datastructure/...  # 测试datastructure包下的所有测试用例
go test -v ./... # to test all packages within the current working directory, recursively
```

#### github
1. [leveldb](https://github.com/syndtr/goleveldb)
2. [fastcache](https://github.com/VictoriaMetrics/fastcache)
3. [cast(convert between different go types)](https://github.com/spf13/cast)
4. [viper(a complete configuration solution for Go applications)](https://github.com/spf13/viper)
5. [调试工具dlv](https://github.com/go-delve/delve)
