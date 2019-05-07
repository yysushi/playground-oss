#!/bin/bash
set -e

# configure database for heat
./wait-for-it.sh $MYSQL_HOST:$MYSQL_PORT -- echo "mysql is up"
cat << EOF > /tmp/heat.sql
CREATE DATABASE IF NOT EXISTS heat;
GRANT ALL PRIVILEGES ON heat.* TO 'heat'@'localhost' IDENTIFIED BY '$HEAT_DBPASS';
GRANT ALL PRIVILEGES ON heat.* TO 'heat'@'%' IDENTIFIED BY '$HEAT_DBPASS';
EOF
mysql -uroot -h$MYSQL_HOST -P$MYSQL_PORT -p$MYSQL_ROOT_PASSWORD < /tmp/heat.sql

# change keystone configuration
cat << EOF > /etc/heat/heat.conf
[database]
connection = mysql://heat:$HEAT_DBPASS@$MYSQL_HOST:$MYSQL_PORT/heat

[DEFAULT]
verbose = True

rpc_backend = rabbit
rabbit_host = $RABBIT_HOST
rabbit_password = $RABBITMQ_DEFAULT_PASS
heat_metadata_server_url = http://$(hostname):8000
heat_waitcondition_server_url = http://$(hostname):8000/v1/waitcondition

[keystone_authtoken]
auth_uri = http://$KEYSTONE_HOST:5000/v2.0
identity_uri = http://$KEYSTONE_HOST:35357
admin_tenant_name = service
admin_user = heat
admin_password = $OS_ADMIN_PASSWORD
EOF
su -s /bin/sh -c "heat-manage db_sync" heat

exec "$@"
