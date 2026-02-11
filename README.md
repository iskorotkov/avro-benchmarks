![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/iskorotkov/avro/v2](https://github.com/iskorotkov/avro)
* [github.com/hamba/avro/v2](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)
* [github.com/actgardner/gogen-avro/v9](https://github.com/actgardner/gogen-avro)

#### Results

```
goos: darwin
goarch: arm64
pkg: github.com/iskorotkov/avro-benchmarks
cpu: Apple M4 Pro
BenchmarkGoAvroDecode-12        	 1262725	       952.7 ns/op	     418 B/op      27 allocs/op
BenchmarkGoAvroEncode-12        	  921210	      1285 ns/op	     883 B/op      63 allocs/op
BenchmarkGoGenAvroDecode-12     	 2628411	       464.0 ns/op	     320 B/op      11 allocs/op
BenchmarkGoGenAvroEncode-12     	 4695589	       233.8 ns/op	     240 B/op       3 allocs/op
BenchmarkHambaDecode-12         	 8214915	       144.9 ns/op	      47 B/op       0 allocs/op
BenchmarkHambaEncode-12         	 8826021	       136.4 ns/op	     112 B/op       1 allocs/op
BenchmarkIskorotkovDecode-12    	 8174200	       145.7 ns/op	      47 B/op       0 allocs/op
BenchmarkIskorotkovEncode-12    	 8739876	       138.0 ns/op	     112 B/op       1 allocs/op
BenchmarkLinkedinDecode-12      	 1925295	       622.8 ns/op	    1688 B/op      35 allocs/op
BenchmarkLinkedinEncode-12      	 4478540	       270.0 ns/op	     248 B/op       5 allocs/op
```
