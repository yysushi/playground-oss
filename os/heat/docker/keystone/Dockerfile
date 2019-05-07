FROM ubuntu:14.04

ENV DEBIAN_FRONTEND	noninteractive

RUN apt-get update -y \
 && apt-get dist-upgrade -y \
 && apt-get install -y software-properties-common \
 && add-apt-repository cloud-archive:juno

ENV MYSQL_HOST			mysql
ENV MYSQL_PORT			3306
ENV MYSQL_ROOT_PASSWORD		root_pass
ENV KEYSTONE_DBPASS		keystone_pass
ENV ADMIN_TOKEN			admin_token
ENV OS_ADMIN_PASSWORD		admin_pass
ENV OS_DEMO_PASSWORD		demo_pass

RUN apt-get update \
 && apt-get install -y python-pip python-dev libssl-dev \
 && apt-get install -y mysql-client python-mysqldb \
 && apt-get install -y keystone python-keystoneclient python-openstackclient \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./
RUN chmod +x wait-for-it.sh
COPY ./docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD /usr/bin/keystone-all
