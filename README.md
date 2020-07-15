# Planet Express

A fully functional end-to-end app with everything except a DB. To see it in action, run `make` and then the following three commands, each in a separate terminal and all from the root directory

```sh
./ship/ship
./headquarters/headquarters
yarn --cwd ./dashboard start
```

Aside from the tasks accomplished below, I also:
* Refactored the `main.go` files into a few different files depending on responsibilities
* Implemented the `GetShip` `GetCrew`, `GetDelivery` and `ListDeliveries` methods all in GraphQL, with `GetDelivery` searching based on UUID (only `GetShip` is shown in the React UI)
* Added a `CrewMember` resource to specify the members of the `Crew`

### Tasks Completed

#### Phase 1

*Note: Some of this functionality I removed when I started building Phases 2 & 3 (like writing to the JSON file). All of this phase's functionality can be found in the `assignment1` branch*

✅  Implement `crew.proto` and `delivery.proto` and the functions listed in `planet_express_service.proto`

✅  Refactor `ship.proto` to use the `crew` and `delivery` resources.

✅  Implement the rpc functions (defined in `planet_express_service.proto`) in both the client (`./headquarters`) and server (`./ship`), following the `GetShip` example. You do not need to add a database. Experiment with using data contained in the gRPC request.

✅  Output data about the planet express ship, deliveries, and crew members to a json file called `planet_express.json`. Feel free to get creative with your data and protobuf resources.

✅  Add a dockerfile to `./ship` so that we can run the server with docker.

#### Phase 2

❌  Provide a basic helm template for the `./ship` server. This should contain a kubernetes deployment and service. We should be able to port-forward your service to call the `./ship` rpc functions.

✅  Add additional protobuf resources like `ShipEngine`. You can get creative with this.
        - Added `CrewMember` to the `Crew` resource

✅  Add GraphQL support to the gateway (headquarters client). A good library for this is https://github.com/graph-gophers/graphql-go

✅  Map a graphql resource and endpoint to a gRPC call such that you can use a GraphQL client to get data from the ship backend.

#### Phase 3

✅  Add a simple React SPA that displays data about the ship using GraphQL to communicate with the gateway (headquarters service). You can put this in the `./dashboard` directory.

❌  Add some go unit tests to the ship server.
