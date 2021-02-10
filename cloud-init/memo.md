# cloud-init

## example

```
koketani: ~/g/g/k/p/cloud-init (main ?)$ echo $VAGRANT_DEFAULT_PROVIDER
virtualbox
koketani: ~/g/g/k/p/cloud-init (main ?)$ echo $VAGRANT_EXPERIMENTAL
cloud_init,disks

koketani: ~/g/g/k/p/cloud-init (main ?)$ vagrant up
==> vagrant: You have requested to enabled the experimental flag with the following features:
==> vagrant:
==> vagrant: Features:  cloud_init, disks
==> vagrant:
==> vagrant: Please use with caution, as some of the features may not be fully
==> vagrant: functional yet.
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'focal-server-cloudimg-amd64-vagrant'...
==> default: Matching MAC address for NAT networking...
==> default: Setting the name of the VM: ubuntu
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Preparing user data for cloud-init...
==> default: Configuring storage mediums...
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default:
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default:
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Waiting for cloud init to finish running
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /vagrant => /Users/y-tsuji/git/github.com/koketani/playground-oss/cloud-init
koketani: ~/g/g/k/p/cloud-init (main ?)$ vagrant ssh nginx
==> vagrant: You have requested to enabled the experimental flag with the following features:
==> vagrant:
==> vagrant: Features:  cloud_init, disks
==> vagrant:
==> vagrant: Please use with caution, as some of the features may not be fully
==> vagrant: functional yet.
The machine with the name 'nginx' was not found configured for
this Vagrant environment.
koketani: ~/g/g/k/p/cloud-init (main ?)$ vagrant ssh
==> vagrant: You have requested to enabled the experimental flag with the following features:
==> vagrant:
==> vagrant: Features:  cloud_init, disks
==> vagrant:
==> vagrant: Please use with caution, as some of the features may not be fully
==> vagrant: functional yet.
Welcome to Ubuntu 20.04.2 LTS (GNU/Linux 5.4.0-65-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  System information as of Wed Feb 10 12:53:17 UTC 2021

  System load:  0.69              Processes:               125
  Usage of /:   3.6% of 38.71GB   Users logged in:         0
  Memory usage: 10%               IPv4 address for enp0s3: 10.0.2.15
  Swap usage:   0%


4 updates can be installed immediately.
3 of these updates are security updates.
To see these additional updates run: apt list --upgradable


vagrant@ubuntu:~$ nginx
nginx: [alert] could not open error log file: open() "/var/log/nginx/error.log" failed (13: Permission denied)
2021/02/10 12:53:20 [warn] 2797#2797: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /etc/nginx/nginx.conf:1
2021/02/10 12:53:20 [emerg] 2797#2797: open() "/var/log/nginx/access.log" failed (13: Permission denied)
```
