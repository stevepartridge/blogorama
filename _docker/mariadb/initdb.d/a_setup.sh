#!/bin/bash -e

#
# temporary until a specific container can be 
# configured to use this as part of start up
#

# echo "Run"
# docker-entrypoint.sh
# echo "Ran"

echo "Root: ${MYSQL_ROOT_PASSWORD}"
echo "User: ${MYSQL_USER}"
echo "Pass: ${MYSQL_PASS}"

# service mysql start || true

TEMP_FILE='/tmp/mariadb-start.sql'

touch $TEMP_FILE

IFS=',' read -r -a dbnames <<< "$CREATE_DATABASES"
for dbname in "${dbnames[@]}"
do
    echo "CREATE DATABASE IF NOT EXISTS $dbname ;" >> "$TEMP_FILE"
    echo "GRANT ALL PRIVILEGES ON $dbname.* TO '$MYSQL_USER'@'%' IDENTIFIED BY '$MYSQL_PASS' WITH GRANT OPTION ;" >> "$TEMP_FILE"
done

echo "FLUSH PRIVILEGES ;" >> "$TEMP_FILE"

mysql -u root -p$MYSQL_ROOT_PASSWORD < $TEMP_FILE
# service mysql stop

# mysqld

# docker-entrypoint.sh