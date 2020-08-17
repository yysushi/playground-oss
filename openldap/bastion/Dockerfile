FROM ubuntu:20.04

# packages
ENV DEBIAN_FRONTEND noninteractive
COPY debconf /debconf
RUN debconf-set-selections /debconf
RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    build-essential \
    python3 \
    python3-dev \
    python3-pip \
    libssl-dev \
    libffi-dev \
    libxml2-dev libxslt1-dev \
    libsasl2-dev libldap2-dev \
    libnss-ldap \
    libpam-ldap \
    ldap-utils \
    vim \
    openssh-server \
    libpam-google-authenticator \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN pip3 install ssh-ldap-pubkey

# ldap client and nss
COPY ldap.conf /etc/ldap/
COPY nsswitch.conf /etc/

COPY pam_sshd /etc/pam.d/sshd
RUN mkdir /google_authenticators

# ssh server
EXPOSE 22
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
CMD ["/usr/sbin/sshd", "-D", "-e", "-f", "/etc/ssh/sshd_config"]
