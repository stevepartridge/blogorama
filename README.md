# blogorama

```

the server side of a blogging application consisting of an HTTP API and a database.
The API should be able to:

1. Create, update, delete users.
2. Allow users to manage their posts.

```

## Development

The local development setup _should_ be a fairly pain-free experience. Using the following instructions should help get you up and running in short order.

Note: This has only been tested in a \*nix environment so your milage may vary if you're using something different such as Windows or a typewriter.

#### Prerequisites

##### [Go 1.13+](https://golang.org/dl/)

_The has only been tested with 1.13 and 1.14_

##### Docker 18.06+ and Compose 1.25+

[Docker Desktop](https://www.docker.com/products/docker-desktop) makes setup and usage fairly easy for Docker and Docker Compose, as well as Kubernetes.

##### Bash 4+

Due to some limitations with Bash 3 it's recommended to have Bash 4+.

### Running the Service

Everything is managed via a [Makefile](https://github.com/stevepartridge/blogorama/tree/master/_scripts) and a collection of [helper scripts](https://github.com/stevepartridge/blogorama/tree/master/Makefile) in the effort to ease the setup and development processes.

The service is intended to run in a semi-real way that simulates a deployed service. To this it is recommended to add `host.local` to your `/etc/hosts` file. You can use the helper command to this easily:

```
make hosts
```

**TL;DR:**

After creating a `local.env` from copying `template.env`, this _should_ be the only command needed to get things setup and started

```
make run
```

If all is successful you can test that it's up and working by going to `https://host.local:8000/info` (it will require allowing the [self-signed SSL certificate](https://github.com/stevepartridge/blogorama#create-self-signed-local-ssl-certificates))

#### What's happening behind the curtain with the `make run` command?

###### Check prerequisites

`make check-prereqs`

Ensures you're running the minimum versions for Go, Docker, and Compose. It will also check for `local.env` and stop if it is not found.

###### Create self-signed local ssl certificates

`make certs`

In the effort to use TLS/SSL locally as it would be in the wild, these are used to simulate a more realistic protocol. However, it still requires skiping TLS verification (and/or accepting an unverified Certificate when using a browser)

###### Start/restart Docker containers

`make docker` and `make stop-docker`

This sets up a bridged network and necessary dependency containers. The stop command shuts down the containers but does not remove the network.

###### Dependencies (go mod)

`make deps`

Ensure all the necessary go mod dependencies are present.

###### Protobufs

`make protos`

This service leverages protobufs to define the payload contracts/data for [gRPC](https://grpc.io/) and HTTP/1 API provided automatically with [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway). This step attempts to install [protoc](https://github.com/protocolbuffers/protobuf) which can be tricky based on various development environment quirks. The [script](https://github.com/stevepartridge/blogorama/blob/master/_scripts/protos.sh) that handles it tries to determine the best way to install it and then attempts to do so. Once protoc and the grpc-gateway packages are in place it will auto generate the Go files for use within the service. They can also be referenced from external services wishing to communicate with it requiring little effort. The JavaScript files for doing so are included as an example of simplicity for supporting other languages.

###### Bindata/Assets

`make bindata`

The command packages up any necessary asset files to be included in the binary in the effort to make the service a zero-dependency binary.

###### Swagger Assets

`make swagger`

As part of using the protobuf files swagger assets to present the auto-generated documentation resources with no additional work. You can see that by hitting the `https://host.local:8000/docs` endpoint once the service is up and running.

### Testing
To run the tests you can use
```
make tests
```
OR
```
make test-show-coverage
```
This will open up the coverage reports in the browser after successfully running the tests.

## Using the API

#### Authentication

Instead of adding the complexity of leveraging OAuth2 or the like to handle authentication there's a [_Faux Auth_](https://github.com/stevepartridge/blogorama/tree/master/pkg/fauxauth) mechanism to fake the authorization and access. With the exception `/info` and `POST:/v1/users` all endpoints are protected and require an API Key in the `X-API-Key` header.

Generating an API Key is handled by creating a user.

```
curl -X POST "https://host.local:8000/v1/users" -H "accept: application/json" -H "Content-Type: application/json" -d "{  \"name\": \"Creed Bratton\"}"
```

Which should result in a response similar to this:

```json
{
  "user": {
    "id": 1,
    "api_key": "c2fcf476-9d46-11ea-9862-dca904888369",
    "name": "Creed Bratton",
    "active": true,
    "created_at": "2020-05-23T22:43:03.336523Z",
    "updated_at": "2020-05-23T22:43:03.336523Z"
  }
}
```

The API Key is only shown once when creating a user.

#### Users

Managing users is fairly basic but it does support the following:

- Create
- Update
- Delete
- Get
- Get List of Users (with pagination)

#### Posts

The posts that users can create are handles like so:

- Create
- Update - only creator
- Delete - only creator
- Get - only creator if flagged as private
- Get All (with pagination)
- Get All for Specific User (with pagination)

#### API Documenation

Once the service is running you can view the full API OpenAPI/Swagger documentation at `https://host.local:8000/docs` or view it in your preferred viewer by referencing the JSON at `https://host.local:8000/docs/swagger.json`. It will also allow for running test calls with an API Key via the `Authorize` button (after you've created a user).
