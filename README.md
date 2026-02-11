![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

#### Libraries

* [github.com/iskorotkov/avro/v2](https://github.com/iskorotkov/avro)
* [github.com/hamba/avro/v2](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)
* [github.com/actgardner/gogen-avro/v9](https://github.com/actgardner/gogen-avro)
* [github.com/heetch/avro](https://github.com/heetch/avro)

#### Results

```
goos: darwin
goarch: arm64
pkg: github.com/iskorotkov/avro-benchmarks
cpu: Apple M4 Pro
BenchmarkGoAvroDecode-12        	 1220336	       983.2 ns/op	     418 B/op      27 allocs/op
BenchmarkGoAvroEncode-12        	  966108	      1298 ns/op	     869 B/op      63 allocs/op
BenchmarkGoGenAvroDecode-12     	 2619901	       459.0 ns/op	     320 B/op      11 allocs/op
BenchmarkGoGenAvroEncode-12     	 4637014	       236.2 ns/op	     240 B/op       3 allocs/op
BenchmarkHambaDecode-12         	 8073816	       147.2 ns/op	      47 B/op       0 allocs/op
BenchmarkHambaEncode-12         	 8566797	       141.2 ns/op	     112 B/op       1 allocs/op
BenchmarkHeetchDecode-12        	   95692	     12716 ns/op	   37773 B/op     385 allocs/op
BenchmarkHeetchEncode-12        	 5241297	       222.0 ns/op	     384 B/op       5 allocs/op
BenchmarkIskorotkovDecode-12    	 8109548	       147.4 ns/op	      47 B/op       0 allocs/op
BenchmarkIskorotkovEncode-12    	 8666217	       138.9 ns/op	     112 B/op       1 allocs/op
BenchmarkLinkedinDecode-12      	 1864414	       646.4 ns/op	    1688 B/op      35 allocs/op
BenchmarkLinkedinEncode-12      	 4233098	       279.7 ns/op	     248 B/op       5 allocs/op
```
