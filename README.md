# Golang Test

We are PackAndGo, a small bus company. We want to create a REST API that helps us manage the trips that we offer.

## Goal

Your goal is to **build a REST API** with the following characteristics:

| Method | Endpoint  | Description          |
|--------|-----------|----------------------|
| GET    | /trip     | List all trips       |
| POST   | /trip     | Add a new trip       |
| GET    | /trip/:id | Get trip with ID :id |

### REST API Description

We would like the trips to be obtained in the following format:

```json
{
    origin: "Barcelona",
    destination: "Seville",
    dates: "Mon Tue Wed Fri",
    price: 40.55
}
```

Whereas to add a trip, we would send the following:

```json
{
    originId: 2,
    destinationId: 1,
    dates: "Sat Sun",
    price: 40.55
}
```

The **trip ID** should be added automatically. Each trip should have a unique trip ID.

The **list of cities** is in a text file, *cities.txt*, but perhaps we will change that in the future as our company grows. Right now, every line in the text file is a city. The first line corresponds to cityId=1, the second to cityId=2, etc.

### General guidelines

We want you to build a REST API that works with our current needs, but that can be ready to **change easily in the future**, without having to rewrite the whole system or fearing that something will break. So please, try to make it as future-proof as possible.

Use the packages that you think are suitable for the job, as well as the **architecture and code structure** that makes most sense from your point of view. Feel free to move, split, etc. the provided files into the files and folders of your choice.

You can also challenge and *change the proposed API structure if you feel it is necessary*, as long as you give a reason why you have decided to do things in a different way.

## Changes!

First of all, i make some changes in API:

* __GET /trip/:id__:
The returned element will contiains two objects, _origin_ and _destination_, this will show the ID to the client, this let the user knows the Id for the POST operation:
```json
{
    "id": 1,
    "origin": {
        "id": 1,
        "name": "Barcelona"
    },
    "destination": {
        "id": 2,
        "name": "Seville"
    },
    "dates": "Mon Tue Wed Fri",
    "price": 40.55
}
```

* __POST /trip__:
The same way as GET, just to create a "standar" way of Data, to make this request you can send something like:
```json
{
    "origin": {
        "id": 1
    },
    "destination": {
        "id": 2
    },
    "dates": "Mon Tue Wed Fri",
    "price": 40.55
}
```

* __GET /trip__:
I add some queryParams to make List paginated, this is really important to manage big amounts of trips (not now but in future). And will return the list of trips, the total count of elements and the limit you request. By default this querys are _limit=100_ and _start=0_
_http://localhost:8090/trip?limit=10&start=0_
```json
{
    "trips": [
        {
            "id": 1,
            "origin": {
                "id": 1,
                "name": "Barcelona"
            },
            "destination": {
                "id": 2,
                "name": "Seville"
            },
            "dates": "Mon Tue Wed Fri",
            "price": 40.55
        },
        {
            "id": 2,
            "origin": {
                "id": 2,
                "name": "Seville"
            },
            "destination": {
                "id": 1,
                "name": "Barcelona"
            },
            "dates": "Sat Sun",
            "price": 40.55
        },
        {
            "id": 3,
            "origin": {
                "id": 3,
                "name": "Madrid"
            },
            "destination": {
                "id": 6,
                "name": "Malaga"
            },
            "dates": "Mon Tue Wed Thu Fri",
            "price": 32.1
        }
    ],
    "total": 3,
    "limit": 10
}
```
* __GET /health__:
I add this new endpoint because is really usefull when you deploy it, just to be sure your service is up and ready to recieve request.
_http://localhost:8090/health_
```json
{}
```

## App structure
This application have been structured thinking of future changes, this is the schema:
```
.
├── cmd							//Package of main aplications
│   ├── client.go  				//Client for acceptance test
│   └── main.go 				//Entry point of app
├── docker-compose.yml 			//Definition to stand up everything and make automatic request
├── Dockerfile 					//App docker definition
├── DockerfileClient 			//Client container definition
├── go.mod
├── go.sum
├── Makefile 					//To make commands like build or test in more abstract way
├── pkg 						//App principal packages
│   ├── api 					//API public structs
│   │   └── model.go
│   ├── health 					//The health handler
│   │   └── health.go
│   ├── server 					//Logic with server and mapping handler
│   │   ├── handler.go 			//Get requests and convert to data for application logic
│   │   ├── handler_test.go
│   │   └── server.go 
│   └── trip 					//Main logic
│       ├── city.go 			//City Implemtations, get the city names
│       ├── city_test.go
│       ├── trip.go 			//Trip Implementations
│       └── trip_test.go
├── README.md
└── resources 					//resource files, like cities file and configuration
    ├── cities.txt
    └── config.env 
```

## How to run

You can configure the application port using `/resources/config.env` file.

if you have make installed in your computer can type:
```
$ make build
```
To create the _trip_ bin.

If just want to test everything and stand up the service you can type:
```
$ docker-compose up
```

This will stand up the application service, and run one client who list the elements, add one trip, list back again, and will get the added trip.