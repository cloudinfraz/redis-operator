# Redis Operator

**How to manage Redis configuration **

Redis configuration is very important for the different scenario, the redisConfig abstract will help us to handle the different configurations

- global redisConfig

  The global redisConfig is only adopted by Redis standalone mode, I think it's a bit of a redundancy if the global redisConfig will be 
  specify for Redis cluster, the master and slave redisConfig already can be specified seperately

- master redisConfig

  The master redisConfig will be adopted by Redis cluster mode for the master nodes, the initial redisConfig will enabled if the master or slave 
  redisConfig will not be specified

- slave redisConfig

  The slave redisConfig will be adopted by Redis cluster mode for the slave nodes, the initial redisConfig will enabled if the master and slave
  redisConfig will not be specified

- initial redisConfig 

  the initial redisConfig is not very clear defined, which can be elaborated

redisConfig crd data type
```
RedisConfig       map[string][]string        `json:"redisConfig"` \\the duplicated keys for some parameters

# code
masterRediscfg := &cr.Spec.Master.RedisConfig
slaveRediscfg := &cr.Spec.Slave.RedisConfig
standaloneRediscfg := &cr.Spec.RedisConfig

```
**How to implement to change Redis configuration **

Currently, only the static configuration can't be deployed, but the dynamic parameters will be implemented soon, there are two approach for the static configuration

- init container
Init containers can contain redis config file and will copy to the reids container ,then start the standalone/cluster of redis

- the entrypoint.sh shell script of redis container
the entrypoint.sh will read the redis config file,then append the initial redis config

**redisConfig crd definition **

In Redis cluster mode

```
master:
  service:
    type: NodePort
  redisConfig:
    save:
      - "900 2"
      - "300 10"
      - "60 10000"
    timeout:
      - "5"
```
In Redis standalone mode

```
redisConfig: 
  timeout:
    - "2"
```

<div align="center">
    <img src="../../../static/configmap.png">
</div>
