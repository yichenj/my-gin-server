# My Gin Server
An API service scaffold based on Golang Gin framework.

## Features
- [X] Strict layered design pattern, suitable for large scale project in real world
- [X] Lightweight and clean implementation
- [X] Least third-party dependency, highly customizable
- [X] Go modules for dependency management
- [X] MIT License

## Usage

### Setup

#### Create Database
We use MySQL database in this demo, you can refer the [MySQL Reference Manual](https://dev.mysql.com/doc/refman/8.0/en/installing.html) to install it on your platform.

After installation, login the MySQL and create a database for the demo. For example, on Mac OS X with `mysql` shell command,
```zsh
mysql -uroot
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 10
Server version: 8.0.18 Homebrew

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> create database test;
Query OK, 1 row affected (0.01 sec)
```
Then, run the database init script to create tables we need. For example,
```zsh
mysql -uroot -Dtest < sql/init_db.sql
```
To check if the database is correctly initialized, you may check if the following table is created in the database you created just now.
```zsh
mysql -uroot -Dtest
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 12
Server version: 8.0.18 Homebrew

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> desc demos;
+-------------+--------------+------+-----+---------+----------------+
| Field       | Type         | Null | Key | Default | Extra          |
+-------------+--------------+------+-----+---------+----------------+
| id          | bigint(20)   | NO   | PRI | NULL    | auto_increment |
| name        | varchar(255) | YES  |     | NULL    |                |
| description | varchar(255) | YES  |     | NULL    |                |
| create_time | datetime     | YES  |     | NULL    |                |
+-------------+--------------+------+-----+---------+----------------+
4 rows in set (0.01 sec)
```
#### Configure for Database

Edit `conf/app_cfg.yml` to tell the server about the database. In this example, we set the `url` field of `database:mysql` to
```yaml
database:
  mysql:
    url: root@tcp(127.0.0.1)/test
```
where `root` is the login name, `127.0.0.1` is the IP address where MySQL service is on, and `test` is the database name you haved created.
For more details about the URL, refer to [DSN](https://github.com/go-sql-driver/mysql/) of MySQL driver for Golang.

#### Compile the Server and Run
Run `go build` under the directory of the project, [Go Modules](https://blog.golang.org/using-go-modules) will download the dependencies automatically and start to build the project.

If the build goes no wrong, there will be a `my-gin-server` executable file output in the directory.

Run it, then you can see the server starts working.
```zsh
./my-gin-server
[INFO] 2020/08/18 16:00:57 main.go:40: Server initializing with config: {Server:{Host: Port:5678 Mode:debug} Database:{Mysql:{Url:root@tcp(127.0.0.1)/test MaxIdleConns:5 MaxOpenConns:0}}}
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /demo                     --> github.com/my-gin-server/api.(*DemoAPI).Create-fm (4 handlers)
[GIN-debug] DELETE /demo                     --> github.com/my-gin-server/api.(*DemoAPI).DeleteRange-fm (4 handlers)
[GIN-debug] DELETE /demo/:id                 --> github.com/my-gin-server/api.(*DemoAPI).Delete-fm (4 handlers)
[GIN-debug] GET    /demo                     --> github.com/my-gin-server/api.(*DemoAPI).List-fm (3 handlers)
[GIN-debug] GET    /demo/:id                 --> github.com/my-gin-server/api.(*DemoAPI).Query-fm (3 handlers)
[GIN-debug] Listening and serving HTTP on :5678
```
#### CLI Flags
Run `my-gin-server` with `--help` args will show all the CLI flags supported.
```zsh
./my-gin-server --help
Usage of ./my-gin-server:
  -config string
    	path to config file (default "conf/app_cfg.yml")
  -log-level string
    	default log level (default "INFO")
```
You can add your own CLI flags in `main.go`.

#### Configuration
```yaml
server:
  host:
  port: 5678
  mode: debug
database:
  mysql:
    url: root@tcp(127.0.0.1)/test
    max_idle_conns: 5
    max_open_conns: 0 #unlimited
```
`conf/app_cfg.yml` gives the full config of our demo server:
* `server:host` and `server:port`: the listen host and port of the server
* `server:mode`: Gin running mode, `debug` or `release`
* `mysql:url`: database access URL
* `max_idle_conns` and `max_open_conns`: connection pool settings for the `database/sql` library

Add your own config items in `base/appconfig/appconfig.go`
### Project Structure
We have implemtented a series of APIs to CRUD records in table `demos` in this project so as to explain how to use this scaffold. Basically, the project follows MVC pattern, and the controller part is further splitted into three layers: API layer, Service layer and DAO layer. The following file structure describes what each file is used for.

```zsh
.
├── LICENSE
├── README.md
├── api                         # API layer, defines all handler functions bound to a URL, handles http input/output
│   ├── demo_api.go             # 'demo' API implementation
│   └── vo                      # Defines the API layer input/output speficition, kind of closing to View of MVC
│       └── demo_vo.go          # 'demo' API JSON input/output struct
├── base
│   ├── appconfig
│   │   └── appconfig.go        # Config definations and implementation
│   ├── apperror
│   │   ├── apperror.go         # Error handling implementation
│   │   └── error_code.go       # All error code definations
│   ├── applog
│   │   └── applog.go           # Log functions wrapper
│   └── db
│       ├── facade.go           # Database interface wrapper
│       └── mysql.go            # MySQL related implementation
├── conf
│   └── app_cfg.yml             # Server config file
├── dao                         # DAO layer, defines database access functions (SQL implementation mostly)
│   └── demo_dao.go             # 'demo' API DAO implementation
├── go.mod                      # Go module file
├── go.sum                      # Go module generated
├── init.go                     # Init API service process: initialize API/Service/Dao instance mainly
├── main.go                     # Entry: initialize infrastructure, CLI flags, log, databases and so on
├── middleware                  # Gin middleware functions
│   └── middleware.go           # A simple middleware example
├── model                       # Model layer, persist the object in database
│   └── demo.go                 # Map to 'demos' table in database
├── my-gin-server               # Binary file to run
├── router                      # Router protocol, map URL path to its handler function
│   └── demo_router.go          # 'demo' API supported path 
├── service                     # Service layer, main service logic
│   ├── demo_service.go         # 'demo' API service implementation
│   └── dto                     # Defines the Service layer input/output, often used when need to re-organize multiple inputs/outputs and transfer them in services.
│       └── demo_dto.go         # 'demo' API DTO definations
└── sql
    └── init_db.sql             # Database init script
```

## Design Considerations
### SQL vs ORM
### Dependency Injection
### Log Library
