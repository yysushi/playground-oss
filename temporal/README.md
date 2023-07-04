# [Temporal](https://temporal.io/)

## References

### GitHub

- [temporalite](https://github.com/temporalio/temporalite)
- [temporalio/samples-go](https://github.com/temporalio/samples-go)
- [temporal-cli](https://github.com/temporalio/tctl)

- [temporal with docker compose](https://github.com/temporalio/docker-compose)

- [proto files](https://github.com/temporalio/api)

### Nix

- [temporal-cli](https://github.com/NixOS/nixpkgs/blob/nixos-23.05/pkgs/applications/networking/cluster/temporal-cli/default.nix#L79), [temporalite](https://github.com/NixOS/nixpkgs/blob/nixos-23.05/pkgs/applications/networking/cluster/temporalite/default.nix#L30)

### Official Documents

- [concepts](https://docs.temporal.io/concepts/)
- [deployments](https://docs.temporal.io/cluster-deployment-guide)

## CLI

- run server: `temporalite start --namespace default`
- run server: `tctl namespace list`

## Sample: Branch
## Sample: xxx

## Investigation Items

- Core idea on the workflow engine
- DSL
- Catch up with the latest approach of binding to SQL database ... sqlx?

### SQL appraoch

```
y-tsuji: ~/g/g/t/temporal (master)$ grep "type sqlExecutionStore" common/persistence/sql/execution.go
type sqlExecutionStore struct {
y-tsuji: ~/g/g/t/temporal (master)$ grep "CREATE TABLE executions" schema/mysql/v8/temporal/schema.sql
CREATE TABLE executions(
```

https://github.com/temporalio/temporal/pull/4346



https://github.com/temporalio/temporal/blob/master/CONTRIBUTING.md

make install-schema

https://github.com/temporalio/temporal/blob/master/cmd/tools/sql/main.go

tools/sql/README.md

https://github.com/temporalio/temporal/blob/master/develop/docs/new-rpcs.md#adding-new-rpcs
