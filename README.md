# GRB-Server
基于Gin + Gorm + Mysql + Redis + Meilisearch的个人博客后端
## 依赖
| Environment | Version     |
| ----------- | ----------- |
| Golang      | 1.22.4      |
| Mysql       | 8.0.37      |
| Redis       | 7.0.15      |
| [Meilisearch](https://www.meilisearch.com) | 1.9.0       |

## 工具
### swag用于生成接口文档
```
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```
## 配置文件
自行新建开发环境配置文件/config/dev.yaml.
参考/config/pro.yaml编写配置文件,默认配置为开发环境.
## 数据库
新建mysql数据库``databasename``.
```sql
CREATE DATABASE databasename DEFAULT CHARACTER SET = 'utf8mb4';
```
在dev/pro.yaml中指定``databasename``数据库名，运行项目即可利用Gorm自动初始化数据库表结构.
## 运行
```shell
make dev # 开发环境  args: -config ./config/dev.yaml
make pro # 生产环境  args: -config ./config/pro.yaml
```
## 接口文档

