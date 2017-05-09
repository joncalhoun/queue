# queue

A generator for creating typed queues + the queue implementations themselves.

## Example usage

All of the sample files here were created by navigating to thie `queue` directory and then running the following:

```
go run gen/main.go -name=String -type=string > string.go
go run gen/main.go -name=Int -type=int > int.go
go run gen/main.go -name=IntSlice -type="[]int" > int_slice.go
# Getting a little meta
go run gen/main.go -name=List -type="*list.List" > container_list.go
```
