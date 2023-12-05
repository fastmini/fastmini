# GoApp

## 基于fiber框架改造的性能炸裂的一站式解决方案

### 功能优势

* 整合MySql客户端
* 整合Redis客户端
* 整合ES客户端
* 整合xxl-job客户端
* 整合dingding客户端

### 稳定性测试

* 上亿数据实践

#### 生成model

```
gentool -dsn "test:test@@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local" -tables "user,post" -onlyModel
```