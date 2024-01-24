# Learning golang sintax and features
This repository is made to document the learning process I go trhough with golang. 

## Testing
Use `go test -v` to get the result of all tests

## Differnces between Maps and Structs
Map
- All keys must be of the same type
- All values must be of the same type
- Since keys support indexing, then can be iterated through
- Usable for representing a collection of related properties
- No need for knowing all keys at compile time
- Is a Reference Type

Struct
- Values can be of different types
- Keys don't support indexing
- All different fields must be known at compile time
- Usable to represent "thing" with a lot of different properties
- Is a Value Type