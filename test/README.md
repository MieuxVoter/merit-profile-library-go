The files in this directory are only used during development and none of them are included in the distributed library.

In Golang, this means that none of the variables, types, etc. in this directory start with an uppercase letter.

Run this from project root:

```shell
max=6
for i in `seq 1 $max`
do
    echo "Running test${i}.go to generate test${i}.svg"
    go run test/test${i}.go > test/test${i}.svg
done
```
