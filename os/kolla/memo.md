# Kolla

## memo

<https://docs.openstack.org/kolla/latest/admin/image-building.html#building-kolla-images>

```shell-sesion
koketani:kolla (master=)$ ls etc/kolla/
koketani:kolla (master=)$ tox -e genconfig
genconfig create: /Users/koketani/Developments/git/github.com/koketani/playground-oss/os/kolla/kolla/.tox/genconfig
genconfig installdeps: -chttps://releases.openstack.org/constraints/upper/master, -r/Users/koketani/Developments/git/github.com/koketani/playground-oss/os/kolla/kolla/requirements.txt, -r/Users/koketani/Developments/git/github.com/koketani/playground-oss/os/kolla/kolla/test-requirements.txt
genconfig develop-inst: /Users/koketani/Developments/git/github.com/koketani/playground-oss/os/kolla/kolla
genconfig installed: appdirs==1.4.4,Babel==2.8.0,bandit==1.6.2,bashate==2.0.0,beautifulsoup4==4.9.1,certifi==2020.4.5.1,cffi==1.14.0,chardet==3.0.4,cliff==3.1.0,cmd2==0.8.9,coverage==5.1,cryptography==2.9.2,ddt==1.4.1,debtcollector==2.1.0,decorator==4.4.2,docker==4.2.1,dogpile.cache==0.9.2,entrypoints==0.3,extras==1.0.0,fixtures==3.0.0,flake8==3.7.9,future==0.18.2,gitdb==4.0.5,GitPython==3.1.3,graphviz==0.14,hacking==3.0.1,idna==2.9,iso8601==0.1.12,Jinja2==2.11.2,jmespath==0.10.0,jsonpatch==1.25,jsonpointer==2.0,keystoneauth1==4.0.0,-e git+https://github.com/openstack/kolla@778d0339c4127dd2ca55a4cac75948083bc94141#egg=kolla&subdirectory=../../../../os/kolla/kolla,linecache2==1.0.0,MarkupSafe==1.1.1,mccabe==0.6.1,mock==3.0.5,msgpack==0.6.1,munch==2.5.0,netaddr==0.7.19,netifaces==0.10.9,openstacksdk==0.46.0,os-client-config==2.1.0,os-service-types==1.7.0,osc-lib==2.1.0,oslo.config==8.1.0,oslo.context==3.1.0,oslo.i18n==5.0.0,oslo.log==4.2.1,oslo.serialization==3.2.0,oslo.utils==4.2.0,oslotest==4.3.0,pbr==5.4.5,prettytable==0.7.2,pycodestyle==2.5.0,pycparser==2.20,pyflakes==2.1.1,pyparsing==2.4.7,pyperclip==1.8.0,python-barbicanclient==4.10.0,python-cinderclient==7.0.0,python-dateutil==2.8.1,python-heatclient==2.1.0,python-keystoneclient==4.0.0,python-mimeparse==1.6.0,python-neutronclient==7.1.1,python-novaclient==17.0.0,python-openstackclient==5.2.0,python-subunit==1.4.0,python-swiftclient==3.9.0,pytz==2020.1,PyYAML==5.3.1,requests==2.23.0,requestsexceptions==1.4.0,rfc3986==1.4.0,simplejson==3.17.0,six==1.15.0,smmap==3.0.4,soupsieve==2.0.1,stestr==3.0.1,stevedore==2.0.0,testscenarios==0.5.0,testtools==2.4.0,traceback2==1.4.0,unittest2==1.1.0,urllib3==1.25.9,voluptuous==0.11.7,wcwidth==0.2.3,websocket-client==0.57.0,wrapt==1.12.1
genconfig run-test-pre: PYTHONHASHSEED='1102031089'
genconfig run-test: commands[0] | oslo-config-generator --config-file etc/oslo-config-generator/kolla-build.conf
/Users/koketani/Developments/git/github.com/koketani/playground-oss/os/kolla/kolla/.tox/genconfig/lib/python3.8/site-packages/oslo_config/types.py:57: UserWarning: converting '[]' to a string
  warnings.warn('converting \'%s\' to a string' % str_val)
______________________________________________________ summary _______________________________________________________
  genconfig: commands succeeded
  congratulations :)
koketani:kolla (master=)$ ls etc/kolla/
kolla-build.conf

$ python setup.py install

$ koketani:kolla (master=)$ kolla-build tacker
INFO:kolla.common.utils:Found the docker image folder at /Users/koketani/Developments/git/github.com/koketani/playground-oss/os/kolla/.venv/share/kolla/docker
INFO:kolla.common.utils:Added image base to queue
INFO:kolla.common.utils:Attempt number: 1 to run task: BuildTask(base)
INFO:kolla.common.utils.base:Building started at 2020-06-06 15:45:36.113890
INFO:kolla.common.utils.base:[WARNING]: Empty continuation line found in:

...

koketani:kolla (master=)$ docker images  | grep tacker
kolla/centos-binary-tacker-conductor                10.1.0              13523d038da6        6 seconds ago        927MB
kolla/centos-binary-tacker-server                   10.1.0              a712421746b6        12 seconds ago       927MB
kolla/centos-binary-tacker-base                     10.1.0              e29367e34d2b        About a minute ago   893MB
```
