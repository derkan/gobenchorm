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
raw
                   Insert:   2000     7.08s      3538114 ns/op     696 B/op     18 allocs/op
pq: there is no parameter $1
       BulkInsert 100 row:    500     0.00s      1.20 ns/op       0 B/op      0 allocs/op
                   Update:   2000     0.32s       158640 ns/op     712 B/op     19 allocs/op
                     Read:   4000     0.65s       163536 ns/op     888 B/op     24 allocs/op
     MultiRead limit 1000:   2000     4.63s      2316797 ns/op  272016 B/op  11657 allocs/op
dbr
                   Insert:   2000     6.36s      3178876 ns/op    2985 B/op     74 allocs/op
      BulkInsert 100 rows:    500     0.02s        31893 ns/op    2081 B/op     39 allocs/op
                   Update:   2000     0.39s       195959 ns/op    2619 B/op     57 allocs/op
                     Read:   4000     1.10s       274037 ns/op    2176 B/op     36 allocs/op
     MultiRead limit 1000:   2000     6.75s      3373897 ns/op  514960 B/op  16705 allocs/op
gorp
                   Insert:   2000     7.26s      3632476 ns/op    1688 B/op     44 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     6.72s      3361338 ns/op    1344 B/op     39 allocs/op
                     Read:   4000     0.85s       212743 ns/op    3952 B/op    188 allocs/op
     MultiRead limit 1000:   2000     6.05s      3026069 ns/op  736514 B/op  15861 allocs/op
godb
                   Insert:   2000     7.77s      3885032 ns/op    4730 B/op    115 allocs/op
       BulkInsert 100 row:    500     3.40s      6791558 ns/op  289700 B/op   5994 allocs/op
                   Update:   2000     7.33s      3666445 ns/op    5377 B/op    154 allocs/op
                     Read:   4000     1.57s       391860 ns/op    4192 B/op    102 allocs/op
     MultiRead limit 1000:   2000     9.53s      4764432 ns/op  997680 B/op  31738 allocs/op
rel
                   Insert:   2000     6.95s      3477098 ns/op    2447 B/op     49 allocs/op
       BulkInsert 100 row:    500     3.28s      6566187 ns/op  287076 B/op   4053 allocs/op
                   Update:   2000     7.15s      3575988 ns/op    2608 B/op     50 allocs/op
                     Read:   4000     0.80s       200248 ns/op    1616 B/op     44 allocs/op
     MultiRead limit 1000:   2000     9.57s      4786094 ns/op 1010636 B/op  24674 allocs/op
qbs
                   Insert:   2000     7.52s      3757715 ns/op    5681 B/op    123 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert, err driver: bad connection
                   Update:   2000     8.84s      4420961 ns/op    5898 B/op    149 allocs/op
                     Read:   4000     reflect: call of reflect.Value.Bytes on string Value
     MultiRead limit 1000:   2000     reflect: call of reflect.Value.Bytes on string Value
genmai
                   Insert:   2000    10.05s      5026864 ns/op    4502 B/op    148 allocs/op
       BulkInsert 100 row:    500     3.30s      6598140 ns/op  205003 B/op   3066 allocs/op
                   Update:   2000     9.23s      4615697 ns/op    3521 B/op    146 allocs/op
                     Read:   4000     1.89s       472793 ns/op    3313 B/op    171 allocs/op
     MultiRead limit 1000:   2000     7.28s      3640718 ns/op  420674 B/op  12845 allocs/op
hood
                   Insert:   2000     9.70s      4850088 ns/op    7088 B/op    173 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     9.35s      4673388 ns/op   13481 B/op    324 allocs/op
                     Read:   4000     reflect: call of reflect.Value.Bytes on string Value
     MultiRead limit 1000:   2000     reflect: call of reflect.Value.Bytes on string Value
gorm
                   Insert:   2000     9.00s      4499248 ns/op    6853 B/op     97 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255
                   Update:   2000    13.96s      6978174 ns/op    7556 B/op     93 allocs/op
                     Read:   4000     1.55s       386842 ns/op    4612 B/op     93 allocs/op
     MultiRead limit 1000:   2000    12.56s      6279390 ns/op  876112 B/op  36740 allocs/op
beego
                   Insert:   2000     7.53s      3766770 ns/op    2424 B/op     56 allocs/op
       BulkInsert 100 row:    500     2.91s      5828279 ns/op  196637 B/op   2845 allocs/op
                   Update:   2000     7.62s      3811246 ns/op    1801 B/op     47 allocs/op
                     Read:   4000     1.45s       362332 ns/op    2113 B/op     75 allocs/op
     MultiRead limit 1000:   2000    12.24s      6118856 ns/op  746817 B/op  32474 allocs/op
