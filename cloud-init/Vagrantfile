# -*- mode: ruby -*-
# vi: set ft=ruby :

# frozen_string_literal: true

specs = {
  memory: 2048,
  cpus: 2
}

Vagrant.configure('2') do |config|
  config.vm.box = 'focal-server-cloudimg-amd64-vagrant'
  config.vm.box_url = 'https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64-vagrant.box'
  config.vm.provider 'virtualbox' do |vb|
    vb.name = 'ubuntu'
    vb.cpus = specs[:cpus]
    vb.memory = specs[:memory]
  end
  config.vm.cloud_init do |cloud_init|
    cloud_init.content_type = 'text/cloud-config'
    cloud_init.inline = <<-PKG
      package_update: true
      packages:
        - nginx
    PKG
  end
end
