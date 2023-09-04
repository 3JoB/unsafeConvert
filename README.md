# Unsafe string↔[]byte conversion library

The library functions for unsafely convert between a string and
a slice of bytes.  You probably shouldn’t use it unless you need to
squeeze extra performance from your performance-critical code path.

See https://mina86.com/2017/golang-string-and-bytes/ for some more info.

## Bench
20230904
```
goos: windows
goarch: amd64
pkg: github.com/3JoB/unsafeConvert
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
Benchmark_StringToByte_G_Lite-20            	507678423	         2.442 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_Copy_U_Lite-20       	100000000	        11.43 ns/op	      16 B/op	       1 allocs/op
Benchmark_StringToByte_Slice_U_Lite-20      	1000000000	         0.2499 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_Bytes_U_Lite-20      	45056714	        30.45 ns/op	      64 B/op	       1 allocs/op
Benchmark_StringToByte_Pointer_U_Lite-20    	1000000000	         0.1175 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_G_Big-20             	  848536	      1418 ns/op	   10880 B/op	       1 allocs/op
Benchmark_StringToByte_Copy_U_Big-20        	 1000000	      1247 ns/op	   10880 B/op	       1 allocs/op
Benchmark_StringToByte_Slice_U_Big-20       	1000000000	         0.4381 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_Bytes_U_Big-20       	  900399	      1322 ns/op	   10880 B/op	       1 allocs/op
Benchmark_StringToByte_U_Pointer_Big-20     	1000000000	         0.1121 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_G_Lite-20            	575563584	         2.093 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_Slice_U_Lite-20      	1000000000	         0.2444 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_Strings_U_Lite-20    	59169846	        20.45 ns/op	      16 B/op	       1 allocs/op
Benchmark_ByteToString_Pointer_U_Lite-20    	1000000000	         0.1185 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_G_Big-20             	  924814	      1107 ns/op	   10880 B/op	       1 allocs/op
Benchmark_ByteToString_Slice_U_Big-20       	1000000000	         0.4491 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_U_Pointer_Big-20     	1000000000	         0.1167 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_U_Strings_Big-20     	 1021662	      1293 ns/op	   10880 B/op	       1 allocs/op
PASS
ok  	github.com/3JoB/unsafeConvert	18.308s
```