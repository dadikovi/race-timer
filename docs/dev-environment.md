# Development environment

**What do you need before you would start to contribute to the project?**

# Start required external services

1. Install Docker
2.  ```
    cd environment
    docker-compose up -d
    ```
This starts up the database and to o11y stack. Access the latter:

[Grafana](localhost:3000) (admin / admin) 

# Start server and clients

Ofc you can start only the server (and maybe one of the clients if you want to develop that)

## Admin client (React application)

1. Install React
2. `npm start`

## Scanner client (Android application)

1. Install [Flutter](flutter.dev)
2. Install [Android studio](https://developer.android.com/studio)
3. Install `Flutter` and `Dart` plugins for `Android studio`
4. Run `lib/main.dart`, for help check [this](https://flutter.dev/docs/get-started/test-drive?tab=androidstudio)

## Server (Go)

1. Install [Go](https://golang.org/doc/install)
3. Set DB access:
    ```bash
    export RACE_TIMER_DB_USER=postgres
    export RACE_TIMER_DB_NAME=postgres
    export RACE_TIMER_DB_PASSWORD=mysecretpassword
    ```

## Run server tests

```bash
cd server # If you are in the root directory of this repo
go test -v
```

# How to connect to database
After you initiated the test database:

```bash
docker exec -it {docker container id} bash
psql -U postgres -d postgres
```