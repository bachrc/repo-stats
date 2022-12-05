# RepoStats

This utility, written in Golang and TDD with acceptance tests, can fetch data on a sample of 100 repositories.

## Installation

You can first clone the project in your desired folder :

```shell
git clone https://github.com/bachrc/repo-stats # Clone the repository
cd repo-stats # Go inside the freshly created
```

### With go cli
This project needs your Go version to be 1.17+.

You can clone the repository :

```shell
<parameters> go run github.com/bachrc/repo-stats/cmd/web # And then you launch the web server
```

### With docker-compose

In the [docker-compose.yml](./docker-compose.yml) file, you'll need to fill the required parameters [specified below](#parameters).

When done, you can launch the application :

```shell
docker-compose up -d
```

### Parameters

To make the application run smoothly, you'll need to provide some environment variables :

- `GITHUBACCESSTOKEN` : This parameter is **required**. Because of Github public API rate limits, 
you'll need to [generate a token](https://github.com/settings/tokens) with your Github account.
- `PORT` : The port on which the webserver will run.

## Documentation

The API description is available as an [OpenAPI file](https://editor.swagger.io/?url=https://raw.githubusercontent.com/bachrc/repo-stats/master/api/openapi.yml).

## Architecture

The description of the architecture is available in the [ARCHITECTURE.md](./ARCHITECTURE.md) file.

A little more context about the realisation of this exercise can be found (in French) in the [CONTEXTE-EXERCICE.md](./CONTEXTE-EXERCICE.md) file. 
