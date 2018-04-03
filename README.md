# National Parks

Restful API for US National Park Service (unofficial). The resources are the
states and their parks. Additional features include pagination and
rate-limiting. The API is written with minimal dependencies (the
only external dependency is the router).

Technology
----------
* Go
* PostgreSQL

Endpoints
---------

| Method     | URI                             | Action                                |
|------------|---------------------------------|---------------------------------------|
| `GET`      | `/api/states`                   | `Retrieve all states`<sup>1</sup>     |
| `GET`      | `/api/states/{sid}`             | `Retrieve state`                      |
| `POST`     | `/api/states`                   | `Add state`                           |
| `PUT`      | `/api/states/{sid}`             | `Update state`                        |
| `DELETE`   | `/api/states/{sid}`             | `Remove state`                        |
| `GET`      | `/api/states/{sid}/parks`       | `Retrieve all state parks`<sup>1</sup>|
| `GET`      | `/api/states/{sid}/parks/{id}`  | `Retrieve state park`                 |
| `POST`     | `/api/states/{sid}/parks`       | `Add state park`                      |
| `PATCH`    | `/api/states/{sid}/parks/{id}`  | `Partially update state park`<sup>2</sup>         |
| `DELETE`   | `/api/states/{sid}/parks/{id}`  | `Delete state's park`                 |

1. Optional query parameters: count, start
2. Send only the fields you need to update (or all fields for full updates)

Sample Responses
---------------

`http get http://localhost:5000/api/states`
```
{
    "data": [
        {
            "id": 1, 
            "name": "California"
        }, 
        {
            "id": 2, 
            "name": "Arizona"
        }, 
        {
            "id": 3, 
            "name": "New York"
        }, 
...
```
`http get http://localhost:5000/api/states/1/parks`
```
{
    "data": [
        {
            "description": "Recognized by its granite cliffs, waterfalls, clear streams, giant sequoia groves, lakes, mountains, glaciers, and biological diversity.", 
            "id": 1, 
            "name": "Yosemite", 
            "nearestCity": "Mariposa", 
            "stateId": 1, 
            "visitors": 4336890
        }, 
...
```
Run
---
With docker:
```
docker-compose build
docker-compose up -d db
docker-compose up
Go to http://localhost:5000 and visit one of the above endpoints
```

Alternatively, create a database named 'national_parks', run the migration
scripts (located in `./migrations/`), and then open `main.go` and
point the PostgreSQL URI to your server.

`cd` into `./national-parks` (if you are not already); then run:
```
go build
./national-parks
Go to http://localhost:5000 and visit one of the above endpoints
```

TODO
---
Add hypermedia links and meta data  
Add sort (by visitors and name)  
Add number of parks per state  
Add more unit tests
