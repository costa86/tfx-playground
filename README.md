
# list workspace
```sh 
tfx workspace list --json | jq .
```
```sh 
tfx workspace list --json | jq '. [0]'
```

# plan
```sh 
tfx plan -w api-test -d ./workspace
```

## speculate
```sh 
tfx plan -w api-test -d ./workspace --speculative
```

# apply
```sh 
tfx apply -i {run id}
```

# locking
```sh 
tfx workspace lock all
```
```sh 
tfx workspace unlock all
```

# variables
```sh 
tfx workspace variable list -w api-test
```

```sh 
tfx ws var create -w api-test -k var4 -v "value1" -d "some important"
tfx ws var create -w api-test -k var5 -v "value2" -d "some very important" --sensitive
tfx ws var create -w api-test -k var6 -v "value3" -d "some env" --env
```
## from file
```sh 
tfx ws var create -w api-test -k var7 -v ./variables/some_map.hcl -d "from a file" --hcl
```
## delete
```sh 
tfx ws var delete -w api-test -k var4 
tfx ws var delete -w api-test -k var5
tfx ws var delete -w api-test -k var7 
```
# runs
```sh 
tfx workspace run list -w api-test
```
## show
```sh 
tfx workspace run show -i {id}
```

## download
```sh 
tfx workspace cv download -i {cv} -d ./
```

# module

## create
```sh 
tfx registry module create -n tls-private -p tls
```

## list
```sh 
tfx registry module version list -n tls-private -p tls
```

## create version
```sh 
tfx registry module version create -n tls-private -p tls -d ./workspace -v 0.0.1
```

## update version
```sh 
tfx registry module version create -n tls-private -p tls -d ./workspace -v 0.0.2
```

# releases
```sh 
tfx release tfe list
```