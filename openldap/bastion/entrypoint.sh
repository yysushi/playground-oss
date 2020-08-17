#!/bin/bash

set -e
mkdir -p /run/sshd

mkdir -p /etc/ssh/keys
if [ ! -f /etc/ssh/keys/ssh_host_ed25519_key ]; then
    ssh-keygen -A
    mv /etc/ssh/ssh_host_* /etc/ssh/keys/
fi

cat >> /etc/ssh/sshd_config <<EOT
HostKey /etc/ssh/keys/ssh_host_rsa_key
HostKey /etc/ssh/keys/ssh_host_dsa_key
HostKey /etc/ssh/keys/ssh_host_ecdsa_key
HostKey /etc/ssh/keys/ssh_host_ed25519_key
EOT

sed -i -e 's/#PermitRootLogin prohibit-password/PermitRootLogin no/g' /etc/ssh/sshd_config
sed -i -e 's~#AuthorizedKeysCommand none~AuthorizedKeysCommand /usr/local/bin/ssh-ldap-pubkey-wrapper~g' /etc/ssh/sshd_config
sed -i -e 's~#AuthorizedKeysCommandUser nobody~AuthorizedKeysCommandUser nobody~g' /etc/ssh/sshd_config
sed -i -e 's~UsePAM yes~UsePAM yes~g' /etc/ssh/sshd_config
sed -i -e 's~ChallengeResponseAuthentication no~ChallengeResponseAuthentication yes~g' /etc/ssh/sshd_config
echo 'AuthenticationMethods publickey,keyboard-interactive:pam' >> /etc/ssh/sshd_config
exec "$@"
