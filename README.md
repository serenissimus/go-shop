# Go Shop

## Environment setup

You need to have

* [Go](https://golang.org/),
* [Node.js](https://nodejs.org/),
* [Yarn](https://yarnpkg.com/),
* [Docker](https://www.docker.com/), and
* [Docker Compose](https://docs.docker.com/compose/)

installed on your computer.

Verify the tools by running the following commands:

```sh
go version
yarn --version
docker --version
docker-compose --version
```

If you are using Windows you will also need [gcc](https://gcc.gnu.org/). It comes
installed on Mac and almost all Linux distributions.

## Start in development mode

In the **docker** directory run the command (you might need to prepend it with
`sudo` depending on your setup):
```sh
docker-compose -f docker-compose-dev.yml up
```

This starts a local PostgreSQL on `localhost:5432`.
The DB will be populated with test records from the [init-db.sql](init-db.sql) file.

Navigate to the `server` folder and start the back end:

```sh
cd server
make run
```
The back end will serve on http://localhost:8080.

Navigate to the `webapp` folder, install dependencies, and start the front end
development server by running:

```sh
cd webapp
yarn
yarn start
```
The application will be available on http://localhost:3000.

## Start in production mode

Perform in the **docker**:
```sh
docker-compose up
```
This will build the application and start it together with its database. Access the
application on http://localhost:8080.
