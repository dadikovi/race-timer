# Development environment
** What do you need before you would start to contribute to the project?**

## Scanner client (Android application)

1. Install [Flutter](flutter.dev)
2. Install [Android studio](https://developer.android.com/studio)
3. Install `Flutter` and `Dart` plugins for `Android studio`
4. Run `lib/main.dart`, for help check [this](https://flutter.dev/docs/get-started/test-drive?tab=androidstudio)

## Server (Go)

1. Install [Go](https://golang.org/doc/install)
2. Install Docker
3. Setup PostgreSQL: 
    ```bash
    docker ps --format {{.ID}} | xargs docker stop {} # Stop all docker containers 
    docker run -it -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres

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