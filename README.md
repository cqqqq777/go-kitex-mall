# go-kitex-mall
The mall/eshop system based on kitex and hertz

## Project Implementation

### Technology Selection

- HTTP frame: [Hertz](https://www.cloudwego.io/zh/docs/hertz/)
- RPC frame: [Kitex](https://www.cloudwego.io/zh/docs/kitex/)
- Relational Database: Mysql
- Non-relational database: Redis  MongoDB
- The configuration center: [Nacos](https://nacos.io/zh-cn/docs/what-is-nacos.html)
- The service center: Nacos
- The message queue: [Nsq](https://nsq.io/)
- The object storage service: [Minio](https://min.io/)
- Traceing: Yaeger
- Monitoring: Prometheus 
- Permission Management: Casbin

### Function

#### User

- [ ] Register
- [ ] Login
- [ ] Login by Github
- [ ] Forget password
- [ ] User info
- [ ] Goods list
- [ ] Search product
- [ ] Collect product
- [ ] View the collected
- [ ] Manage the Collected
- [ ] Shopping carts
- [ ] Order service
- [ ] Comment
- [ ] History
- [ ] Buy
- [ ] Chat with merchant

#### Merchant

- [ ] Publish product
- [ ] Delete product
- [ ] Change product info
- [ ] Chat with consumer

### Service Split

- User
- Product
- Cart
- Pay
- Order
- Comment
- Chat
- Merchant
- Operation

