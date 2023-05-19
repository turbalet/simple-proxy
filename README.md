## Endpoints
### Do proxy things
```
POST /jobs
{
  "method": "GET",
  "url": "http://google.com",
  "headers": {
    "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
    ...
   }
}
```
### Response
```
{
  "id": "generated unique id",
  "status": "status code from 3rd party service",
  "length": "content length from 3rd party service",
  "headers": "headers list from third party service"
}
```

### Request saved job
```
GET /jobs/{ID}
```
### Response 
```
{
  "id": "generated unique id",
  "status": "status code from 3rd party service",
  "length": "content length from 3rd party service",
  "method": "GET",
  "url": "http://google.com",
  "req_headers": {
    "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
    ...
   },
  "resp_headers": "headers list from third party service"
}
```
