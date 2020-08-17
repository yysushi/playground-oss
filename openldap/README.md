# Bastion server

## Deploy

1. Create LDAP admin/config password as docker secret.

```shell-session
$ printf "<admin_pw>" | docker secret create authentication_admin_pw -
nx7gioomcshqw4j58b4ieywrq
$ printf "<config_pw>" | docker secret create authentication_config_pw -
itiqkmn0b625n5jpd6qmz7iux
```

2. Build images and deploy stack

```shell-skoketanison
$ docker-compose build
$ docker stack deploy openldap --compose-file docker-compose.yml
```

3. Distribute CA certificate to client

```shell-session
$ docker cp openldap_openldap.1.0tfwk0n8bcgmkv6hsj4x6wqiv:/container/service/\:ssl-tools/assets/default-ca/default-ca.pem bastion/certs/ca.crt
```

## Add User

1. Append LDAP user entry

```shell-session
$ cat openldap/custom/99_users.ldif
dn: uid=koketani,ou=People,dc=example,dc=org
objectClass: top
objectClass: person
objectClass: posixAccount
objectClass: shadowAccount
objectClass: ldapPublicKey
uid: koketani
cn: Koke Tani
sn: Tani
loginShell: /sbin/nologin
uidNumber: 1000
gidNumber: 1000
homeDirectory: /home/koketani
description: This is an example user
sshPublicKey: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIC
 8utV7RcMz3Kja9xSqj++VZnzdALWADbjnpUaUlNqQI
```

2. Create Google Authenticator secret

```shell-session
$ docker run -v $(pwd)/bastion/google_authenticators:/google_authenticators --entrypoint bash -it koketani/bastion -c "google-authenticator -t -d -f -Q ANSI -r 3 -R 30 -s /google_authenticators/koketani -w 3"
$ sudo chown koketani:koketani bastion/google_authenticators/koketani
```

Enter this secret key in mobile's Google Authenticator app.

3. Submit new LDAP user entry

```shell-session
$ docker service update openldap_openldap
```

## Remove User

1. Remove LDAP user entry

```shell-session
$  cat openldap/custom/99_users.ldif
```

2. Remove Google Authenticator file

```shell-session
$ rm bastion/google_authenticators/koketani
```

3. Update docker services in the stack

```shell-session
$ docker service update openldap_openldap
```

## Commands Reference

- List up users by using ldap server

```shell-session
root@bastion:/# getent passwd | tail -n1
koketani:x:1000:1000:Koke Tani:/home/koketani:/sbin/nologin
```

- Get ssh public key stored in ldap server

```shell-session
root@bastion:/# ssh-ldap-pubkey-wrapper koketani
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIC8utV7RcMz3Kja9xSqj++VZnzdALWADbjnpUaUlNqQI
```

- Search LDAP entries with no filters by readonly user

```shell-session
root@bastion:/# ldapsearch -b dc=example,dc=org -D "cn=readonly,dc=example,dc=org" -W | head -n5
Enter LDAP Password:
# extended LDIF
#
# LDAPv3
# base <dc=example,dc=org> with scope subtree
# filter: (objectclass=*)
root@bastion:/# ldapsearch -b dc=example,dc=org -D "cn=readonly,dc=example,dc=org" -W | head -n10
Enter LDAP Password:
# extended LDIF
#
# LDAPv3
# base <dc=example,dc=org> with scope subtree
# filter: (objectclass=*)
# requesting: ALL
#

# example.org
dn: dc=example,dc=org
```

- Search LDAP entries with specifying host (non-SSL) and debug level

```shell-session
root@bastion:/# ldapsearch -d255 -H ldap://ldap-server -b dc=example,dc=org -D "cn=readonly,dc=example,dc=org" -W
```

- Connect lab A koketani by using bation as proxy

```shell-skoketanison
ssh -o ProxyCommand="ssh -W %h:%p 172.23.16.118 -p 2222 -l koketani -i id_rsa" koketani@internal
```
