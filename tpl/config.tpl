package {{.pkgName}}

type Server struct {
    Name string
    Host string
    Port string
}

type Mysql struct{
    Host string
    Port string
    User string
    Pwd  string
    Db   string
}

type Redis struct{
    Host string
    Port string
    User string
    Pwd  string
    Db   string
}