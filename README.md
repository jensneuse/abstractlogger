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
