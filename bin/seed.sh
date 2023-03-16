#!/bin/bash

d=$(cd $(dirname $0)/..; pwd)
mysql -h"$MYSQL_HOST" -u"$MYSQL_USER" -p"$MYSQL_PASS" "$MYSQL_DB" --connect-timeout=3 < "${d}/mysql/seed.sql"
