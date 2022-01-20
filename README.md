# adService

## Prerequisites

- **[Docker][]**:  latest version.

## Set up servers

To run the servers use [Docker-compose]

```console
docker-compose up
```

[Docker]: https://docs.docker.com
[Docker-compose]: https://docs.docker.com/compose/

## Healthcheck

### Request
    GET api/ping

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 54

    {
    "message": "pong",
    "serviceUpTime": "14m2s"
    }

## Get list of ad

### Request

`GET api/adList?orderField=price&order=DESC` 

### QueryParams
    orderField = price | create_ts
    order = DESC | ASC

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 54

    [
        {
        "name": "test2",
        "price": 1000,
        "image_links": [
        "ttt2"
        ]
    },
        {
        "name": "test3",
        "price": 800,
        "image_links": [
        "ttt3"
        ]
    },
    {
        "name": "test1",
        "price": 500,
        "image_links": [
        "ttt1"
        ]
    },
    ]

## Get current ad by id

### Request

`GET api/ad/:id` - 

### BodyParams
    ["images","description"]
    -- images - get ad images list
    -- description - get description  
### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 54
    {
        "name":"test1",
        "price":500,
        "description":"testDesc"
        "image_links":[
            "ttt1"
        ]
    },

## Add new ad

### Request

`POST api/ad/`

### BodyParams
    {
    "name":"test",
    "description":"description",
    "price":20
    }   
### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 54
    {
        "id": 8
    }
