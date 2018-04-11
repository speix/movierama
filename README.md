# MovieRama
MovieRama is a simple app that supports a user system where users can share their favorite movies and express their opinion about each others submissions.  

- The backend (app folder) is implemented using Golang and exposes a RESTful API to interact with client requests.  
- The stateful layer consists of an SQLite database and an in-memory session storage using hashmaps.
- The frontend (static folder) is implemented using Vue js (https://vuejs.org/) and Bulma css framework (https://bulma.io/) to support the user interface.  
- Test coverage: hasher_test, validator_test, sessionData_test, movieHandler_test, userHandler_test, voteHandler_test
- The app can also be easily used with Docker (docker folder) setting a PORT through docker-compose.yml. (https://hub.docker.com/r/supergramm/movierama/)

### API
```
Endpoint:  /users/session

Action: User session status
Method: GET
Params: -

Action: User login
Method: POST
Params: email, password

Action: User logout
Method: DELETE
Params: -
```

```
Endpoint:  /users

Action: User registration
Method: POST
Params: email, password, name
```

```
Endpoint: /movies

Action: Submit movie
Method: POST
Params: title, description
```
```
Endpoint: /movies/{sorting}/{user_id}

Action: Sort movies
Method: GET
Params: {sorting}: likes, hates or date
Notes:  {user_id} is optional
```
```
Endpoint: /movies/{movie_id:[0-9]+/vote/{vote:[0-1]}

Action: Add vote to a movie
Method: POST
Params: {movie_id}, {vote}: 0 for negative or 1 for positive
```
```
Endpoint: /movies/{movie_id:[0-9]+/vote/

Action: Retract vote from movie
Method: DELETE
Params: {movie_id}
```

### Installation
Before using the app you need to setup the Server Port that will listen for http requests:
1. MR_SERVER_PORT

Make sure that no other application is using that port.

Select the way to run the app that suits you the most:

* Using the Dockerfile to build and run the container:
```
docker build -t supergramm/movierama:latest .
docker run -d -p PORT:PORT --name movierama -e MR_SERVER_PORT=you_selected_port supergramm/movierama
```


* Using Docker Compose tool:
```
docker-compose up -d
```


* Using go:
```
go get github.com/speix/movierama
go install github.com/speix/movierama

Dependencies (go get):
github.com/jmoiron/sqlx
github.com/mattn/go-sqlite3
github.com/twinj/uuid
github.com/gorilla/mux
```


### Docker image hosted on DockerHub
https://hub.docker.com/r/supergramm/movierama/

### Preview
The app is deployed on the Elastic Container Service of AWS:  
http://ec2-18-196-125-219.eu-central-1.compute.amazonaws.com