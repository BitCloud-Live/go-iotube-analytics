# Go Iotube analytics
fast, lightweight [Iotube](https://tube.iotex.io/) bridge analytics and API.
![Flow](assets/flow.png "Flow")

## Quickstart using docker-compose
```sh
$ cp env.example .env # edit and add postgres credentials
```
run using docker-compose:
```sh
$ docker-compose up -d --build # deploy and run api, defiscrapers and postgres all at once
```
## Deployment using k8s manifest files.