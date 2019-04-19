# liveness

### Install

Requires [Go](https://golang.org/)

```
$ go get github.com/kroppt/liveness
```

### Usage

```
$ liveness < blocks.txt
```

### Format

```
<name> <defines> <uses0>[,<uses1>,...,<usesN>]
<required empty line>
<from> <to>
<required empty line>
```

Spaces are rquired between `name`, `defines` ande `uses list`.

### Example

Input:
```
B0 a 
B1 b 
B2 c a,b
B3  c

B0 B1
B1 B2
B2 B3

```

Output:
```
B0
  LVin: { }
  LVout: { a }
B1
  LVin: { a }
  LVout: { a b }
B2
  LVin: { a b }
  LVout: { c }
B3
  LVin: { c }
  LVout: { }
```
