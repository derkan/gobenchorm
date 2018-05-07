# GO ORM Benchmarks

[![Build Status](https://travis-ci.com/derkan/gobenchorm.svg?branch=master)](https://travis-ci.com/derkan/gobenchorm) [![GoDoc](https://godoc.org/github.com/derkan/gobenchorm/benchs?status.svg)](https://godoc.org/github.com/derkan/gobenchorm/benchs)

### About

ORM benchmarks for GoLang. Originally forked from [orm-benchmark](https://github.com/milkpod/orm-benchmark).
Contributions are wellcome.

### Environment

* go version go1.9 linux/amd64

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

#### QBS
- `qbs` needs patch for reflecting `string` values:

    `github.com/coocood/qbs/base.go:69` should be patched to:

    `fieldValue.SetString(string(driverValue.Elem().String()))`
- `BulkInsert` is not working as expected.


#### Gorm
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
cd gobenchorm/cmd
go install
# build
cd gobenchorm/cmd
go build
# all
cd gobenchorm/cmd
./gobenchorm -multi=1 -orm=all
# portion
./gobenchorm -multi=1 -orm=xorm -orm=raw -orm=godb
```

### Reports

```
godb
                   Insert:   2000     4.15s      2077096 ns/op    6057 B/op    139 allocs/op
       BulkInsert 100 row:    500     2.13s      4268255 ns/op  328261 B/op   6712 allocs/op
                   Update:   2000     4.19s      2097205 ns/op    6832 B/op    177 allocs/op
                     Read:   4000     0.68s       170512 ns/op    5184 B/op    123 allocs/op
     MultiRead limit 1000:   2000    11.84s      5921726 ns/op  988936 B/op  36004 allocs/op
gorp
                   Insert:   2000     4.08s      2037957 ns/op    2137 B/op     57 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     4.04s      2022381 ns/op    1752 B/op     51 allocs/op
                     Read:   4000     0.69s       172499 ns/op    5308 B/op    213 allocs/op
     MultiRead limit 1000:   2000     7.63s      3815172 ns/op  462660 B/op  16134 allocs/op
modl
                   Insert:   2000     4.10s      2050692 ns/op    2137 B/op     56 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert
                   Update:   2000     4.06s      2028871 ns/op    1688 B/op     52 allocs/op
                     Read:   4000     0.63s       158635 ns/op    2209 B/op     56 allocs/op
     MultiRead limit 1000:   2000     7.52s      3758069 ns/op  353503 B/op  17947 allocs/op
gorm
                   Insert:   2000     4.33s      2165074 ns/op    9650 B/op    181 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255
                   Update:   2000     4.35s      2174938 ns/op    8512 B/op    207 allocs/op
                     Read:   4000     0.82s       204606 ns/op    7884 B/op    171 allocs/op
     MultiRead limit 1000:   2000    21.91s     10957177 ns/op 2807977 B/op  65047 allocs/op
beego_orm
                   Insert:   2000     4.11s      2056054 ns/op    3169 B/op     75 allocs/op
       BulkInsert 100 row:    500     1.85s      3704520 ns/op  205420 B/op   2860 allocs/op
                   Update:   2000     4.04s      2021564 ns/op    2464 B/op     60 allocs/op
                     Read:   4000     0.68s       170224 ns/op    3201 B/op    104 allocs/op
     MultiRead limit 1000:   2000    10.85s      5422641 ns/op  587895 B/op  35009 allocs/op
pg
                   Insert:   2000     4.06s      2030079 ns/op     688 B/op     10 allocs/op
       BulkInsert 100 row:    500     1.89s      3783487 ns/op   14507 B/op    212 allocs/op
                   Update:   2000     4.02s      2012278 ns/op     656 B/op      9 allocs/op
                     Read:   4000     0.38s        95899 ns/op     776 B/op     14 allocs/op
     MultiRead limit 1000:   2000     4.29s      2145322 ns/op  233240 B/op   6025 allocs/op
db.v3
                   Insert:   2000     5.88s      2938549 ns/op   37563 B/op   1642 allocs/op
       BulkInsert 100 row:    500     2.39s      4777852 ns/op  486200 B/op  20014 allocs/op
                   Update:   2000     5.80s      2900725 ns/op   42698 B/op   1937 allocs/op
                     Read:   4000     0.92s       229124 ns/op    7105 B/op    296 allocs/op
     MultiRead limit 1000:   2000     8.00s      4000497 ns/op  463540 B/op  15300 allocs/op
qbs
                   Insert:   2000     3.95s      1977438 ns/op    6870 B/op    133 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert, err driver: bad connection
                   Update:   2000     4.07s      2032930 ns/op    7278 B/op    163 allocs/op
                     Read:   4000     0.42s       103880 ns/op    8280 B/op    179 allocs/op
     MultiRead limit 1000:   2000    12.85s      6424494 ns/op 1735061 B/op  54136 allocs/op
pop
                   Insert:   2000    14.69s      7343211 ns/op   61768 B/op    720 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     5.57s      2782799 ns/op   32931 B/op    428 allocs/op
                     Read:   4000    10.33s      2581398 ns/op   29504 B/op    313 allocs/op
     MultiRead limit 1000:   2000    11.73s      5865713 ns/op  406412 B/op  16232 allocs/op
hood
                   Insert:   2000     4.22s      2109764 ns/op    8914 B/op    237 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     4.19s      2097071 ns/op   15496 B/op    390 allocs/op
                     Read:   4000     0.91s       227133 ns/op    7040 B/op    204 allocs/op
     MultiRead limit 1000:   2000    18.56s      9280663 ns/op 2010209 B/op  86139 allocs/op
xorm
                   Insert:   2000     4.16s      2081905 ns/op    4348 B/op    120 allocs/op
       BulkInsert 100 row:    500     2.57s      5134646 ns/op 2406765 B/op   8676 allocs/op
                   Update:   2000     4.03s      2014333 ns/op    4064 B/op    138 allocs/op
                     Read:   4000     1.20s       299337 ns/op   10168 B/op    252 allocs/op
     MultiRead limit 1000:   2000    23.56s     11780300 ns/op 1277733 B/op  57111 allocs/op
raw
                   Insert:   2000     4.01s      2006702 ns/op     848 B/op     25 allocs/op
       BulkInsert 100 row:    500     1.76s      3529739 ns/op  144840 B/op   1432 allocs/op
                   Update:   2000     0.11s        56153 ns/op     856 B/op     25 allocs/op
                     Read:   4000     0.31s        77569 ns/op    1082 B/op     33 allocs/op
     MultiRead limit 1000:   2000     4.40s      2200151 ns/op  283096 B/op  12918 allocs/op
sqlx
                   Insert:   2000     4.15s      2073620 ns/op    2768 B/op     65 allocs/op
       BulkInsert 100 row:    500     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
                   Update:   2000     4.06s      2030586 ns/op    1424 B/op     35 allocs/op
                     Read:   4000     0.37s        92822 ns/op    1936 B/op     47 allocs/op
     MultiRead limit 1000:   2000     7.49s      3745893 ns/op  330544 B/op  14957 allocs/op
genmai
                   Insert:   2000     4.45s      2223445 ns/op    5802 B/op    185 allocs/op
       BulkInsert 100 row:    500     1.95s      3905387 ns/op  223381 B/op   3536 allocs/op
                   Update:   2000     4.16s      2081260 ns/op    4320 B/op    170 allocs/op
                     Read:   4000     0.79s       198012 ns/op    3872 B/op    187 allocs/op
     MultiRead limit 1000:   2000     6.83s      3413104 ns/op  432162 B/op  14119 allocs/op
dbr
                   Insert:   2000     4.49s      2243182 ns/op    4313 B/op     81 allocs/op
       BulkInsert 100 row:    500     0.01s        27247 ns/op    2866 B/op     42 allocs/op
                   Update:   2000     0.20s        98571 ns/op    3157 B/op     58 allocs/op
                     Read:   4000     0.39s        97510 ns/op    2261 B/op     39 allocs/op
     MultiRead limit 1000:   2000    16.03s      8013595 ns/op 2070974 B/op  45957 allocs/op

Reports:

  2000 times - Insert
       qbs:     3.95s      1977438 ns/op    6870 B/op    133 allocs/op
       raw:     4.01s      2006702 ns/op     848 B/op     25 allocs/op
        pg:     4.06s      2030079 ns/op     688 B/op     10 allocs/op
      gorp:     4.08s      2037957 ns/op    2137 B/op     57 allocs/op
      modl:     4.10s      2050692 ns/op    2137 B/op     56 allocs/op
 beego_orm:     4.11s      2056054 ns/op    3169 B/op     75 allocs/op
      sqlx:     4.15s      2073620 ns/op    2768 B/op     65 allocs/op
      godb:     4.15s      2077096 ns/op    6057 B/op    139 allocs/op
      xorm:     4.16s      2081905 ns/op    4348 B/op    120 allocs/op
      hood:     4.22s      2109764 ns/op    8914 B/op    237 allocs/op
      gorm:     4.33s      2165074 ns/op    9650 B/op    181 allocs/op
    genmai:     4.45s      2223445 ns/op    5802 B/op    185 allocs/op
       dbr:     4.49s      2243182 ns/op    4313 B/op     81 allocs/op
     db.v3:     5.88s      2938549 ns/op   37563 B/op   1642 allocs/op
       pop:    14.69s      7343211 ns/op   61768 B/op    720 allocs/op

   500 times - BulkInsert 100 row
       dbr:     0.01s        27247 ns/op    2866 B/op     42 allocs/op
       raw:     1.76s      3529739 ns/op  144840 B/op   1432 allocs/op
 beego_orm:     1.85s      3704520 ns/op  205420 B/op   2860 allocs/op
        pg:     1.89s      3783487 ns/op   14507 B/op    212 allocs/op
    genmai:     1.95s      3905387 ns/op  223381 B/op   3536 allocs/op
      godb:     2.13s      4268255 ns/op  328261 B/op   6712 allocs/op
     db.v3:     2.39s      4777852 ns/op  486200 B/op  20014 allocs/op
      xorm:     2.57s      5134646 ns/op 2406765 B/op   8676 allocs/op
      gorm:     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255
       pop:     Problematic bulk insert, too slow
      hood:     Problematic bulk insert, too slow
      modl:     Don't support bulk insert
      sqlx:     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
      gorp:     Problematic bulk insert, too slow
       qbs:     Don't support bulk insert, err driver: bad connection

  2000 times - Update
       raw:     0.11s        56153 ns/op     856 B/op     25 allocs/op
       dbr:     0.20s        98571 ns/op    3157 B/op     58 allocs/op
        pg:     4.02s      2012278 ns/op     656 B/op      9 allocs/op
      xorm:     4.03s      2014333 ns/op    4064 B/op    138 allocs/op
 beego_orm:     4.04s      2021564 ns/op    2464 B/op     60 allocs/op
      gorp:     4.04s      2022381 ns/op    1752 B/op     51 allocs/op
      modl:     4.06s      2028871 ns/op    1688 B/op     52 allocs/op
      sqlx:     4.06s      2030586 ns/op    1424 B/op     35 allocs/op
       qbs:     4.07s      2032930 ns/op    7278 B/op    163 allocs/op
    genmai:     4.16s      2081260 ns/op    4320 B/op    170 allocs/op
      hood:     4.19s      2097071 ns/op   15496 B/op    390 allocs/op
      godb:     4.19s      2097205 ns/op    6832 B/op    177 allocs/op
      gorm:     4.35s      2174938 ns/op    8512 B/op    207 allocs/op
       pop:     5.57s      2782799 ns/op   32931 B/op    428 allocs/op
     db.v3:     5.80s      2900725 ns/op   42698 B/op   1937 allocs/op

  4000 times - Read
       raw:     0.31s        77569 ns/op    1082 B/op     33 allocs/op
      sqlx:     0.37s        92822 ns/op    1936 B/op     47 allocs/op
        pg:     0.38s        95899 ns/op     776 B/op     14 allocs/op
       dbr:     0.39s        97510 ns/op    2261 B/op     39 allocs/op
       qbs:     0.42s       103880 ns/op    8280 B/op    179 allocs/op
      modl:     0.63s       158635 ns/op    2209 B/op     56 allocs/op
 beego_orm:     0.68s       170224 ns/op    3201 B/op    104 allocs/op
      godb:     0.68s       170512 ns/op    5184 B/op    123 allocs/op
      gorp:     0.69s       172499 ns/op    5308 B/op    213 allocs/op
    genmai:     0.79s       198012 ns/op    3872 B/op    187 allocs/op
      gorm:     0.82s       204606 ns/op    7884 B/op    171 allocs/op
      hood:     0.91s       227133 ns/op    7040 B/op    204 allocs/op
     db.v3:     0.92s       229124 ns/op    7105 B/op    296 allocs/op
      xorm:     1.20s       299337 ns/op   10168 B/op    252 allocs/op
       pop:    10.33s      2581398 ns/op   29504 B/op    313 allocs/op

  2000 times - MultiRead limit 1000
        pg:     4.29s      2145322 ns/op  233240 B/op   6025 allocs/op
       raw:     4.40s      2200151 ns/op  283096 B/op  12918 allocs/op
    genmai:     6.83s      3413104 ns/op  432162 B/op  14119 allocs/op
      sqlx:     7.49s      3745893 ns/op  330544 B/op  14957 allocs/op
      modl:     7.52s      3758069 ns/op  353503 B/op  17947 allocs/op
      gorp:     7.63s      3815172 ns/op  462660 B/op  16134 allocs/op
     db.v3:     8.00s      4000497 ns/op  463540 B/op  15300 allocs/op
 beego_orm:    10.85s      5422641 ns/op  587895 B/op  35009 allocs/op
       pop:    11.73s      5865713 ns/op  406412 B/op  16232 allocs/op
      godb:    11.84s      5921726 ns/op  988936 B/op  36004 allocs/op
       qbs:    12.85s      6424494 ns/op 1735061 B/op  54136 allocs/op
       dbr:    16.03s      8013595 ns/op 2070974 B/op  45957 allocs/op
      hood:    18.56s      9280663 ns/op 2010209 B/op  86139 allocs/op
      gorm:    21.91s     10957177 ns/op 2807977 B/op  65047 allocs/op
      xorm:    23.56s     11780300 ns/op 1277733 B/op  57111 allocs/op
```
