## Requirement

```
go 1.11 with go mudules
github.com/gorilla/context
github.com/jinzhu/gorm/dialects/mysql
github.com/gorilla/sessions
github.com/jinzhu/gorm
github.com/spf13/viper
```

## Get Started

### clone

```
cd $GOPATH/src
git clone git@github.com:pikapiyyy/blog.git
```

### go mod tidy

Before install required package,you need make your `go mod` can be used which means your go version must upper or equal than 1.11.
```
vim /etc/profile
```
在文件末尾加入
```
export GOPROXY=https://goproxy.io
export GO111MODULE=on
```
回到项目目录`$GOPATH/src/blog`
```
go mod tidy
```

### database initialize
```
cp config.yml.config config.yml
```
创建数据库`blog`(名称可自定义)并编辑`config.yml`修改成你的数据库的对应信息

```
go run console/db_init.go
```

### start
```
go run main.go
```
Your blog web server will listen on port 8888,visit URL `http://你的ip:8888` on website,you can see your blog.