xorm
                   Insert:   2000     6.98s      3492309 ns/op    3164 B/op     98 allocs/op
       BulkInsert 100 row:    500     3.63s      7263580 ns/op  319965 B/op   7542 allocs/op
                   Update:   2000     7.77s      3883802 ns/op    3217 B/op    126 allocs/op
                     Read:   4000     1.75s       438095 ns/op    8795 B/op    252 allocs/op
     MultiRead limit 1000:   2000    22.47s     11233651 ns/op 1447692 B/op  55858 allocs/op
pop
                   Insert:   2000     8.61s      4305478 ns/op   10335 B/op    247 allocs/op
       BulkInsert 100 row:    500     Problematic bulk insert, too slow
                   Update:   2000     6.52s      3258479 ns/op    6795 B/op    197 allocs/op
                     Read:   4000     0.76s       189248 ns/op    3669 B/op     72 allocs/op
     MultiRead limit 1000:   2000    10.09s      5045536 ns/op  690531 B/op  14754 allocs/op
sqlx
                   Insert:   2000     7.58s      3788224 ns/op    2319 B/op     51 allocs/op
       BulkInsert 100 row:    500     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
                   Update:   2000     7.18s      3590337 ns/op    1016 B/op     21 allocs/op
                     Read:   4000     0.83s       208223 ns/op    1744 B/op     38 allocs/op
     MultiRead limit 1000:   2000     7.00s      3501095 ns/op  499424 B/op  13691 allocs/op
pg
                   Insert:   2000     6.86s      3429856 ns/op    1559 B/op     11 allocs/op
       BulkInsert 100 row:    500     3.19s      6378572 ns/op   14913 B/op    214 allocs/op
                   Update:   2000     7.00s      3500675 ns/op     992 B/op     13 allocs/op
                     Read:   4000     0.87s       216339 ns/op    1262 B/op     14 allocs/op
     MultiRead limit 1000:   2000     4.57s      2284203 ns/op  319984 B/op   5027 allocs/op
upper
                   Insert:   2000     9.94s      4969764 ns/op   27734 B/op   1184 allocs/op
       BulkInsert 100 row:    500     4.01s      8027425 ns/op  482063 B/op  19820 allocs/op
                   Update:   2000     9.98s      4989815 ns/op   33009 B/op   1491 allocs/op
                     Read:   4000     1.90s       474667 ns/op    7385 B/op    293 allocs/op
     MultiRead limit 1000:   2000     7.90s      3949170 ns/op  647963 B/op  14046 allocs/op
modl
                   Insert:   2000     7.23s      3615847 ns/op    1688 B/op     43 allocs/op
       BulkInsert 100 row:    500     Don't support bulk insert
                   Update:   2000     7.92s      3961305 ns/op    1296 B/op     40 allocs/op
                     Read:   4000     3.00s       749301 ns/op    1776 B/op     41 allocs/op
     MultiRead limit 1000:   2000     7.59s      3796295 ns/op  514018 B/op  16675 allocs/op

