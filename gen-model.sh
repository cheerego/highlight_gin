#!/bin/bash


table=$1
echo "table name: ${table}"
table | awk
exit
db2struct \
--json \
--gorm \
--guregu \
--package mysql \
--host localhost \
--mysql_port 3306 \
--user root \
-p root \
-d haha \
-t $1 \
--struct User \
| sed -e "s/sql.NullFloat64/null.Float     /g" \
| sed -e "s/sql.NullInt64/null.Int     /g" \
| sed -e "s/sql.NullString/null.String   /g" \
| sed -e "s/time.Time/null.Time/g" \
> ./model/user.go#!/usr/bin/env bash