# Redis Operator

**How to manage Redis configuration **

Redis configuration is very important for the different scenario, the redisConfig abstract will help us to handle the different configurations

- global redisConfig
  The global redisConfig is only adopted by Redis standalone mode, I think it's too compliated to handle the global redisConfig for Redis cluster,
  the master and slave redisConfig can be specified seperately

- master redisConfig
  The master redisConfig will be adopted by Redis cluster mode for the master nodes, the initial redisConfig will enabled if the master or slave 
  redisConfig will't be specified

- slave redisConfig
  The slave redisConfig will be adopted by Redis cluster mode for the slave nodes, the initial redisConfig will enabled if the master and slave
  redisConfig will't be specified

- initial redisConfig is not very clear defined, which can be elaborated

redisConfig crd definition

```
RedisConfig       map[string][]string        `json:"redisConfig"` \\the duplicated keys for some parameters
masterRediscfg := &cr.Spec.Master.RedisConfig
slaveRediscfg := &cr.Spec.Slave.RedisConfig
standaloneRediscfg := &cr.Spec.RedisConfig

```
<div align="center">
    <img src="../../static/configmap.png">
</div>
