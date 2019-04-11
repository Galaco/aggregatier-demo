# Aggregatier

##### Aggregatier is a MOBA tierlist creator and aggregator.
This is really a simple demonstrative application for structuring a Vue front-end backed by a Go API.

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

* A CLI tool is provided to import Smite god data. It will require your own developer API credentials, as well as aws 
credentials for s3. Eventually this will be moved to the api. It is located at `cmd/smite/importheroes/main.go`

##### Production

###### @TODO