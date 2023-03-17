#!/bin/bash

d=$(cd $(dirname $0)/..; pwd)
mysql -h"$MYSQL_HOST" -u"$MYSQL_USER" -p"$MYSQL_PASS" "$MYSQL_DATABASE" --connect-timeout=3 --protocol=tcp < "${d}/mysql/seed.sql"
