# goattest

Outputs an ASCII-drawn goat whenever output of `go test` contains failed tests.

```
$ go test | goattest
goat_test.go:23 assertion failed (Validation failed)
             / /
          (\/_//')
           /   '/
          0  0   \
         /        \
        /    __/   \
       /,  _/ \     \_
       '-./ )  |     ~^~^~^~^~^~^~^~\~.
           (   /                     \_}
              |               /      |
              ;     |         \      /
               \/ ,/           \    |
               / /~~|~|~~~~~~|~|\   |
              / /   | |      | | '\ \
             / /    | |      | |   \ \
            / (     | |      | |    \ \
           /,_)    /__)     /__)   /,_/
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
--- FAIL: TestGoatValidity (0.00s)
FAIL
exit status 1
FAIL    github.com/goats-corp/invasion-app     0.042s
```
