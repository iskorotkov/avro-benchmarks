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
BenchmarkGoAvroDecode-12                 1248079               960.1 ns/op           418 B/op         27 allocs/op
BenchmarkGoAvroEncode-12                  941415              1310 ns/op             877 B/op         63 allocs/op
BenchmarkGoGenAvroDecode-12              2594593               466.0 ns/op           320 B/op         11 allocs/op
BenchmarkGoGenAvroEncode-12              5187818               229.2 ns/op           240 B/op          3 allocs/op
BenchmarkHambaDecode-12                  8118740               147.3 ns/op            47 B/op          0 allocs/op
BenchmarkHambaEncode-12                  8540484               139.7 ns/op           112 B/op          1 allocs/op
BenchmarkHeetchDecode-12                   96644             12441 ns/op           37773 B/op        385 allocs/op
BenchmarkHeetchEncode-12                 5392798               221.5 ns/op           384 B/op          5 allocs/op
BenchmarkIskorotkovDecode-12            10105106               118.6 ns/op            47 B/op          0 allocs/op
BenchmarkIskorotkovEncode-12             8573349               139.3 ns/op           112 B/op          1 allocs/op
BenchmarkLinkedinDecode-12               1882018               637.8 ns/op          1688 B/op         35 allocs/op
BenchmarkLinkedinEncode-12               4127374               288.0 ns/op           248 B/op          5 allocs/op
```
