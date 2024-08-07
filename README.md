
### Terraform Cache Plugin Scaffolding

> This project demonstrates plugin architecture and different resources/datasource inside the plugin.


This Cache Plugin Creates
1. cache cluster ->  return ocid of cache cluster
2. cache_cluster_patching  -> apply patch and return ocid of patch
3. cache_cluster_datasource-> return all clusters

Make sure move .terraformrc file under home directory

```shell
mv .terraformrc ~/.
```

### Create Module in root Dir(If go.mod not present)

```shell

#initial (if go.mod not present)
go mod init github.io/tvajjala/terraform-provider-cache

#check env
go env 

#download dependencies
go get github.io/tvajjala/terraform-provider-cache
# build

make 
```

### Importing sub module fails 

When you change project structure(name/module). submodule import fails. 

run below command from mod directory

```shell
go get github.io/tvajjala/terraform-provider-cache
```
then you are able to import 

`"github.com/hashicorp/terraform-provider-scorpio/internal/provider"`

If above change not works try 
```shell

# add relative path , to edit go.mod that look for module directory
go mod edit -replace github.io/service=./service
```

### Binary file name not changing

check make file

### Update dependency

```shell
go mod why -m google.golang.org/protobuf
```

In Go, `go mod why` shows the shortest path in the import graph from the main module to each package listed.

`go list` may refer to a command that should not unintentionally update minor versions in go.mod.

