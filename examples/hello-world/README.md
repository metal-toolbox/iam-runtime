# hello-world - Example IAM runtime and service

This example provides an IAM runtime that only authenticates a single subject with a static credential token `hello` and only authorizes that subject to perform the action `greet` to the resource `world`.

## Running

To run the example, build and run the runtime first:

```
$ go build ./cmd/runtime
$ ./runtime
```

Then, in another terminal window in this directory, build and run the service:

```
$ go build ./cmd/service
$ ./service
```

Then, try the following curl commands:

```
$ curl --oauth2-bearer "hello" http://localhost:8080/whoami
$ curl --oauth2-bearer "goodbye" http://localhost:8080/whoami
$ curl --oauth2-bearer "hello" 'http://localhost:8080/can-i?what=greet&who=world'
$ curl --oauth2-bearer "goodbye" 'http://localhost:8080/can-i?what=greet&who=world'
$ curl --oauth2-bearer "hello" 'http://localhost:8080/can-i?what=greet&who=universe'
```
