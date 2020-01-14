# abstractlogger

Abstractlogger is a logging frontend to abstract away logging from your frontend of choice.
It's a very thin interface that makes a tradeoff against too much flexibility in favour of no performance overhead.

Abstractlogger enables you to use your logging backend of choice.
Currently there's a zap and logrus implementation.
Feel free to add additional logging backend implementations.

You should consider using abstract logger in two situations:
1. You're building a library and don't want to make the choice for your user on which logging backend to use.
2. You're building an application and want logging in the hot path. You're unsure which logging library to use. In this case Abstractlogger helps you to keep your domain logic separated from a 3rd party logging library as you can change it every time without updating all of your code.

If you look at the logging interface itself you'll find that it's not as dynamic as most logging libraries are.
This is because Abstractlogger doesn't want to add any additional runtime overhead.
So the decision was made to not allow variadic field arguments.

If you feel an important method is missing please file an issue or create a PR.
Tested code is welcomed.

## Benchmarks

```text
BenchmarkNoopLogger/log_level_invalid/variadic/noop-16    	             30566610	      42.5 ns/op	     224 B/op	       1 allocs/op
BenchmarkNoopLogger/log_level_invalid/variadic/logrus-16  	              4010529          300 ns/op	    1393 B/op	      10 allocs/op
BenchmarkNoopLogger/log_level_invalid/variadic/zap-16     	             30208502	      41.0 ns/op	     192 B/op	       1 allocs/op
BenchmarkNoopLogger/log_level_invalid/variadic/abstract_zap-16         	364706740	      3.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoopLogger/log_level_invalid/variadic/abstract_logrus-16      	343806495	      3.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoopLogger/log_level_invalid/non_variadic/noop-16             	390515324         2.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoopLogger/log_level_invalid/non_variadic/abstract_zap-16     	407610915	      3.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoopLogger/log_level_invalid/non_variadic/abstract_logrus-16  	377882335	      3.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoopLogger/log_level_valid/variadic/noop-16                   	 22977513	      51.4 ns/op	     224 B/op	       1 allocs/op
BenchmarkNoopLogger/log_level_valid/variadic/logrus-16                 	   172395	      6411 ns/op	    2741 B/op	      38 allocs/op
BenchmarkNoopLogger/log_level_valid/variadic/zap-16                    	  9942966	       135 ns/op	     192 B/op	       1 allocs/op
BenchmarkNoopLogger/log_level_valid/variadic/abstract_zap-16           	  8113160	       150 ns/op	     192 B/op	       1 allocs/op
BenchmarkNoopLogger/log_level_valid/variadic/abstract_logrus-16        	   215976	      4982 ns/op	    2019 B/op	      32 allocs/op
BenchmarkNoopLogger/log_level_valid/non_variadic/noop-16               	426628474	      2.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoopLogger/log_level_valid/non_variadic/abstract_zap-16       	  8872508	       144 ns/op	     192 B/op	       1 allocs/op
BenchmarkNoopLogger/log_level_valid/non_variadic/abstract_logrus-16    	   154567	      7034 ns/op	    2786 B/op	      41 allocs/op
```