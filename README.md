# flamingo-readinglist [![Build Status](https://github.com/tessig/flamingo-readinglist/workflows/CI/badge.svg)](https://github.com/tessig/flamingo-readinglist/actions?query=workflow%3ACI)
A simple [Flamingo](https://www.flamingo.me/) application as showcase.

## Building the app

```bash
go build -o readinglist .
```

## Development setup

A docker-compose setup for development can be found in devenv directory.

This contains:
  * Article service on http://localhost:8000/ 
  * Jaeger on http://localhost:16686/
  * Prometheus on http://localhost:9090/
  * Grafana on http://localhost:3000/ 
  
    user: admin, password: password
  * Alert Manager on http://localhost:9093/
  * Node Exporter on http://localhost:9100/
  * cAdvisor on http://localhost:8080/
 
  
Simply run `docker-compose up` from within devenv.

Then start the app via `CONTEXT=dev go run main.go serve`

The config in `config/config_dev.cue` matches the docker-compose setup.

The app will be under http://localhost:3322/
The metrics endpoint will be under http://localhost:13210/metrics

## PHP Service

The source code of the PHP service that parses the golang blog can be founde here: https://github.com/carstendietrich/php-golang-blog-parser
