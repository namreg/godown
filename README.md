## Godown

[![Build Status](https://www.travis-ci.org/namreg/godown.svg?branch=master)](https://www.travis-ci.org/namreg/godown)
[![Go Report Card](https://goreportcard.com/badge/github.com/namreg/godown)](https://goreportcard.com/report/github.com/namreg/godown)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/namreg/godown/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/namreg/godown/branch/master/graph/badge.svg)](https://codecov.io/gh/namreg/godown)

A simple, distributed, fault-tolerant key-value storage inspired by Redis. It uses [Raft](https://raft.github.io) protocotol as consensus algorithm.
It supports the following data structures: `String`, `Bitmap`, `Map`, `List`.

[![asciicast](https://asciinema.org/a/lNp3lOJlnnp9WQW3kKnguL35e.png)](https://asciinema.org/a/lNp3lOJlnnp9WQW3kKnguL35e)

### How to install

#### Install via binaries
You can find binaries on the [Github releases page](https://github.com/namreg/godown/releases).

#### Install via docker
```bash
# creating a network
docker network create godown 

# creating a volume
docker volume create godown 

# bootstrap a cluster with a single node
docker run -it --rm -v godown:/var/lib/godown --name=godown_1 --net=godown -p 5000:5000 \
    namreg/godown-server -id 1 -listen godown_1:5000 -raft godown_1:6000

# join the second node to the cluster
docker run -it --rm -v godown:/var/lib/godown --name=godown_2 --net=godown -p 5001:5001 \
    namreg/godown-server -id 2 -listen godown_2:5001 -join godown_1:5000 -raft godown_2:6001

# join the third node to the cluster
docker run -it --rm -v godown:/var/lib/godown --name=godown_3 --net=godown -p 5002:5002  \
    namreg/godown-server -id 3 -listen godown_3:5001 -join godown_1:5000 -raft godown_3:6002
```

Available options to run a server:
```bash
  -dir string
        Directory where data is stored.
  -gc duration
        Garbage collector interval.
  -id string
        Server unique id.
  -join string
        Server address to join.
  -listen string
        Server address to listen.
  -raft string
        Raft protocol listen address.
  -resp string
        Redis Serialization Protocol listen address.      
  -version
        Show version.
```

### How to connect

You can connect to any godown node. All modifications will be replicated to all nodes.

#### Connect via any redis client
If you have specified `resp` address while starting a node, you can connect to the one by any redis client.
```go
package main

import (
	"fmt"
	"net"
	"time"

	"github.com/garyburd/redigo/redis"
)

const connectionTimeout = 100 * time.Millisecond

func main() {
	conn, err := net.Dial("tcp", "6380")
	if err != nil {
		panic(err)
	}
	rconn := redis.NewConn(conn, connectionTimeout, connectionTimeout)

	reply, err := rconn.Do("LRANGE", "list", 0, 100)
	vals, err := redis.Strings(reply, err)

	if err != nil {
		panic(err)
	}

	fmt.Println(vals)
}

```

#### Connect via CLI
```bash
godown-cli
```

Available options:
```bash
  -host string
    	Host to connect to a server (default "127.0.0.1")
  -port string
    	Port to connect to a server (default "4000")
  -version
    	Show godown version.
```

Supported commands:

| Command| Description |
|---|---|
| HELP&nbsp;command | Show the usage of the given command. |
| TYPE&nbsp;key | Returns the type stored at key. |
| KEYS&nbsp;pattern | Find all keys matching the given pattern. |
| PING&nbsp;[message] | Returns PONG if no argument is provided, otherwise return a copy of the argument as a bulk. |
| EXPIRE&nbsp;key&nbsp;seconds | Set a timeout on key. After the timeout has expired, the key will automatically be deleted. |
| TTL&nbsp;key | Returns the remaining time to live of a key. -1 returns if key does not have timeout. |
|---|---|---|
| SET&nbsp;key&nbsp;value | Set key to hold the string value. If key already holds a value, it is overwritten. |
| GET&nbsp;key | Get the value by key. If provided key does not exist NIL will be returned. |
| STRLEN&nbsp;key | Returns length of the given key. If key does not exists, 0 will be returned. |
| DEL&nbsp;key | Delete the given key. |
|---|---|---|
| SETBIT&nbsp;key&nbsp;offset&nbsp;value | Sets or clears the bit at offset in the string value stored at key. |
| GETBIT&nbsp;key&nbsp;offset | Returns the bit value at offset in the string value stored at key. |
|---|---|---|
| LPUSH&nbsp;key&nbsp;value&nbsp;[value&nbsp;...] | Prepend one or multiple values to a list. |
| LPOP&nbsp;key | Removes and returns the first element of the list stored at key. |
| RPUSH&nbsp;key&nbsp;value&nbsp;[value&nbsp;...] | Append one or multiple values to a list. |
| RPOP&nbsp;key | Removes and returns the last element of the list stored at key. |
| LLEN&nbsp;key | Returns the length of the list stored at key. If key does not exist, it is interpreted as an empty list and 0 is returned. |
| LINDEX&nbsp;key&nbsp;index | Returns the element at index index in the list stored at key. <br>The index is zero-based, so 0 means the first element, 1 the second element and so on. Negative indices can be used to designate elements starting at the tail of the list. |
| LRANGE&nbsp;key&nbsp;start&nbsp;stop | Returns the specified elements of the list stored at key.<br> The offsets start and stop are zero-based indexes, with 0 being the first element of the list (the head of the list), 1 being the next element and so on. |
| LREM&nbsp;key&nbsp;value | Removes all occurrences of elements equal to value from the list stored at key. |
|---|---|---|
| HSET&nbsp;key&nbsp;field&nbsp;value | Sets field in the hash stored at key to value. |
| HGET&nbsp;key&nbsp;field | Returns the value associated with field in the hash stored at key. |
| HKEYS&nbsp;key | Returns all field names in the hash stored at key. Order of fields is not guaranteed. |
| HVALS&nbsp;key | Returns all values in the hash stored at key. |
| HDEL&nbsp;key&nbsp;field&nbsp;[field&nbsp;...] | Removes the specified fields from the hash stored at key. Returns the number of fields that were removed. |

#### Connect via go client
```go
package main

import (
	"fmt"

	"github.com/namreg/godown/client"
)

func main() {
	c, err := client.New("127.0.0.1:4000")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	res := c.Get("key")
	if res.Err() != nil {
		panic(res.Err())
	}

	if res.IsNil() {
		fmt.Print("key does not exist")
	} else {
		fmt.Println(res.Int64())
	}
}
```
Client documentation available at [godoc](https://godoc.org/github.com/namreg/godown/client)

### TODO
- [ ] Write more docs
- [ ] Write more tests
