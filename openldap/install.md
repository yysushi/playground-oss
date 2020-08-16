# OpenLDAP

## Links

<https://www.openldap.org/doc/admin24/quickstart.html>

## Environment

please renew secrets and keys

- create secrets

```shell-session
$ docker secret create authentication_admin_pw ./openldap/pw/admin_pw.txt
$ docker secret create authentication_config_pw ./openldap/pw/config_pw.txt
$ docker secret ls
ID                          NAME                       DRIVER              CREATED             UPDATED
ipkkdl68y4yiq46ssarhk78n3   authentication_admin_pw                        13 minutes ago      13 minutes ago
m0qpkuv6uwr5spieid8vqq1ki   authentication_config_pw                       12 minutes ago      12 minutes ago
```

- run openldap servers

```shell-session
$ docker stack deploy openldap --compose-file docker-compose.yml
Ignoring unsupported options: domainname

Ignoring deprecated options:

container_name: Setting the container name is not supported.

Creating network openldap_default
Creating service openldap_openldap
Creating service openldap_phpldapadmin
```

- confirm its running

```shell-session
$ docker exec -it openldap_openldap.1.8yyv70gu1e4h7zlozmzqw7moq ldapsearch -H ldaps://localhost -b dc=example,dc=org -D "cn=admin,dc=example,dc=org" -W
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
objectClass: top
objectClass: dcObject
objectClass: organization
o: Example Inc.
dc: example

# admin, example.org
dn: cn=admin,dc=example,dc=org
objectClass: simpleSecurityObject
objectClass: organizationalRole
cn: admin
description: LDAP administrator
userPassword:: e1NTSEF9a1g5Q2NBa1ZlL2Y2LzBmcndrZk5QeFpPOUFCRDlFcVY=

# search result
search: 2
result: 0 Success

# numResponses: 3
# numEntries: 2
```

- confirm acl configuration
https://www.openldap.org/doc/admin24/access-control.html

```shell-session
root@ldap-server:/# ldapsearch -Y EXTERNAL -H ldapi:/// -b olcDatabase={1}mdb,cn=config
SASL/EXTERNAL authentication started
SASL username: gidNumber=0+uidNumber=0,cn=peercred,cn=external,cn=auth
SASL SSF: 0
# extended LDIF
#
# LDAPv3
# base <olcDatabase={1}mdb,cn=config> with scope subtree
# filter: (objectclass=*)
# requesting: ALL
#

# {1}mdb, config
dn: olcDatabase={1}mdb,cn=config
objectClass: olcDatabaseConfig
objectClass: olcMdbConfig
olcDatabase: {1}mdb
olcDbDirectory: /var/lib/ldap
olcSuffix: dc=example,dc=org
olcAccess: {0}to * by dn.exact=gidNumber=0+uidNumber=0,cn=peercred,cn=external
 ,cn=auth manage by * break
olcAccess: {1}to attrs=userPassword,shadowLastChange by self write by dn="cn=a
 dmin,dc=example,dc=org" write by anonymous auth by * none
olcAccess: {2}to * by self read by dn="cn=admin,dc=example,dc=org" write by *
 none
olcLastMod: TRUE
olcRootDN: cn=admin,dc=example,dc=org
olcRootPW: {SSHA}kX9CcAkVe/f6/0frwkfNPxZO9ABD9EqV
olcDbCheckpoint: 512 30
olcDbIndex: uid eq
olcDbIndex: mail eq
olcDbIndex: memberOf eq
olcDbIndex: entryCSN eq
olcDbIndex: entryUUID eq
olcDbIndex: objectClass eq
olcDbMaxSize: 1073741824

# {0}memberof, {1}mdb, config
dn: olcOverlay={0}memberof,olcDatabase={1}mdb,cn=config
objectClass: olcOverlayConfig
objectClass: olcMemberOf
olcOverlay: {0}memberof
olcMemberOfDangling: ignore
olcMemberOfRefInt: TRUE
olcMemberOfGroupOC: groupOfUniqueNames
olcMemberOfMemberAD: uniqueMember
olcMemberOfMemberOfAD: memberOf

# {1}refint, {1}mdb, config
dn: olcOverlay={1}refint,olcDatabase={1}mdb,cn=config
objectClass: olcOverlayConfig
objectClass: olcRefintConfig
olcOverlay: {1}refint
olcRefintAttribute: owner
olcRefintAttribute: manager
olcRefintAttribute: uniqueMember
olcRefintAttribute: member
olcRefintAttribute: memberOf

# search result
search: 2
result: 0 Success

# numResponses: 4
# numEntries: 3
```

- confirm all config

`ldapsearch -Y EXTERNAL -H ldapi:/// -b cn=config`
