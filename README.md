# Factory Method Pattern

The Factory method pattern (or simply, Factory) is probably the second-best known and used design pattern in the industry. Its purpose is to abstract the user from the knowledge of the struct he needs to achieve for a specific purpose, such as retrieving some value, maybe from a web service or a database. The user only needs an interface that provides him this value. By delegating this decision to a Factory, this Factory can provide an interface that fits the user needs. It also eases the process of downgrading or upgrading of the implementation of the underlying type if needed.

## The Example

For our example, we are going to implement a device inventory with the Factory method, which will provide us with different ways to categorize the devices in a store inventory. At first, we will have two methods to categorize the devices. We will also have an interface with the methods, Brand, and Stock, with their respective getters and setters that each structure that wants to use a device must implement.

## Running the client

```sh
go run main.com
```

## Unit Tests

The next commands should be execute inside the folder named devices.
```sh
cd devices/
```


* Execute a single unit test per function name
```sh
go test -v -run=TestGetDeviceFactorySmartWatch
```


* Execute all unit tests
```sh
go test -v
```


* Generate code coverage by unit tests metrics and show these in HTML report

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out        
```
* Code coverage with unit tests

![alt coverage](/images/code_coverage.png)