# DOCKER CONTAINER MANAGEMENT SYSTEM
## Author : Shreevatshan B

This is a simple docker container management system. 
It is a simple web application that can be used to manage docker containers and view logs.
The application and containers must be running on the same host.

# steps to run the application
## 1. build
go build
## 2. run, pass the addr in which you would like to run the application
./docker host:port
## example:
./docker localhost:9090
## 3. access the application
http://localhost:9090/docker
