# Anti-bruteforce service

## General description

The service is designed to combat the selection of passwords when logging in to any system.
The service is called before user authorization and can either allow or block the attempt.
It is assumed that the service is used only for server-server, i.e. hidden from the end user.

## Algorithm of operation

The service limits the frequency of authorization attempts for various combinations of parameters, for example:
* no more than N = 10 attempts per minute for this login.
* no more than M = 100 attempts per minute for this password (reverse brute-force protection).
* no more than K = 1000 attempts per minute for this IP (the number is large, because NAT).

White/black sheets contain lists of network addresses that are processed in a simpler way.
If the incoming ip is in the whitelist, the service definitely allows authorization (ok=true), if it is in the blacklist, it rejects it (ok=false).

## Architecture

The microservice consists of an API, a database for storing settings and black/white lists.

## Description of API methods

### Authorization attempt
Request:
* login
* password
* ip

Answer:
* ok (true/false) - the service should return ok=true if it considers that the request is normal
  and ok=false if he thinks bruteforce is going on.

### Reset bucket
* login
* ip

Must clear the buckets corresponding to the transmitted login and ip

### Adding IP to blacklist
* subnet (ip + mask)

### Removing IP from blacklist
* subnet (ip + mask)

### Adding IP to whitelist
* subnet (ip + mask)

### Removing IP from whitelist
* subnet (ip + mask)

## Configuration
The main configuration parameters: N, M, K are the limits for reaching which the service considers an attempt to be a brute force.

## Command-Line interface
Developed a command-line interface for manual administration of the service.
Through the CLI, it is possible to cause a bucket reset and manage the whitelist/blacklist.
The CLI works through the GRPC interface.

## Deployment
Before deployment, you need to add the `.config.yaml` file to the project directory. Example file `.config.yaml.example`
The deployment of the microservice is carried out by the `make up` command in the directory with the project.

## Testing
The project contains unit tests of leaky-bucket implementation and integration tests for API methods.
The deployment of the microservice is carried out by the 'make test` command in the directory with the project.

## Monitoring
When you start the build in the docker, prometheus rises
To enter prometheus, you need to follow the link: http://localhost:9090
View the status: http://localhost:9090/targets