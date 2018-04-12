# go-todo

## create mock
```
# install
$ go get github.com/golang/mock/gomock
$ go get github.com/golang/mock/mockgen

```

```
# generate mock
$ mockgen -source sample.go -destination sample_mock.go
```

## Test
```
make test
```
