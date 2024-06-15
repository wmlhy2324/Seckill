由于刚刚才匆匆忙忙补完，文档和测试还没开始写，感觉做得不是很好，但是还是交了试试，见谅，后面会慢慢完善一下
### 前言

------

用到的一些技术栈

- go-zero

- mysql

- redis

- kafka

- asynq

- etcd

- jaeger

- Prometheus

- docker

### 项目简介

------

- common

  主要有barrier表，gorm的gorm.model结构体，这里把它单独拿出来了，以及修改的返回模板。

- core

  主要包括对mysql，redis，etcd的初始化，后续在依赖注入的时候使用。

- 用lhy开头的一些前缀

  这里其实放进一个目录里面好一点，当时一直写下去了就没有改了主要有

  - lhy_order

  - lhy_auth

  - lhy_cart

  - lhy_gateway

  - lhy_payment

  - lhy_stock

  - lhy_user

    里面都是对应的api和rpc服务

- mqueue

  这是延迟队列，格式按照go-zero生成的文件格式编写，主要用于取消过期订单

- template

  这是go-zero生成api文件的模板，用于修改模板，这里是把请求回复改成了统一的code，data，msg的格式

- utils

  这是一些工具，包括连接redis，获取主机ip（没有使用了），jwt验证，hash密码，从数据库获取指定字段

- main.go

  主要用于生成表结构，由gorm自动生成，更改表结构之后在项目根目录下go run main.go -db即可更新表结构

  ### 服务注册与发现

  ------

  ​    项目使用了etcd进行服务注册与发现

  ![image-20240615235648157](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240615235648157.png)

在api文件生成之后，开启etcd，传入etcd地址和服务名，具体的服务地址就可以进行服务上送

![image-20240616000018258](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616000018258.png)

得到服务地址

### 网关

这里用网关直接用原生得go写的网关，一开始想用nginx，出了一点小问题就改用这个了，主要逻辑就是先匹配请求找到是否有这个服务

![image-20240616000326063](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616000326063.png)

再请求认证服务，让请求都走一遍认证

![image-20240616000418194](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616000418194.png)

最后代理到原请求

### 订单服务

其实我感觉总体内容比较少，主要的就是订单这一块，其他的就比较简单，这里就介绍一下订单服务

- 预热缓存

  首先在每次启动订单服务的的时候，都会从数据库加载秒杀的商品到redis，这里做了一个sql优化，每次限制10条记录

  ![image-20240616000911165](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616000911165.png)

- 分布式事务

  我了解到的分布式数据一致性可以用mq，分布式锁，分布式事务等等来解决，因为在看go-zero的时候看到了dtm，于是索性就用dtm了。

  ![image-20240616001317145](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616001317145.png)

核心代码如图，首先创建一个全局id，再创建了一个saga事务，这里的dtmServer已经通过配置注册到了etcd里面，add添加了两个分支事务，分别对应了先行函数和补偿函数，最后提交事务。

- kafka消息队列

  这里用了kafka做消息队列，主要是防止创建订单时对数据库压力过大，将创建订单的的信息推送给kafka，让消费者慢慢创建订单，由于kafka写消息比较慢，所以用异步执行的方式

  ![image-20240616002458308](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616002458308.png)

消费者

![image-20240616002521525](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616002521525.png)

- 延时队列

  用了asynq来做延时队列，主要用来关闭过期的订单并回调库存，这里也可以用redis做临时订单，需要定时去扫描

![image-20240616003149956](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240616003149956.png)

- 商品缓存

  在创建商品的时候这里会将商品缓存进redis，以减少对数据库的压力
