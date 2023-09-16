# Unsafe string↔[]byte conversion library

The library functions for unsafely convert between a string and
a slice of bytes.  You probably shouldn’t use it unless you need to
squeeze extra performance from your performance-critical code path.

See https://mina86.com/2017/golang-string-and-bytes/ for some more info.

## Bench
20230916
```
goarch: amd64
pkg: github.com/3JoB/unsafeConvert
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
Benchmark_StringToByte_G_Lite-20            	506960354	         2.390 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_Copy_U_Lite-20       	100000000	        10.06 ns/op	      16 B/op	       1 allocs/op
Benchmark_StringToByte_Slice_U_Lite-20      	1000000000	         0.2269 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_Bytes_U_Lite-20      	52430136	        23.88 ns/op	      64 B/op	       1 allocs/op
Benchmark_StringToByte_Pointer_U_Lite-20    	1000000000	         0.1159 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_G_Big-20             	 1000000	      1086 ns/op	   10880 B/op	       1 allocs/op
Benchmark_StringToByte_Copy_U_Big-20        	 1117935	      1073 ns/op	   10880 B/op	       1 allocs/op
Benchmark_StringToByte_Slice_U_Big-20       	1000000000	         0.4338 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToByte_Bytes_U_Big-20       	 1000000	      1138 ns/op	   10880 B/op	       1 allocs/op
Benchmark_StringToByte_U_Pointer_Big-20     	1000000000	         0.1117 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_G_Lite-20            	613644068	         1.979 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_Slice_U_Lite-20      	1000000000	         0.2174 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_Strings_U_Lite-20    	63382401	        17.89 ns/op	      16 B/op	       1 allocs/op
Benchmark_ByteToString_Pointer_U_Lite-20    	1000000000	         0.1119 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_G_Big-20             	 1000000	      1139 ns/op	   10880 B/op	       1 allocs/op
Benchmark_ByteToString_Slice_U_Big-20       	1000000000	         0.4358 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_U_Pointer_Big-20     	1000000000	         0.1100 ns/op	       0 B/op	       0 allocs/op
Benchmark_ByteToString_U_Strings_Big-20     	 1000000	      1126 ns/op	   10880 B/op	       1 allocs/op
Benchmark_Float32_U-20                      	18850288	        58.05 ns/op	      40 B/op	       2 allocs/op
Benchmark_Float32_FMT-20                    	11351810	       114.2 ns/op	      16 B/op	       2 allocs/op
Benchmark_Itoa_U-20                         	186808952	         6.235 ns/op	       0 B/op	       0 allocs/op
Benchmark_Itoa_G-20                         	88849399	        13.01 ns/op	       5 B/op	       1 allocs/op
Benchmark_Atoi_U-20                         	135130189	         8.779 ns/op	       0 B/op	       0 allocs/op
Benchmark_Atoi_G-20                         	342945426	         3.713 ns/op	       0 B/op	       0 allocs/op
```