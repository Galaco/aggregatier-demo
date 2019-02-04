# Aggregatier

##### Aggregatier is a MOBA tierlist creator and aggregator.



### Running
Docker is required. It may be run without, but that is left up to the reader.

#### Locally
```bash
    docker-compose up -d
```



#### Developing

Dev environments should automatically reload on code changes. `docker-compose logs -f -t web` recommended
to monitor.

Mysql & PHPmyadmin are included in `docker-compose.override.yml`. `aggregatier.sql` is a useable dump for
development until proper seeding is implemented. Just create the database `aggregatier` and run the dump
manually for now.