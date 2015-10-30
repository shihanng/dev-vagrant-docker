# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    config.vm.define "redis" do |r|
        r.vm.provider "docker" do |d| # NOT virtualbox
            d.image = "redis:latest" # Use image, not build one
            d.name = "redis" # docker run --name
            d.ports = ["6379:6379"]
            d.vagrant_vagrantfile = "./docker-host-vm/Vagrantfile" # Host Vagrantfile
        end
    end

    config.vm.define "hello" do |h|
        h.vm.provider "docker" do |d|
            d.build_dir = "./hello" # Build, not using image
            d.build_args = ["-t=go_im"] # docker build --tag
            d.name = "go_con"
            d.vagrant_vagrantfile = "./docker-host-vm/Vagrantfile"
            d.ports = ["8000:8080"]
            d.volumes = ["/src/hello:/go/src/hello/"] # VM-host:Container
            d.cmd = ["go", "run", "/go/src/hello/hello.go"]
            d.link("redis:redis") # Link to redis (see above)
        end
    end
end
