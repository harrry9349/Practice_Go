
go run main.go
# command-line-arguments
.\main.go:12:14: undefined: Appname

go run main.go app.go
Zoo  Application
Glass
Banana
Carrot
Meat
Fish


go test ./animals
--- FAIL: TestElephantFeed (0.00s)
    animals_test.go:12: Grass != Glass
FAIL
FAIL    _/home/... 0.086s

go test -v ./animals
=== RUN   TestElephantFeed
--- PASS: TestElephantFeed (0.00s)
=== RUN   TestMonkeyFeed
--- PASS: TestMonkeyFeed (0.00s)
=== RUN   TestRabbitFeed
--- PASS: TestRabbitFeed (0.00s)
=== RUN   TestLionFeed
--- PASS: TestLionFeed (0.00s)
=== RUN   TestDolphinFeed
--- PASS: TestDolphinFeed (0.00s)
PASS
ok      _/home/... 0.089s

go test -v
=== RUN   TestAppname
--- PASS: TestAppname (0.00s)