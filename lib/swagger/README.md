# Swagger
## API文档生成（Simplify API development for users, teams, and enterprises with the Swagger open source and professional toolset. Find out how Swagger can help you design and document your APIs at scale.）
## [官网](https://swagger.io/)&&[GitHub](https://github.com/swagger-api)
## 生态
| 名称 | 描述 |
| ---- | ---- |
| Swagger-UI | API的可视化界面工具 |
| Swagger-Inspector | 在线的API测试工具，附带chrome插件|
| Swagger-Editor | API可视化的设计工具并可生成相关代码 |
# spring boot 集成
- 添加springfox-swagger2，springfox-swagger-ui
- 注入Docket的配置bean,并使用 @EnableSwagger2注解
- swagger-ui需配置mvc资源访问或将其解压复制到public目录下