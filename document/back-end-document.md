# Hedy Shop线上购物平台项目说明——后端

## 简介

本项目实现了一个线上购物平台，面向买家和卖家。买家可以在系统中购买和秒杀物品，购买物品会直接返回结果，秒杀则会生成秒杀订单，记录秒杀结果。另外买家还可以在系统中充值。卖家则可以对商品的名称、价格、库存、描述、种类等进行编辑，还可以查看所有的秒杀订单。秒杀机制为当商品的库存余量小于100时购买功能自动变为秒杀功能。卖家和买家无法查看对方的界面，也无法触及对方的功能，前后端都有身份验证相关的机制。另外系统还使用了token实现3天免登陆的功能。

## 项目架构

本项目采用了前后端分离技术，前端使用vue编写，并运行在node.js中。后端使用go编写，后端服务器主要有两个。一个负责接收前端项目的请求，并进行业务处理，包括身份认证，购买商品，卖家管理商品等等。如果用户发送的是秒杀请求，服务器则会将秒杀信息放入rabbitmq中等待另一个后端的处理。而另一个后端就专门负责处理秒杀请求，不断的从rabbitmq队列中取出请求并进行业务处理。数据库使用了mysql。用户的身份认证使用了jwt。

## 后端实现

- 后端代码架构

  ![image-20210111212415590](D:\arts\gopath\src\pku-class\market\document\img\b代码架构.png)

  整个后端使用了gin架构，可以分为消费者和生产者两种模式运行。

  backstage意味后端的后端主要包含rabbitmq消费者的处理，主要是处理秒杀订单。

  data包括一些数据结构的定义、打包和变换。

  database包括一些数据库的配置和初始化等。

  document是文档。

  error-handler是一些封装的错误处理的函数，包含一些安全日志的输出。

  handler是各种业务处理函数。

  jwt是jwt的初始化配置和基本操作函数。

  rabbitmq是消息队列的初始化和配置以及一些基本操作。

  router是路由配置，包括一些拦截器中间件(处理跨域，jwt认证等)。

  下面是main.go的代码：

  ![image-20210111212855884](D:\arts\gopath\src\pku-class\market\document\img\main.go.png)

  可以看到程序可以分为消费者和生产者，这一概念对应于消息队列。生产者模式下，程序会监听8081端口，然后接受前端发来的请求进行处理和返回，如果是秒杀请求就将消息放入消息队列中等待消费者模式下的程序处理。消费者模式下则不断从消息队列中取出秒杀订单来处理。

- router(包含认证)

  ![image-20210111213240851](D:\arts\gopath\src\pku-class\market\document\img\cors.png)

  ![image-20210111213310740](D:\arts\gopath\src\pku-class\market\document\img\JWT.png)

  ![image-20210111213530209](D:\arts\gopath\src\pku-class\market\document\img\b-router.png)

  router主要是配置各种路由，将相应的请求交给相应的handler。另外router使用了两个中间件，一个负责处理跨域请求，主要是因为我们是前后端分离的，前端和后端通信涉及跨域。另一个负责JWT的token解析，身份认证，放行login请求，其余请求则必须在header中包含token，解密后可以得到用户id、用户名和用户角色三个信息，并将这三个信息交给handler(如果handler中涉及到需要用户身份信息的业务，则只能使用jwt解析出的身份信息)。并且token解析还需要输出安全日志。

- handler

  ![image-20210111214049643](D:\arts\gopath\src\pku-class\market\document\img\handler.png)

  handler中主要是各种业务处理，每个目录下都有一个server.go，该文件中主要封装了所有和数据库相关的操作。如果是查询数据库则直接查询，如果涉及增删改则需要使用事务。简单业务都可以直接和数据库交互处理，如果是秒杀请求则将请求放入消息队列。如果是登录请求则要生成新的token返回给前端。

- jwt, rabbitmq, database

  这三个包下分别初始化了jwt, rabbitmq和mysql分别为后端提供了，token生成和解密，消息队列的读写和数据库的读写。数据库使用了mysql并且使用了gorm包。

- data

  ![image-20210111214822371](D:\arts\gopath\src\pku-class\market\document\img\data.png)

  data中主要是一些数据结构的定义，在gin框架下与前端的通信中，以及与数据库的交互中都使用了这些数据结构。另外data中还有一些函数用于封装这些结构体为前后端通信接口中的json格式。

- backstage

  backstage意为后端的后端，主要是消息队列的消费者处理，用来序列化处理秒杀请求，判断秒杀是否成功并将订单写入数据库。

- 安全日志

  用户登录，每次请求的用户身份，访问时间，数据库的增删改查都有日志记录。

  - token解析的日志

    ![image-20210111220146515](D:\arts\gopath\src\pku-class\market\document\img\token解析.png)

  - 商品修改日志

    ![image-20210111220403242](D:\arts\gopath\src\pku-class\market\document\img\commodity修改.png)

## 数据库

- 表

  ![image-20210111215513435](D:\arts\gopath\src\pku-class\market\document\img\表.png)

- commodities

  ![image-20210111215557085](D:\arts\gopath\src\pku-class\market\document\img\commodity.png)

- menus

  ![image-20210111215633378](D:\arts\gopath\src\pku-class\market\document\img\menus.png)

- orders

  ![image-20210111215707418](D:\arts\gopath\src\pku-class\market\document\img\orders.png)

- users

  ![image-20210111215801748](D:\arts\gopath\src\pku-class\market\document\img\users.png)