# 项目名：广州缀美画室网站
# 开发语言:Go!!!!

# 项目规划
	环境要求：golang=1.6,mysql=5.6
	起始时间：2016.6.10
	结束时间：2016.

# 项目托管
	阿里云服务器：
	域名：http://www.beautyart.top

# 参考

1. 基于角色的访问控制（Role-Based Access Control）作为传统访问控制
2. 使用beego框架和大量javascript脚本ajax调用
3. Amaze UI v2.7.0和jQuery EasyUI 1.4.2、Bootstrap混合

# 运行步骤

1. 运行init.sh进行包初始化
2. 接着

```
	cd beauty
	go build main.go
	./main -s
	./main
```

3. 平台使用说明参见doc文件夹

可自由修改源代码，但必须保留友好链接

	[http://wwww.beautyart.top](广州缀美画室)

# 联系方式
	https://www.github.com/hunterhug 
	QQ：569929309


# 文件目录
````
beautyart
----conf 配置文件夹

	----app.conf 		应用配置文件
	----local_**.ini 	国际化文件

----controllers 控制器
	----admin	后台控制器	
	----home 	前台控制器

-----lib 公共库

-----models ORM模型
	----admin RBAC主要数据库
	----home 

----routers 路由
----static  静态文件
----views	视图
	----admin 	后台视图
		----default 默认主题
	----home 	前台视图
		----default 默认主题

----log 日志
----doc 说明文档
```

# 项目约定

1. RBAC权限相关的models统一放在admin文件夹，其他都放在home文件夹.
	前台控制相关的controllers统一放在home文件夹，其他都放在admin文件夹
	URL router统一M/C/A方式，该正则url需要验证权限，如rbac/public/index，其他如public/index不验证。

2. 登录说明
	登陆过的用户只能注销后登录，支持定义cookie登录。进入后台时验证session，session不存在则验证cookie，如果用户
	未被冻结，增加session，同时更改用户登录时间、登录IP等，cookie与登录IP绑定。

3. 系统时间默认数据库本地时间为东八区，北京时间。

4. 后台模板在views/admin，前台模板在views/home，子文件夹为主题，默认主题为default

5. 所有配置在conf文件夹conf/app.conf，支持国际化

6. 数据库数据填充在models/*/*Init.go中定义


# 温馨提示

1. Amaze Ui与EasyUi冲突，Amazeui.css第一行注释掉
```
	*,
	*:before,
	*:after {
	  -webkit-box-sizing: border-box;
	          box-sizing: border-box;
	}
```

# 项目进展
1. 开发手脚架搭建完毕，RBAC模块完成，2016/7/2