Reports:

  2000 times - Insert
       dbr:     6.36s      3178876 ns/op    2985 B/op     74 allocs/op
        pg:     6.86s      3429856 ns/op    1559 B/op     11 allocs/op
       rel:     6.95s      3477098 ns/op    2447 B/op     49 allocs/op
      xorm:     6.98s      3492309 ns/op    3164 B/op     98 allocs/op
       raw:     7.08s      3538114 ns/op     696 B/op     18 allocs/op
      modl:     7.23s      3615847 ns/op    1688 B/op     43 allocs/op
      gorp:     7.26s      3632476 ns/op    1688 B/op     44 allocs/op
       qbs:     7.52s      3757715 ns/op    5681 B/op    123 allocs/op
     beego:     7.53s      3766770 ns/op    2424 B/op     56 allocs/op
      sqlx:     7.58s      3788224 ns/op    2319 B/op     51 allocs/op
      godb:     7.77s      3885032 ns/op    4730 B/op    115 allocs/op
       pop:     8.61s      4305478 ns/op   10335 B/op    247 allocs/op
      gorm:     9.00s      4499248 ns/op    6853 B/op     97 allocs/op
      hood:     9.70s      4850088 ns/op    7088 B/op    173 allocs/op
     upper:     9.94s      4969764 ns/op   27734 B/op   1184 allocs/op
    genmai:    10.05s      5026864 ns/op    4502 B/op    148 allocs/op

   500 times - BulkInsert 100 row
       dbr:     0.02s        31893 ns/op    2081 B/op     39 allocs/op
     beego:     2.91s      5828279 ns/op  196637 B/op   2845 allocs/op
        pg:     3.19s      6378572 ns/op   14913 B/op    214 allocs/op
       rel:     3.28s      6566187 ns/op  287076 B/op   4053 allocs/op
    genmai:     3.30s      6598140 ns/op  205003 B/op   3066 allocs/op
      godb:     3.40s      6791558 ns/op  289700 B/op   5994 allocs/op
      xorm:     3.63s      7263580 ns/op  319965 B/op   7542 allocs/op
     upper:     4.01s      8027425 ns/op  482063 B/op  19820 allocs/op
       raw:     0.00s      1.20 ns/op       0 B/op      0 allocs/op
      gorm:     Don't support bulk insert - https://github.com/jinzhu/gorm/issues/255
      hood:     Problematic bulk insert, too slow
       pop:     Problematic bulk insert, too slow
      sqlx:     benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134
       qbs:     Don't support bulk insert, err driver: bad connection
      gorp:     Problematic bulk insert, too slow
      modl:     Don't support bulk insert

  2000 times - Update
       raw:     0.32s       158640 ns/op     712 B/op     19 allocs/op
       dbr:     0.39s       195959 ns/op    2619 B/op     57 allocs/op
       pop:     6.52s      3258479 ns/op    6795 B/op    197 allocs/op
      gorp:     6.72s      3361338 ns/op    1344 B/op     39 allocs/op
        pg:     7.00s      3500675 ns/op     992 B/op     13 allocs/op
       rel:     7.15s      3575988 ns/op    2608 B/op     50 allocs/op
      sqlx:     7.18s      3590337 ns/op    1016 B/op     21 allocs/op
      godb:     7.33s      3666445 ns/op    5377 B/op    154 allocs/op
     beego:     7.62s      3811246 ns/op    1801 B/op     47 allocs/op
      xorm:     7.77s      3883802 ns/op    3217 B/op    126 allocs/op
      modl:     7.92s      3961305 ns/op    1296 B/op     40 allocs/op
       qbs:     8.84s      4420961 ns/op    5898 B/op    149 allocs/op
    genmai:     9.23s      4615697 ns/op    3521 B/op    146 allocs/op
      hood:     9.35s      4673388 ns/op   13481 B/op    324 allocs/op
     upper:     9.98s      4989815 ns/op   33009 B/op   1491 allocs/op
      gorm:    13.96s      6978174 ns/op    7556 B/op     93 allocs/op

  4000 times - Read
       raw:     0.65s       163536 ns/op     888 B/op     24 allocs/op
       pop:     0.76s       189248 ns/op    3669 B/op     72 allocs/op
       rel:     0.80s       200248 ns/op    1616 B/op     44 allocs/op
      sqlx:     0.83s       208223 ns/op    1744 B/op     38 allocs/op
      gorp:     0.85s       212743 ns/op    3952 B/op    188 allocs/op
        pg:     0.87s       216339 ns/op    1262 B/op     14 allocs/op
       dbr:     1.10s       274037 ns/op    2176 B/op     36 allocs/op
     beego:     1.45s       362332 ns/op    2113 B/op     75 allocs/op
      gorm:     1.55s       386842 ns/op    4612 B/op     93 allocs/op
      godb:     1.57s       391860 ns/op    4192 B/op    102 allocs/op
      xorm:     1.75s       438095 ns/op    8795 B/op    252 allocs/op
    genmai:     1.89s       472793 ns/op    3313 B/op    171 allocs/op
     upper:     1.90s       474667 ns/op    7385 B/op    293 allocs/op
      modl:     3.00s       749301 ns/op    1776 B/op     41 allocs/op
       qbs:     reflect: call of reflect.Value.Bytes on string Value
      hood:     reflect: call of reflect.Value.Bytes on string Value

  2000 times - MultiRead limit 1000
        pg:     4.57s      2284203 ns/op  319984 B/op   5027 allocs/op
       raw:     4.63s      2316797 ns/op  272016 B/op  11657 allocs/op
      gorp:     6.05s      3026069 ns/op  736514 B/op  15861 allocs/op
       dbr:     6.75s      3373897 ns/op  514960 B/op  16705 allocs/op
      sqlx:     7.00s      3501095 ns/op  499424 B/op  13691 allocs/op
    genmai:     7.28s      3640718 ns/op  420674 B/op  12845 allocs/op
      modl:     7.59s      3796295 ns/op  514018 B/op  16675 allocs/op
     upper:     7.90s      3949170 ns/op  647963 B/op  14046 allocs/op
      godb:     9.53s      4764432 ns/op  997680 B/op  31738 allocs/op
       rel:     9.57s      4786094 ns/op 1010636 B/op  24674 allocs/op
       pop:    10.09s      5045536 ns/op  690531 B/op  14754 allocs/op
     beego:    12.24s      6118856 ns/op  746817 B/op  32474 allocs/op
      gorm:    12.56s      6279390 ns/op  876112 B/op  36740 allocs/op
      xorm:    22.47s     11233651 ns/op 1447692 B/op  55858 allocs/op
```
