
A CLI tool to bechmark Hazelcast IMDG clusters.

### Installation

```
> brew tap mdogan/hazelcast
> brew install hazelcast-benchmark
```

### Usage

```
Usage of hazelcast-benchmark:
  -c int
    	Number of client threads (default 10)
  -cluster string
    	Hazelcast cluster name (default "dev")
  -d int
    	Data size in bytes (default 128)
  -h string
    	Server hostname (default "127.0.0.1")
  -m string
    	Name of the Hazelcast IMap (default "benchmark")
  -n int
    	Number of total requests (default 1000)
  -p int
    	Server port (default 5701)
  -r int
    	Key space range (default 1000)
  -ratio string
    	Set:Get ratio (default "1:10")
```
