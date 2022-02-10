Signup:
```bash
curl --location --request POST 'localhost:8010/user/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test3",
    "password": "12345678"
}'
```



Login:
```bash
curl --location --request POST 'localhost:8010/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test3",
    "password": "12345678"
}'
```



Register URL:
```bash
curl --location --request POST 'localhost:8010/url/register-url' \
--header 'Authorization: Bearer jwt token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "snapp",
    "url": "https://snapp.ir",
    "threshold": "5"
}'
```



Get URLs:
```bash
curl --location --request GET 'localhost:8010/url/urls' \
--header 'Authorization: Bearer jwt token'
```



Statistics:
```bash
curl --location --request POST 'localhost:8010/url/stats' \
--header 'Authorization: Bearer jwt token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "snapp"
}'
```



Alerts:
```bash
curl --location --request POST 'localhost:8010/alert/alerts' \
--header 'Authorization: Bearer jwt token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "urlname": "google"
}'
```
