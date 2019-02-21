# Aggregatier

##### Aggregatier is a MOBA tierlist creator and aggregator.



### Running
Docker is required. It may be run without, but that is left up to you. 

#### Locally
Run the server using docker composer. This will spin up the server on `localhost:8080`.
```bash
    docker-compose up -d
```
Run the client app using npm. This will be available on `localhost:8082`
```bash
cd client
npm run-script serve
```

You should be presented with something a little like this:
![Homepage](https://raw.githubusercontent.com/Galaco/aggregatier/master/.github/docs/repo/aggregatier.png)


#### Developing

Dev environments should automatically reload on code changes. `docker-compose logs -f -t web` recommended
to monitor.

Mysql & PHPMyAdmin are included in `docker-compose.override.yml`. `aggregatier.sql` is a useable dump for
development until proper seeding is implemented. Just run the dump to setup the database.

A CLI tool is provided to import Smite god data. It will require your own developer API credentials. Eventually
this will be moved to the api. It is located at `server/cmd/smite/importheroes/main.go`