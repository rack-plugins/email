# email
send email using http API

# usage


```bash
curl -XPOST 127.0.0.1:8000/mailto \
-H "content-type: application/json" \
-d '{
    "mailto": [
        "lmr@epurs.com"
    ],
    "subject": "test subject",
    "body": "测试发送成功！"
}'
```