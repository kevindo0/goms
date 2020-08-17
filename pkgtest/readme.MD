#### [jwtauth](https://stackoverflow.com/questions/31195469/golang-jwt-no-signature-crypto-rsa-verification-error)
```bash
openssl genrsa -out private.pem 1024
openssl rsa -in private.pem -pubout -out public.pem
```
