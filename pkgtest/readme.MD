#### [jwtauth](https://stackoverflow.com/questions/31195469/golang-jwt-no-signature-crypto-rsa-verification-error)
```bash
openssl genrsa -out private.pem 1024
openssl rsa -in private.pem -pubout -out public.pem
go test ./datastructure/...  # 测试datastructure包下的所有测试用例
go test -v ./... # to test all packages within the current working directory, recursively
```
