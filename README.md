# GO ORM Benchmarks

### About

ORM benchmarks for GoLang. Originally forked from [orm-benchmark](https://github.com/milkpod/orm-benchmark).
Contributions are wellcome.

### Environment

* go version go1.8 linux/amd64

### PostgreSQL

* PostgreSQL 9.6 for Linux on x86_64

### ORMs

* [dbr](https://github.com/gocraft/dbr)
* [genmai](https://github.com/naoina/genmai)
* [gorm](https://github.com/jinzhu/gorm)
* [gorp](https://github.com/go-gorp/gorp/tree/v2.1)
* [pg](https://github.com/go-pg/pg)
* [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
* [sqlx](https://github.com/jmoiron/sqlx)
* [xorm](https://github.com/xormplus/xorm)
* [godb](https://github.com/samonzeweb/godb)
* [upper.io/db.v3](https://upper.io/db.v3)
* [hood](https://github.com/eaigner/hood)
* [modl](https://github.com/jmoiron/modl)
* [qbs](https://github.com/coocood/qbs)
* [pop](https://github.com/gobuffalo/pop)


### Notes:

#### Hood
- `hood` needs patch for reflecting `string` values:

    `github.com/eaigner/hood/base.go:50` should be patched to:

    `fieldValue.SetString(string(driverValue.Elem().String()))`
- Multi insert is too slow(over 100 seconds), need check/help

### QBS
- `qbs` needs patch for reflecting `string` values:

    `github.com/coocood/qbs/base.go:54` should be patched to:

    `fieldValue.SetString(string(driverValue.Elem().String()))`
- `BulkInsert` is not working as expected.


### Gorm
- [No support for multi insert](https://github.com/jinzhu/gorm/issues/255)

#### Genmai

- Fails on reading 10000 rows (err=>sql: expected 4464 arguments, got 70000)

#### Gorp
- `BulkInsert` is not working as expected. It's too slow.

#### Pop
- `BulkInsert` is not working as expected. It's too slow.

### Prepare DB:

```sql
CREATE ROLE bench LOGIN PASSWORD 'pass'
   VALID UNTIL 'infinity';
CREATE DATABASE benchdb
  WITH OWNER = bench;
```

### Run

```go
go get github.com/derkan/gobenchorm
# install
go install
# build
go build
# all
orm-benchmark -multi=20 -orm=all
# portion
orm-benchmark -multi=20 -orm=xorm -orm=raw
```

### Reports

```
dbr
                   Insert:   2000     4.01s      2005213 ns/op    4297 B/op     81 allocs/op
       BulkInsert 100 row:    500     0.02s        37854 ns/op    2866 B/op     42 allocs/op
                   Update:   2000     0.19s        95029 ns/op    3158 B/op     58 allocs/op
                     Read:   4000     0.37s        93290 ns/op    2261 B/op     39 allocs/op
     MultiRead limit 1000:   2000    16.90s      8449983 ns/op 2071583 B/op  45961 allocs/op
pop
                   Insert:   2000    15.41s      7705992 ns/op   61776 B/op    720 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     6.28s      3140076 ns/op   32931 B/op    428 allocs/op
                     Read:   4000    12.98s      3245945 ns/op   29511 B/op    313 allocs/op
     MultiRead limit 1000:   2000    12.10s      6048901 ns/op  406315 B/op  16231 allocs/op
sqlx
                   Insert:   2000     4.10s      2051602 ns/op    2769 B/op     65 allocs/op
       BulkInsert 100 row:    500     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
                   Update:   2000     4.14s      2069727 ns/op    1424 B/op     35 allocs/op
                     Read:   4000     0.37s        92089 ns/op    1936 B/op     47 allocs/op
     MultiRead limit 1000:   2000     7.69s      3842905 ns/op  330543 B/op  14956 allocs/op
db.v3
                   Insert:   2000     6.54s      3270646 ns/op   37549 B/op   1642 allocs/op
       BulkInsert 100 row:    500     2.41s      4815923 ns/op  486190 B/op  20014 allocs/op
                   Update:   2000     5.82s      2911505 ns/op   42683 B/op   1937 allocs/op
                     Read:   4000     0.95s       236915 ns/op    7105 B/op    296 allocs/op
     MultiRead limit 1000:   2000     8.18s      4091602 ns/op  463532 B/op  15300 allocs/op
pg
                   Insert:   2000     4.02s      2008151 ns/op    1990 B/op     10 allocs/op
       BulkInsert 100 row:    500     1.88s      3766258 ns/op   14507 B/op    212 allocs/op
                   Update:   2000     4.04s      2018755 ns/op     656 B/op      9 allocs/op
                     Read:   4000     0.44s       110691 ns/op     776 B/op     14 allocs/op
     MultiRead limit 1000:   2000     4.13s      2063212 ns/op  233240 B/op   6025 allocs/op
modl
                   Insert:   2000     4.10s      2052064 ns/op    2137 B/op     56 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert
                   Update:   2000     4.03s      2016194 ns/op    1688 B/op     52 allocs/op
                     Read:   4000     0.62s       155328 ns/op    2208 B/op     56 allocs/op
     MultiRead limit 1000:   2000     7.80s      3897565 ns/op  353337 B/op  17946 allocs/op
hood
                   Insert:   2000     4.13s      2066006 ns/op    8914 B/op    237 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     4.34s      2169109 ns/op   15496 B/op    390 allocs/op
                     Read:   4000     0.97s       242829 ns/op    7040 B/op    204 allocs/op
     MultiRead limit 1000:   2000    19.36s      9680767 ns/op 2010210 B/op  86139 allocs/op
godb
                   Insert:   2000     4.26s      2127786 ns/op    6266 B/op    144 allocs/op
       BulkInsert 100 row:    500     2.22s      4443246 ns/op  341082 B/op   7013 allocs/op
                   Update:   2000     4.13s      2065343 ns/op    7056 B/op    182 allocs/op
                     Read:   4000     0.71s       176409 ns/op    5376 B/op    128 allocs/op
     MultiRead limit 1000:   2000    13.33s      6667297 ns/op 1102239 B/op  39008 allocs/op
gorm
                   Insert:   2000     4.42s      2208930 ns/op    8343 B/op    181 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255
                   Update:   2000     4.43s      2212758 ns/op    8506 B/op    207 allocs/op
                     Read:   4000     0.85s       212769 ns/op    7880 B/op    171 allocs/op
     MultiRead limit 1000:   2000    23.06s     11530125 ns/op 2807338 B/op  65042 allocs/op
genmai
                   Insert:   2000     4.39s      2192525 ns/op    5801 B/op    185 allocs/op
       BulkInsert 100 row:    500     1.89s      3786738 ns/op  223382 B/op   3536 allocs/op
                   Update:   2000     4.13s      2062883 ns/op    4320 B/op    170 allocs/op
                     Read:   4000     0.88s       220733 ns/op    3872 B/op    187 allocs/op
     MultiRead limit 1000:   2000     6.89s      3444379 ns/op  432161 B/op  14119 allocs/op
qbs
                   Insert:   2000     3.92s      1960408 ns/op    6610 B/op    133 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert, err driver: bad connection
                   Update:   2000     4.04s      2022347 ns/op    7017 B/op    163 allocs/op
                     Read:   4000     0.43s       107389 ns/op    8024 B/op    179 allocs/op
     MultiRead limit 1000:   2000    14.31s      7156007 ns/op 1734739 B/op  54135 allocs/op
gorp
                   Insert:   2000     4.06s      2028993 ns/op    2137 B/op     57 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     3.99s      1997102 ns/op    1752 B/op     51 allocs/op
                     Read:   4000     0.70s       176123 ns/op    5305 B/op    213 allocs/op
     MultiRead limit 1000:   2000     7.78s      3890168 ns/op  462424 B/op  16133 allocs/op
raw
                   Insert:   2000     3.89s      1947092 ns/op     848 B/op     25 allocs/op
       BulkInsert 100 row:    500     1.74s      3472205 ns/op  144840 B/op   1432 allocs/op
                   Update:   2000     0.11s        56468 ns/op     856 B/op     25 allocs/op
                     Read:   4000     0.28s        70573 ns/op    1082 B/op     33 allocs/op
     MultiRead limit 1000:   2000     4.62s      2311384 ns/op  283095 B/op  12917 allocs/op
xorm
                   Insert:   2000     4.12s      2058652 ns/op    4348 B/op    120 allocs/op
       BulkInsert 100 row:    500     2.52s      5049331 ns/op 2406760 B/op   8676 allocs/op
                   Update:   2000     4.05s      2026978 ns/op    4064 B/op    138 allocs/op
                     Read:   4000     1.25s       311981 ns/op    9880 B/op    240 allocs/op
     MultiRead limit 1000:   2000    22.02s     11010617 ns/op 1277728 B/op  57111 allocs/op
beego_orm
                   Insert:   2000     4.03s      2015357 ns/op    3168 B/op     75 allocs/op
       BulkInsert 100 row:    500     1.87s      3738687 ns/op  205212 B/op   2860 allocs/op
                   Update:   2000     4.07s      2032942 ns/op    2464 B/op     60 allocs/op
                     Read:   4000     0.67s       168516 ns/op    3200 B/op    104 allocs/op
     MultiRead limit 1000:   2000    11.24s      5619829 ns/op  587550 B/op  35007 allocs/op

Reports:

  2000 times - Insert
       raw:     3.89s      1947092 ns/op     848 B/op     25 allocs/op
       qbs:     3.92s      1960408 ns/op    6610 B/op    133 allocs/op
       dbr:     4.01s      2005213 ns/op    4297 B/op     81 allocs/op
        pg:     4.02s      2008151 ns/op    1990 B/op     10 allocs/op
 beego_orm:     4.03s      2015357 ns/op    3168 B/op     75 allocs/op
      gorp:     4.06s      2028993 ns/op    2137 B/op     57 allocs/op
      sqlx:     4.10s      2051602 ns/op    2769 B/op     65 allocs/op
      modl:     4.10s      2052064 ns/op    2137 B/op     56 allocs/op
      xorm:     4.12s      2058652 ns/op    4348 B/op    120 allocs/op
      hood:     4.13s      2066006 ns/op    8914 B/op    237 allocs/op
      godb:     4.26s      2127786 ns/op    6266 B/op    144 allocs/op
    genmai:     4.39s      2192525 ns/op    5801 B/op    185 allocs/op
      gorm:     4.42s      2208930 ns/op    8343 B/op    181 allocs/op
     db.v3:     6.54s      3270646 ns/op   37549 B/op   1642 allocs/op
       pop:    15.41s      7705992 ns/op   61776 B/op    720 allocs/op

   500 times - BulkInsert 100 row
       dbr:     0.02s        37854 ns/op    2866 B/op     42 allocs/op
       raw:     1.74s      3472205 ns/op  144840 B/op   1432 allocs/op
 beego_orm:     1.87s      3738687 ns/op  205212 B/op   2860 allocs/op
        pg:     1.88s      3766258 ns/op   14507 B/op    212 allocs/op
    genmai:     1.89s      3786738 ns/op  223382 B/op   3536 allocs/op
      godb:     2.22s      4443246 ns/op  341082 B/op   7013 allocs/op
     db.v3:     2.41s      4815923 ns/op  486190 B/op  20014 allocs/op
      xorm:     2.52s      5049331 ns/op 2406760 B/op   8676 allocs/op
      modl:     Don't support bulk insert
      hood:     Problematic bulk insert, too slow
       qbs:     Don't support bulk insert, err driver: bad connection
      gorp:     Problematic bulk insert, too slow
       pop:     Problematic bulk insert, too slow
      sqlx:     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
      gorm:     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255

  2000 times - Update
       raw:     0.11s        56468 ns/op     856 B/op     25 allocs/op
       dbr:     0.19s        95029 ns/op    3158 B/op     58 allocs/op
      gorp:     3.99s      1997102 ns/op    1752 B/op     51 allocs/op
      modl:     4.03s      2016194 ns/op    1688 B/op     52 allocs/op
        pg:     4.04s      2018755 ns/op     656 B/op      9 allocs/op
       qbs:     4.04s      2022347 ns/op    7017 B/op    163 allocs/op
      xorm:     4.05s      2026978 ns/op    4064 B/op    138 allocs/op
 beego_orm:     4.07s      2032942 ns/op    2464 B/op     60 allocs/op
    genmai:     4.13s      2062883 ns/op    4320 B/op    170 allocs/op
      godb:     4.13s      2065343 ns/op    7056 B/op    182 allocs/op
      sqlx:     4.14s      2069727 ns/op    1424 B/op     35 allocs/op
      hood:     4.34s      2169109 ns/op   15496 B/op    390 allocs/op
      gorm:     4.43s      2212758 ns/op    8506 B/op    207 allocs/op
     db.v3:     5.82s      2911505 ns/op   42683 B/op   1937 allocs/op
       pop:     6.28s      3140076 ns/op   32931 B/op    428 allocs/op

  4000 times - Read
       raw:     0.28s        70573 ns/op    1082 B/op     33 allocs/op
      sqlx:     0.37s        92089 ns/op    1936 B/op     47 allocs/op
       dbr:     0.37s        93290 ns/op    2261 B/op     39 allocs/op
       qbs:     0.43s       107389 ns/op    8024 B/op    179 allocs/op
        pg:     0.44s       110691 ns/op     776 B/op     14 allocs/op
      modl:     0.62s       155328 ns/op    2208 B/op     56 allocs/op
 beego_orm:     0.67s       168516 ns/op    3200 B/op    104 allocs/op
      gorp:     0.70s       176123 ns/op    5305 B/op    213 allocs/op
      godb:     0.71s       176409 ns/op    5376 B/op    128 allocs/op
      gorm:     0.85s       212769 ns/op    7880 B/op    171 allocs/op
    genmai:     0.88s       220733 ns/op    3872 B/op    187 allocs/op
     db.v3:     0.95s       236915 ns/op    7105 B/op    296 allocs/op
      hood:     0.97s       242829 ns/op    7040 B/op    204 allocs/op
      xorm:     1.25s       311981 ns/op    9880 B/op    240 allocs/op
       pop:    12.98s      3245945 ns/op   29511 B/op    313 allocs/op

  2000 times - MultiRead limit 1000
        pg:     4.13s      2063212 ns/op  233240 B/op   6025 allocs/op
       raw:     4.62s      2311384 ns/op  283095 B/op  12917 allocs/op
    genmai:     6.89s      3444379 ns/op  432161 B/op  14119 allocs/op
      sqlx:     7.69s      3842905 ns/op  330543 B/op  14956 allocs/op
      gorp:     7.78s      3890168 ns/op  462424 B/op  16133 allocs/op
      modl:     7.80s      3897565 ns/op  353337 B/op  17946 allocs/op
     db.v3:     8.18s      4091602 ns/op  463532 B/op  15300 allocs/op
 beego_orm:    11.24s      5619829 ns/op  587550 B/op  35007 allocs/op
       pop:    12.10s      6048901 ns/op  406315 B/op  16231 allocs/op
      godb:    13.33s      6667297 ns/op 1102239 B/op  39008 allocs/op
       qbs:    14.31s      7156007 ns/op 1734739 B/op  54135 allocs/op
       dbr:    16.90s      8449983 ns/op 2071583 B/op  45961 allocs/op
      hood:    19.36s      9680767 ns/op 2010210 B/op  86139 allocs/op
      xorm:    22.02s     11010617 ns/op 1277728 B/op  57111 allocs/op
      gorm:    23.06s     11530125 ns/op 2807338 B/op  65042 allocs/op
```
=======
# gobenchorm
GO ORM/ODM benchmarks
>>>>>>> da8af2c82c284032bcfae7dfb4fd6feff375b449
