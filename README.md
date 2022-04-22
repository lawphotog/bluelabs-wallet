# bluelabs

This service manages user wallets. This is not supposed to be production ready service. 
This is proof of concept with all the requirement implemented with limited time available.

various aspects of this service can be improved especially error handling and validation and other edge cases that might have. error logging and monitoring is also something outside the scope.

Authentication authorisation is also something I considerred outside the scope

How to run it

set up db first. docker-compose.yml is provided to set up local dynamodb

```docker compose up```

this should bring up both dynamodb and the service on a machine with latest version of Docker installed. 

VSCode

You can open this in VSCode and F5 to run the project.

userId needs to be provided to this service for everything.
HTTP is used for simplicity in this service but may not be the most suitable one depends on other requirements within the organisation.

POST
localhost:8080/wallet/create/{userId}
localhost:8080/wallet/deposit/{userId}/{amount}
localhost:8080/wallet/withdraw/{userId}/{amount}

The wallet needs to be created first before other sequence requests can be made.

GET
localhost:8080/wallet/getbalance/{userId}

How to run tests

`go test ./...`

Main business logics (wallet.go) are unit tested for handling transactions. I have considered integration tests, performance tests, etc to be outside of the scope.

Assumptions

the client will be javascript/http client.
should this required from another Go client, this can be done using more efficient protocol such as GRPC.
Or even have this service reads from a Kafka topic for resiliency.

DB

I am using Dynamodb in this example for ease of setting up. This solution can be done in any databases including SQL databases.
Race condition are handled with optimistic locking using Condition Expression feature in dynamodb.

Libraries

gin is used as http server for ease of setup
testify/mock and ginkgo is used for unit testing

