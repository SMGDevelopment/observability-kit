# observability-kit
The idea with this is to create a small library or kit to help with general observability.

## Logger
Implements a basic slog logger with adjustable log levels via the initializer. 
Provides functions for logging with context at different log levels.

## Metrics
Provides wrapper functions for interacting with prometheus via a RED strategy based connection.

## Errors
Provides helpful functions for wrapping and unwrapping errors. 
As well as provides a wrapper function for recording errors via the logger and metrics utility.