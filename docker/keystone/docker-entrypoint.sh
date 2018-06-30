#!/bin/bash
set -e

# configure database for keystone
./wait-for-it.sh $MYSQL_HOST:$MYSQL_PORT -- echo "mysql is up"
cat << EOF > /tmp/keystone.sql
CREATE DATABASE IF NOT EXISTS keystone;
GRANT ALL PRIVILEGES ON keystone.* TO 'keystone'@'localhost' IDENTIFIED BY '$KEYSTONE_DBPASS';
GRANT ALL PRIVILEGES ON keystone.* TO 'keystone'@'%' IDENTIFIED BY '$KEYSTONE_DBPASS';
EOF
mysql -uroot -h$MYSQL_HOST -P$MYSQL_PORT -p$MYSQL_ROOT_PASSWORD < /tmp/keystone.sql

# change keystone configuration
cat << EOF > /etc/keystone/keystone.conf
[database]
connection = mysql://keystone:$KEYSTONE_DBPASS@$MYSQL_HOST:$MYSQL_PORT/keystone

[DEFAULT]
admin_token = $ADMIN_TOKEN
verbose = True

[token]
provider = keystone.token.providers.uuid.Provider
driver = keystone.token.persistence.backends.sql.Token

[revoke]
driver = keystone.contrib.revoke.backends.sql.Revoke
EOF
su -s /bin/sh -c "keystone-manage db_sync" keystone

# initialize keystone service
export OS_SERVICE_TOKEN=$ADMIN_TOKEN
export OS_SERVICE_ENDPOINT=http://$(hostname):35357/v2.0
export OS_ADMIN_AUTH_URL=http://$(hostname):35357/v2.0
export OS_DEMO_AUTH_URL=http://$(hostname):5000/v2.0
export OS_ADMIN_TENANT_NAME=admin
export OS_ADMIN_USERNAME=admin
[[ -z "$OS_ADMIN_PASSWORD" ]] && export OS_ADMIN_PASSWORD=admin_pass
export OS_DEMO_TENANT_NAME=demo
export OS_DEMO_USERNAME=demo
[[ -z "$OS_DEMO_PASSWORD" ]] && export OS_DEMO_PASSWORD=demo_pass
cat << EOF > keystone_init.sh
sleep 5
## create admin role/tenant/user
keystone tenant-create --name $OS_ADMIN_TENANT_NAME --description "Admin Tenant"
keystone user-create --name $OS_ADMIN_USERNAME --pass $OS_ADMIN_PASSWORD --email admin@domain.com
keystone role-create --name admin
keystone user-role-add --user $OS_ADMIN_USERNAME --tenant $OS_ADMIN_TENANT_NAME --role admin
## create member tenant/user
keystone tenant-create --name $OS_DEMO_TENANT_NAME --description "Demo Tenant"
keystone user-create --name $OS_DEMO_USERNAME --tenant $OS_DEMO_TENANT_NAME --pass $OS_DEMO_PASSWORD --email demo@domain.com
## create service tenant
keystone tenant-create --name service --description "Service Tenant"
## create service
keystone service-create --name keystone --type identity --description "OpenStack Identity"
## create service endpoint
keystone endpoint-create \
--service-id \$(keystone service-list | awk '/ identity / {print \$2}') \
--publicurl http://$(hostname):5000/v2.0 \
--internalurl http://$(hostname):5000/v2.0 \
--adminurl http://$(hostname):35357/v2.0 \
--region regionOne
EOF
cat << EOF > adminrc
export OS_TENANT_NAME=$OS_ADMIN_TENANT_NAME
export OS_USERNAME=$OS_ADMIN_USERNAME
export OS_PASSWORD=$OS_ADMIN_PASSWORD
export OS_AUTH_URL=$OS_ADMIN_AUTH_URL
EOF
cat << EOF > demorc
export OS_TENANT_NAME=$OS_DEMO_TENANT_NAME
export OS_USERNAME=$OS_DEMO_USERNAME
export OS_PASSWORD=$OS_DEMO_PASSWORD
export OS_AUTH_URL=$OS_DEMO_AUTH_URL
EOF
bash -x keystone_init.sh &
unset OS_SERVICE_TOKEN OS_SERVICE_ENDPOINT

exec "$@"
