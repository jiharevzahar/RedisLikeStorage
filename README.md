# Redis like key-value store in Go

This is a basic implementation of a key-value storage system in Go that supports `Set`, `Get`, and `Delete` methods with Time-To-Live (TTL) functionality.

## Features

- Uses a map to achieve O(1) complexity for `Get`, `Set`, and `Delete` operations.
- Supports TTL for automatic data expiration.

## How It Works

- The program defines a `KeyValueStore` struct with a map to store key-value pairs.
- The `Set` method adds key-value pairs with an optional TTL.
- The `Get` method retrieves values, respecting TTL.
- The `Delete` method removes key-value pairs from the store.

## How to Run

   ```bash
   make run
   ```
## Example Output

```plaintext
Get key1: value1
Get key2: value2
Get key1 after expiration: <nil>
Get key2 after deletion: <nil>
```

### Missing parts:

- Tests
- Dockerfile

### What also could be improved

- Service could be implemented as BST and complexity going to be O(log N)
- To achieve even better scalability map can be replaced with Sharded Map algorithm.
