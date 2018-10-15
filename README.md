## godown

[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/namreg/godown/blob/master/LICENSE)

A simple, distributed, fault-tolerant key-value storage inspired by Redis. It uses [Raft](https://raft.github.io) protocotol as consensus algorithm.
It supports the following data structures: `String`, `Bitmap`, `Map`, `List`.

[![asciicast](https://asciinema.org/a/lNp3lOJlnnp9WQW3kKnguL35e.png)](https://asciinema.org/a/lNp3lOJlnnp9WQW3kKnguL35e)

### How to install

#### Install via binaries
You can find binaries on the [Github releases page](https://github.com/namreg/godown/releases).

#### Install via docker
```bash
# creating an network
docker network create godown 

# creating an volume
docker volume create godown 

# bootsrap a cluster with a single node
docker run -it --rm -v godown:/var/lib/godown --name=godown_1 --net=godown -p 5000:5000 \
    namreg/godown-server -id 1 -listen godown_1:5000 -raft godown_1:6000

# join the second node to the cluster
docker run -it --rm -v godown:/var/lib/godown --name=godown_2 --net=godown -p 5001:5001 \ 
    namreg/godown-server -id 2 -listen godown_2:5001 -join godown_1:5000 -raft godown_2:6001

# join the third node to the cluster
docker run -it --rm -v godown:/var/lib/godown --name=godown_2 --net=godown -p 5001:5001  \
    namreg/godown-server -id 2 -listen godown_2:5001 -join godown_1:5000 -raft godown_2:6001
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
  -version
        Show version.
```

### How to connect

You can connect to any godown node. All modifications will be replicated to all nodes.

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
| Command | Usage | Description |
|---|---|---|
| HELP | HELP command | Show the usage of the given command. |
| TYPE | TYPE key | Returns the type stored at key. |
| KEYS | KEYS pattern | Find all keys matching the given pattern. |
| PING | PING [message] | Returns PONG if no argument is provided, otherwise return a copy of the argument as a bulk. |
| EXPIRE | EXPIRE key seconds | Set a timeout on key. After the timeout has expired, the key will automatically be deleted. |
| TTL | TTL key | Returns the remaining time to live of a key. -1 returns if key does not have timeout. |
|---|---|---|
| SET | SET key value | Set key to hold the string value. If key already holds a value, it is overwritten. |
| GET | GET key | Get the value by key. If provided key does not exist NIL will be returned. |
| STRLEN | STRLEN key | Returns length of the given key. If key does not exists, 0 will be returned. |
| DEL | DEL key | Delete the given key. |
|---|---|---|
| SETBIT | SETBIT key offset value | Sets or clears the bit at offset in the string value stored at key. |
| GETBIT | GETBIT key offset | Returns the bit value at offset in the string value stored at key. |
|---|---|---|
| LPUSH | LPUSH key value [value ...] | Prepend one or multiple values to a list. |
| LPOP | LPOP key | Removes and returns the first element of the list stored at key. |
| LLEN | LLEN key | Returns the length of the list stored at key. If key does not exist, it is interpreted as an empty list and 0 is returned. |
| LINDEX | LINDEX key index | Returns the element at index index in the list stored at key. <br>The index is zero-based, so 0 means the first element, 1 the second element and so on. Negative indices can be used to designate elements starting at the tail of the list. |
| LRANGE | LRANGE key start stop | Returns the specified elements of the list stored at key.<br> The offsets start and stop are zero-based indexes, with 0 being the first element of the list (the head of the list), 1 being the next element and so on. |
| LREM | LREM key value | Removes all occurrences of elements equal to value from the list stored at key. |
|---|---|---|
| HSET | HSET key field value | Sets field in the hash stored at key to value. |
| HGET | HGET key field | Returns the value associated with field in the hash stored at key. |
| HKEYS | HKEYS key | Returns all field names in the hash stored at key. Order of fields is not guaranteed. |
| HVALS | HVALS key | Returns all values in the hash stored at key. |

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
Client documentation available at [godoc](https://godoc.org/github.com/namreg/godown)