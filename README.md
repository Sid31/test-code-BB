# test-code-BB

This repository contains Go-based solutions for two problems: a concurrency pattern demonstrating the fan-in, fan-out concept and a thread-safe data access demonstration.

**Usage**:
```bash
cd fanin-fanout-problem
go run fanin-fanout.go <URL1> <URL2> ...
go run fanin-fanout.go https://edition.cnn.com/2023/02/20/world/australian-handfish-photograph-c2e-spc-intl-scn/index.html https://www.google.fr

go mod init fanin-fanout
go test

cd thread-safe-problem
go run ts-data-access.go // default 100
go run ts-data-access.go <INTEGER>  

