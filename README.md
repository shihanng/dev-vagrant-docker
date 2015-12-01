## Run

Setup proxy with [vagrant-proxyconf](http://tmatilai.github.io/vagrant-proxyconf/) if necessary.
```shell
$ export VAGRANT_HTTP_PROXY="http://proxy.local.hde.co.jp:9080"
$ vagrant up --no-parallel
```

## Test

```shell
$ vagrant docker-run hello -- go test hello
```
