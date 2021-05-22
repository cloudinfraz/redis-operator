# Redis Operator
  This is a Red Hat OLM and Operator learning, development and training github repository, which inspired by the repo below:

[OT-CONTAINER-KIT Redis Operator](https://github.com/OT-CONTAINER-KIT/redis-operator)

- 功能比较

<div align="center">
    <img src="./static/redis-comparsion.png">
</div>

# Roadmap

- [ done ] Support the different redis configuration files by RedisConf crd for the master and slave nodes of redis cluster
- [ ] Support Dynamic redis configuration parameter for redis cluster by RedisConf crd
- [ ] Support tls
- [ ] More feature extension
- [ ] e2e tests

- Reference:

[gitlab-runner](https://gitlab.com/gitlab-org/gitlab-runner/blob/f4645bfbf947b761f69e8ba292bce84e5c95766d/executors/kubernetes/executor_kubernetes.go)
