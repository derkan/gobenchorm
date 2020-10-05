# GO ORM Benchmarks

[![Build Status](https://travis-ci.com/derkan/gobenchorm.svg?branch=master)](https://travis-ci.com/derkan/gobenchorm) [![GoDoc](https://godoc.org/github.com/derkan/gobenchorm/benchs?status.svg)](https://godoc.org/github.com/derkan/gobenchorm/benchs) [![Coverage Status](https://coveralls.io/repos/github/derkan/gobenchorm/badge.svg?branch=master)](https://coveralls.io/github/derkan/gobenchorm?branch=master)

## About

ORM benchmarks for GoLang. Originally forked from [orm-benchmark](https://github.com/milkpod/orm-benchmark).
Contributions are wellcome.

## Environment

- go version go1.9 linux/amd64

## PostgreSQL

- PostgreSQL 12.4 for Linux on WSL2

## ORMs

- [dbr](https://github.com/gocraft/dbr/v2)
- [genmai](https://github.com/naoina/genmai)
- [gorm](https://github.com/jinzhu/gorm)
- [gorp](http://gopkg.in/gorp.v3)
- [pg](https://github.com/go-pg/pg/v10)
- [beego](https://github.com/astaxie/beego/tree/master/orm)
- [sqlx](https://github.com/jmoiron/sqlx)
- [xorm](https://github.com/xormplus/xorm)
- [godb](https://github.com/samonzeweb/godb)
- [upper](https://github.com/upper/db/v4)
- [hood](https://github.com/eaigner/hood)
- [modl](https://github.com/jmoiron/modl)
- [qbs](https://github.com/coocood/qbs)
- [pop](https://github.com/gobuffalo/pop)
- [rel](https://github.com/go-rel/rel)

### Notes

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

### Prepare DB

```sql
CREATE ROLE bench LOGIN PASSWORD 'pass'
   VALID UNTIL 'infinity';
CREATE DATABASE benchdb
  WITH OWNER = bench;
```

### Run

```go
go get github.com/derkan/gobenchorm

# build:
cd gobenchorm/cmd
go build

# run all benchmarks:
./gobenchorm -multi=1 -orm=all

# run given benchmarks:
./gobenchorm -multi=1 -orm=xorm -orm=raw -orm=godb
```

### Reports

```yaml
gorp
                   Insert:   2000     6.87s      3432760 ns/op    1688 B/op     44 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     7.00s      3497929 ns/op    1344 B/op     39 allocs/op
                     Read:   4000     0.84s       211248 ns/op    3952 B/op    188 allocs/op
     MultiRead limit 1000:   2000     6.07s      3036872 ns/op  735574 B/op  15853 allocs/op
dbr
                   Insert:   2000     6.42s      3208658 ns/op    2984 B/op     74 allocs/op
      BulkInsert 100 rows:    500     0.02s        30645 ns/op    2081 B/op     39 allocs/op
                   Update:   2000     0.40s       201692 ns/op    2619 B/op     57 allocs/op
                     Read:   4000     0.79s       197834 ns/op    2176 B/op     36 allocs/op
     MultiRead limit 1000:   2000     6.19s      3093118 ns/op  514960 B/op  16705 allocs/op
genmai
                   Insert:   2000     8.01s      4006918 ns/op    4503 B/op    148 allocs/op
       BulkInsert 100 row:    500     3.12s      6246188 ns/op  204861 B/op   3066 allocs/op
                   Update:   2000     7.41s      3703812 ns/op    3521 B/op    146 allocs/op
                     Read:   4000     1.87s       466538 ns/op    3313 B/op    171 allocs/op
     MultiRead limit 1000:   2000     5.60s      2802090 ns/op  420667 B/op  12844 allocs/op
pop
                   Insert:   2000     7.53s      3765400 ns/op   10327 B/op    248 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     6.76s      3377949 ns/op    6794 B/op    197 allocs/op
                     Read:   4000     0.95s       236338 ns/op    3668 B/op     72 allocs/op
     MultiRead limit 1000:   2000     7.50s      3748938 ns/op  693135 B/op  14755 allocs/op
qbs
                   Insert:   2000     6.41s      3203275 ns/op    5681 B/op    123 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert, err driver: bad connection
                   Update:   2000     7.06s      3528732 ns/op    5899 B/op    149 allocs/op
                     Read:   4000     reflect: call of reflect.Value.Bytes on string Value
     MultiRead limit 1000:   2000     reflect: call of reflect.Value.Bytes on string Value
upper
                   Insert:   2000     8.89s      4445811 ns/op   27757 B/op   1184 allocs/op
       BulkInsert 100 row:    500     3.83s      7665536 ns/op  482127 B/op  19821 allocs/op
                   Update:   2000     8.95s      4476088 ns/op   33013 B/op   1491 allocs/op
                     Read:   4000     1.59s       398329 ns/op    7385 B/op    293 allocs/op
     MultiRead limit 1000:   2000     6.92s      3462435 ns/op  648016 B/op  14047 allocs/op
gorm
                   Insert:   2000     7.40s      3698926 ns/op    6842 B/op     97 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255
                   Update:   2000     7.33s      3665767 ns/op    7468 B/op     93 allocs/op
                     Read:   4000     0.76s       190584 ns/op    4612 B/op     93 allocs/op
     MultiRead limit 1000:   2000    11.50s      5749897 ns/op  876110 B/op  36740 allocs/op
rel
                   Insert:   2000     6.90s      3451333 ns/op    2447 B/op     49 allocs/op
       BulkInsert 100 row:    500     4.32s      8647441 ns/op  287076 B/op   4053 allocs/op
                   Update:   2000     8.32s      4159590 ns/op    2608 B/op     50 allocs/op
                     Read:   4000     0.93s       231918 ns/op    1616 B/op     44 allocs/op
     MultiRead limit 1000:   2000     9.10s      4551945 ns/op 1010637 B/op  24674 allocs/op
hood
                   Insert:   2000     6.99s      3493561 ns/op    7089 B/op    173 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     7.29s      3643943 ns/op   13483 B/op    324 allocs/op
                     Read:   4000     reflect: call of reflect.Value.Bytes on string Value
     MultiRead limit 1000:   2000     reflect: call of reflect.Value.Bytes on string Value
pg
                   Insert:   2000     6.63s      3316215 ns/op    1596 B/op     11 allocs/op
       BulkInsert 100 row:    500     2.94s      5878255 ns/op   17011 B/op    214 allocs/op
                   Update:   2000     6.50s      3251419 ns/op     992 B/op     13 allocs/op
                     Read:   4000     0.77s       191609 ns/op    1000 B/op     14 allocs/op
     MultiRead limit 1000:   2000     3.84s      1921496 ns/op  320023 B/op   5027 allocs/op
raw
                   Insert:   2000     6.14s      3071862 ns/op     696 B/op     18 allocs/op
       BulkInsert 100 row:    500     0.00s      1.40 ns/op       0 B/op      0 allocs/op
                   Update:   2000     0.34s       170900 ns/op     712 B/op     19 allocs/op
                     Read:   4000     0.64s       160831 ns/op     888 B/op     24 allocs/op
     MultiRead limit 1000:   2000     4.10s      2049903 ns/op  237312 B/op   6026 allocs/op
beego
                   Insert:   2000     6.43s      3213887 ns/op    2424 B/op     56 allocs/op
       BulkInsert 100 row:    500     2.86s      5726498 ns/op  196505 B/op   2845 allocs/op
                   Update:   2000     6.37s      3182599 ns/op    1801 B/op     47 allocs/op
                     Read:   4000     0.72s       178911 ns/op    2112 B/op     75 allocs/op
     MultiRead limit 1000:   2000     7.66s      3831261 ns/op  746765 B/op  32474 allocs/op
godb
                   Insert:   2000     6.82s      3412419 ns/op    4730 B/op    115 allocs/op
       BulkInsert 100 row:    500     3.26s      6516848 ns/op  289701 B/op   5994 allocs/op
                   Update:   2000     6.86s      3432064 ns/op    5377 B/op    154 allocs/op
                     Read:   4000     1.48s       370590 ns/op    4193 B/op    102 allocs/op
     MultiRead limit 1000:   2000     9.31s      4652732 ns/op  997678 B/op  31738 allocs/op
sqlx
                   Insert:   2000     7.46s      3731319 ns/op    2319 B/op     51 allocs/op
       BulkInsert 100 row:    500     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
                   Update:   2000     8.37s      4185389 ns/op    1016 B/op     21 allocs/op
                     Read:   4000     1.02s       253876 ns/op    1744 B/op     38 allocs/op
     MultiRead limit 1000:   2000     6.99s      3496744 ns/op  499425 B/op  13691 allocs/op
modl
                   Insert:   2000     7.03s      3516723 ns/op    1688 B/op     43 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert
                   Update:   2000     7.91s      3952784 ns/op    1296 B/op     40 allocs/op
                     Read:   4000     1.63s       407979 ns/op    1776 B/op     41 allocs/op
     MultiRead limit 1000:   2000     6.82s      3410109 ns/op  514024 B/op  16675 allocs/op
xorm
                   Insert:   2000     7.05s      3522534 ns/op    3164 B/op     98 allocs/op
       BulkInsert 100 row:    500     3.43s      6861290 ns/op  319570 B/op   7541 allocs/op
                   Update:   2000     7.60s      3798531 ns/op    3217 B/op    126 allocs/op
                     Read:   4000     2.12s       530698 ns/op    8795 B/op    252 allocs/op
     MultiRead limit 1000:   2000    18.23s      9114290 ns/op 1267390 B/op  55857 allocs/op

Reports:

  2000 times - Insert
       raw:     6.14s      3071862 ns/op     696 B/op     18 allocs/op
       qbs:     6.41s      3203275 ns/op    5681 B/op    123 allocs/op
       dbr:     6.42s      3208658 ns/op    2984 B/op     74 allocs/op
     beego:     6.43s      3213887 ns/op    2424 B/op     56 allocs/op
        pg:     6.63s      3316215 ns/op    1596 B/op     11 allocs/op
      godb:     6.82s      3412419 ns/op    4730 B/op    115 allocs/op
      gorp:     6.87s      3432760 ns/op    1688 B/op     44 allocs/op
       rel:     6.90s      3451333 ns/op    2447 B/op     49 allocs/op
      hood:     6.99s      3493561 ns/op    7089 B/op    173 allocs/op
      modl:     7.03s      3516723 ns/op    1688 B/op     43 allocs/op
      xorm:     7.05s      3522534 ns/op    3164 B/op     98 allocs/op
      gorm:     7.40s      3698926 ns/op    6842 B/op     97 allocs/op
      sqlx:     7.46s      3731319 ns/op    2319 B/op     51 allocs/op
       pop:     7.53s      3765400 ns/op   10327 B/op    248 allocs/op
    genmai:     8.01s      4006918 ns/op    4503 B/op    148 allocs/op
     upper:     8.89s      4445811 ns/op   27757 B/op   1184 allocs/op

   500 times - BulkInsert 100 row
       dbr:     0.02s        30645 ns/op    2081 B/op     39 allocs/op
     beego:     2.86s      5726498 ns/op  196505 B/op   2845 allocs/op
        pg:     2.94s      5878255 ns/op   17011 B/op    214 allocs/op
    genmai:     3.12s      6246188 ns/op  204861 B/op   3066 allocs/op
      godb:     3.26s      6516848 ns/op  289701 B/op   5994 allocs/op
      xorm:     3.43s      6861290 ns/op  319570 B/op   7541 allocs/op
     upper:     3.83s      7665536 ns/op  482127 B/op  19821 allocs/op
       rel:     4.32s      8647441 ns/op  287076 B/op   4053 allocs/op
      hood:     Problematic bulk insert, too slow
      gorm:     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255
       raw:     0.00s      1.40 ns/op       0 B/op      0 allocs/op
       qbs:     Don't support bulk insert, err driver: bad connection
       pop:     Problematic bulk insert, too slow
      sqlx:     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
      modl:     Don't support bulk insert
      gorp:     Problematic bulk insert, too slow

  2000 times - Update
       raw:     0.34s       170900 ns/op     712 B/op     19 allocs/op
       dbr:     0.40s       201692 ns/op    2619 B/op     57 allocs/op
     beego:     6.37s      3182599 ns/op    1801 B/op     47 allocs/op
        pg:     6.50s      3251419 ns/op     992 B/op     13 allocs/op
       pop:     6.76s      3377949 ns/op    6794 B/op    197 allocs/op
      godb:     6.86s      3432064 ns/op    5377 B/op    154 allocs/op
      gorp:     7.00s      3497929 ns/op    1344 B/op     39 allocs/op
       qbs:     7.06s      3528732 ns/op    5899 B/op    149 allocs/op
      hood:     7.29s      3643943 ns/op   13483 B/op    324 allocs/op
      gorm:     7.33s      3665767 ns/op    7468 B/op     93 allocs/op
    genmai:     7.41s      3703812 ns/op    3521 B/op    146 allocs/op
      xorm:     7.60s      3798531 ns/op    3217 B/op    126 allocs/op
      modl:     7.91s      3952784 ns/op    1296 B/op     40 allocs/op
       rel:     8.32s      4159590 ns/op    2608 B/op     50 allocs/op
      sqlx:     8.37s      4185389 ns/op    1016 B/op     21 allocs/op
     upper:     8.95s      4476088 ns/op   33013 B/op   1491 allocs/op

  4000 times - Read
       raw:     0.64s       160831 ns/op     888 B/op     24 allocs/op
     beego:     0.72s       178911 ns/op    2112 B/op     75 allocs/op
      gorm:     0.76s       190584 ns/op    4612 B/op     93 allocs/op
        pg:     0.77s       191609 ns/op    1000 B/op     14 allocs/op
       dbr:     0.79s       197834 ns/op    2176 B/op     36 allocs/op
      gorp:     0.84s       211248 ns/op    3952 B/op    188 allocs/op
       rel:     0.93s       231918 ns/op    1616 B/op     44 allocs/op
       pop:     0.95s       236338 ns/op    3668 B/op     72 allocs/op
      sqlx:     1.02s       253876 ns/op    1744 B/op     38 allocs/op
      godb:     1.48s       370590 ns/op    4193 B/op    102 allocs/op
     upper:     1.59s       398329 ns/op    7385 B/op    293 allocs/op
      modl:     1.63s       407979 ns/op    1776 B/op     41 allocs/op
    genmai:     1.87s       466538 ns/op    3313 B/op    171 allocs/op
      xorm:     2.12s       530698 ns/op    8795 B/op    252 allocs/op
       qbs:     reflect: call of reflect.Value.Bytes on string Value
      hood:     reflect: call of reflect.Value.Bytes on string Value

  2000 times - MultiRead limit 1000
       raw:     3.89s      1943745 ns/op  272017 B/op  11657 allocs/op
        pg:     4.10s      2049903 ns/op  237312 B/op   6026 allocs/op
    genmai:     5.60s      2802090 ns/op  420667 B/op  12844 allocs/op
      gorp:     6.07s      3036872 ns/op  735574 B/op  15853 allocs/op
       dbr:     6.19s      3093118 ns/op  514960 B/op  16705 allocs/op
      modl:     6.82s      3410109 ns/op  514024 B/op  16675 allocs/op
     upper:     6.92s      3462435 ns/op  648016 B/op  14047 allocs/op
      sqlx:     6.99s      3496744 ns/op  499425 B/op  13691 allocs/op
       pop:     7.50s      3748938 ns/op  693135 B/op  14755 allocs/op
     beego:     7.66s      3831261 ns/op  746765 B/op  32474 allocs/op
       rel:     9.10s      4551945 ns/op 1010637 B/op  24674 allocs/op
      godb:     9.31s      4652732 ns/op  997678 B/op  31738 allocs/op
      gorm:    11.50s      5749897 ns/op  876110 B/op  36740 allocs/op
      xorm:    18.23s      9114290 ns/op 1267390 B/op  55857 allocs/op
```
