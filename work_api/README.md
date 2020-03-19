# work API
勤怠時間のCRUD API

# Endpoints

GET    /work_records
GET    /work_records/:id
POST   /work_records
PUT    /work_records/:id
DELETE /work_records/:id

## GET
```shell script
curl -X GET http://localhost:8080/work_records
```

```shell script
curl -X GET http://localhost:8080/work_records/:id
```

## POST
```shell script
curl -X POST http://localhost:8080/work_records \
-d '{
"work_date": "2014-10-10T00:00:00+09:00", 
"begin_work_time": "2014-10-10T10:00:00+09:00", 
"end_work_time": "2014-10-10T19:00:00+09:00", 
"begin_break_time": "2014-10-10T12:00:00+09:00",  
"end_break_time": "2014-10-10T13:00:00+09:00" 
}'
```

## PUT
```shell script
curl -X PUT http://localhost:8080/work_records/:id \
-d '{
"work_date": "2015-10-10T00:00:00+09:00", 
"begin_work_time": "2016-10-10T10:00:00+09:00", 
"end_work_time": "2017-10-10T19:00:00+09:00", 
"begin_break_time": "2018-10-10T12:00:00+09:00",  
"end_break_time": "2019-10-10T13:00:00+09:00" 
}'
```

## DELETE
```shell script
curl -X DELETE http://localhost:8080/work_records/:id
```

# TODO
- OPEN API
- echoを使う
