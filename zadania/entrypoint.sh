go test -v exc00.go exc00_test.go
go test -v exc01.go exc01_test.go
go test -v exc02.go exc02_test.go
go test -bench=. -run=^PrimesCounterHandler
