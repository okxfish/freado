fread in golang - simple RSS Server

后端：用 Gin 快速搭建基础restful风格API，Gin 是一个go语言编写的Web框架。

数据库：采用MySql(8.0)版本，使用 gorm 实现对数据库的基本操作。

缓存：Redis

API文档：使用Swagger构建自动化文档。

配置文件：使用 fsnotify 和 viper 实现yaml格式的配置文件。

日志：使用 zap 实现日志记录。


    ├── server
        ├── api             (api层)
        │   └── v1          (v1版本接口)
        ├── config          (配置包)
        ├── core            (核心文件)
        ├── docs            (swagger文档目录)
        ├── global          (全局对象)                    
        ├── initialize      (初始化)                        
        │   └── internal    (初始化内部函数)                            
        ├── middleware      (中间件层)                        
        ├── model           (模型层)                    
        │   ├── request     (入参结构体)                        
        │   └── response    (出参结构体)                            
        ├── packfile        (静态文件打包)                        
        ├── resource        (静态资源文件夹)                        
        │   ├── excel       (excel导入导出默认路径)                        
        │   ├── page        (表单生成器)                        
        │   └── template    (模板)                            
        ├── router          (路由层)                    
        ├── service         (service层)                    
        ├── source          (source层)                    
        └── utils           (工具包)                    
            ├── timer       (定时器接口封装)                        
            └── upload      (oss接口封装)                        



API

    ├── auth
        ├── OAuth2.0          
        │   ├── google
        │   ├── wexin
        │   └── github
        ├── User information
        
    ├── API
        ├── subscription              订阅相关
        │   ├── Add subscription           
        │   ├── Edit subscription           
        │   └──Subscription list  
        ├── Other        
        │   ├── Mark all as read        
        │   └── Unread counts        
        ├── Article                   文章相关
        │   ├── Item IDs                   
        │   ├── Stream contents         
        │   ├── Stream preferences list                              
        │   └──  Stream preferences set       
        ├── Folder/Tag                tag
        │   ├── Folder/Tag list            
        │   ├── Rename tag                             
        │   ├── Delete tag                    
        │   └──  Edit tag                    
        └── search                    搜索
            ├── Create active search                             
            └── Delete active search                           

