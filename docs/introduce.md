# GyuBlog 阅读指引说明

## 待加入两大基础业务功能
- 标签管理：文章所归属的分类
- 文章管理：指文章内容的管理

## 项目结构

- configs：配置文件
- docs：文档集合
- global：全局变量
- internal：内部模块
- dao：数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL、ElasticSearch 等
- middleware：HTTP 中间件
- model：模型层，用于存放 model 对象
- routers：路由相关逻辑处理
- service：项目核心业务逻辑
- pkg：项目相关的模块包
- storage：项目生成的临时文件
- scripts：各类构建，安装，分析等操作的脚本
- third_party：第三方的资源工具，例如 Swagger UI


## MySql 
- Model：指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小写下划线驼峰。若情况特殊，也可以编写该结构体的 TableName 方法用于指定其对应返回的表名。 
- Where：设置筛选条件，接受 map，struct 或 string 作为条件。 
- Offset：偏移量，用于指定开始返回记录之前要跳过的记录数。
- Limit：限制检索的记录数。 
- Find：查找符合筛选条件的记录。
Updates：更新所选字段。
Delete：删除数据。
Count：统计行为，用于统计模型的记录数。