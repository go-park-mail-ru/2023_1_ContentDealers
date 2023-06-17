#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username postgres --dbname postgres <<-EOSQL
	CREATE DATABASE filmium;
EOSQL
psql -U postgres filmium -f /tmp/all.sql




# TYPE | DATABASE | USER | ADDRESS | METHOD
echo  "local all postgres md5" > /var/lib/postgresql/data/pg_hba.conf
echo  "host filmium user_schema_content 192.168.243.11/24 md5" >> /var/lib/postgresql/data/pg_hba.conf
echo  "host filmium user_schema_user 192.168.243.12/24 md5" >> /var/lib/postgresql/data/pg_hba.conf
echo  "host filmium user_schema_action 192.168.243.13/24 md5" >> /var/lib/postgresql/data/pg_hba.conf

# TYPE 		DATABASE 	USER 					ADDRESS 			METHOD
# local	 	all 		postgres 									md5
# host 		filmium 	user_schema_content 	192.168.243.11/24 	md5
# host 		filmium 	user_schema_user 		192.168.243.12/24 	md5
# host 		filmium 	user_schema_action 		192.168.243.13/24 	md5
