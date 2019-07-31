# Sets

This is a simple command line program wrapping the wonderful `github.com/deckarep/golang-set` library. It allows you to easily perform difference, intersection, and union operations on two sets.

# Examples:

```shell
$ sets -d -i -u -a 1,2,3 -b 3,4,5
\: Set{1, 2}
∩: Set{3}
∪: Set{1, 2, 3, 5, 4}
```