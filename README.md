# Drone Navigation Service (DNS)

This is a drone navigation service whose primary goal is to help a drone retrieve the location of a databank. 

## Assumptions
- There are sectors of galaxies (referred to `sectors` here) that when passed coordinates (`x,y and z`) and velocity, can return locations of databanks on these sectors.
- A `JWT` is used to authenticate requests to the endpoint.
- A `droneID` is part of the `JWT` claims. 
- The `droneID` can then be used to retrieve the drone details (from a database) which contain most importantly, the `sectorID` needed to calculate the location. This retrieval is mocked in this service. 

## Likely Use Cases

- A drone can be embedded with a JWT that can be used to retrieve location of a databank
- A drone can authenticate with DNS, generate a JWT that can be used to make a call to retrieve location of a databank


## Usage
- Make a `GET` call to `/locations` with query parameters (coords and velocity). Please see postman collection in `postman` folder
- Pass a bearer token containing `droneID` as part of its claims. You can generate one here - http://jwtbuilder.jamiekurtz.com/
- Please note that retrieval of drone details is mocked in `internal/db.go` to return a constant value. A furistic implementation would be to use a relational database where `sectors` are stored (with the `id` set as primary key) and then, the function in `internal\db.go` can query the `DB` to return the `sectorID`
- This endpoint should return a response in this form
``` bash
{
    "loc": 113.23999999999998
}
```
### Build
``` bash
$  go build -o main ./cmd/drones  
``` 
### Running Tests
``` bash
$ make test
``` 

### Run
``` bash 
$ make run 
```

### Docker
``` bash 
$ docker build -t drones-api .    
```

``` bash 
$ docker run -p 8000:9064 drones-api
```

## OpenAPI Spec
The OpenAPI spec for this service can be found in the `open-api.yaml` file. Upload to https://editor.swagger.io/ to view. 

## Postman Collection
The Postman collection can be found in the `postman` folder

