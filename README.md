# Aggregatier

##### Aggregatier is (planned to be) a MOBA tierlist creator and aggregator.
This is really a simple demonstrative application for structuring a Vue front-end backed by a Go API. It doesn't actually
do very much. It is a demo Vue front end that communicates with a Go api to fetch and display some data from a MySQL 
database. Client and API come with docker configuration, and kubernetes configuration is coming too.

If a feature is broken or incomplete, its probably intentional.

### Running
Docker is required. It can be run without, but that is left up to you.

#### Locally
##### Development
1. Copy `config.example.json` to `config.json`.
2. Run the application using docker-compose. This will spin up the server on `localhost:8080`, and the client on `localhost:8081`.
```bash
    docker-compose up -d
```

###### You should be presented with something a little like this
![Homepage](https://raw.githubusercontent.com/Galaco/aggregatier/master/.github/docs/repo/aggregatier.png)

* A development MySQL database is available on `localhost:3306`. PhpMyAdmin is available at `localhost:8089`. For now, 
import the schema `docker/dev/database/sources.sql` into database `aggregatier` until a better solution.

* A CLI tool is provided to import Smite god data. It will require your own developer API credentials. Eventually
this will be moved to the api. It is located at `server/cmd/smite/importheroes/main.go`

##### Production

###### @TODO