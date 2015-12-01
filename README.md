## Run

Setup proxy with [vagrant-proxyconf](http://tmatilai.github.io/vagrant-proxyconf/) if necessary.
```shell
$ export VAGRANT_HTTP_PROXY="http://proxy.local.hde.co.jp:9080"
$ vagrant up --no-parallel
```
Visit http://192.168.33.10:8000.

## Test

```shell
$ vagrant docker-run hello -- go test hello
```
