### Queue

A [`queue`](https://en.wikipedia.org/wiki/Queue_(abstract_data_type)) is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence.

A simple way to implement a temporary queue data structure in Go is to use a slice:

to enqueue you use the built-in append function, and
to dequeue you slice off the first element.

Please implement your own `queue`.

### Run tests

```
go test .
```

