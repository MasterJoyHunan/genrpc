package config

type Server struct {
    Name string
    Host string
    Port string
}

type Mysql struct{
    Host string
    Port int
    User string
    Pwd  string
    Db   string
}

type Redis struct{
    Host string
    Port int
    User string
    Pwd  string
    Db   int
}
