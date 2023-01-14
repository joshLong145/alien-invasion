# Alien Invasion Simulator 👾

Small map traversal program where aliens fight in locations until the map is destroyed or a move threshold is met.

## Example input file
```
Foo​ ​north=Bar​ ​west=Baz​ ​south=Qu-ux
Bar​ ​south=Foo​ ​west=Bee
```

## Running
**go 1.18 or newer is required**

```bash
go run cmd/main.go <full-path to input file> <number-of-aliens>
```

### Running tests
```bash
go test cmd/world/world_test.go
```