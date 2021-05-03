Go like Laravel


+ Go-Gin framework
+ RouteAPI-Middleware-MVC-Service like Laravel
-- Route: ok, multi-route-grouping ok
-- Middleware*: Auth (basic ok, JWT ?), 
-- MVC-Service: controller ok, model ok, view?X, service ok
----- data json binding OK, validation data OK
+ Logging ok, Testing ok, Debugging not
+ Jobs Queue* (sync, database, redis*, rabbitMQ): Notifications(Mail), Events
+ Config ok
+ RPC OK, gRPC ?


Installation


Testing & benchmark & document(Output: XX)
- go test -v
- go test -coverprofile=coverage.out
- go tool cover -html=coverage.out
- go test -bench=.
- go test -bench=Add

Run Server
- go run *.go