# viper

### 读取配置文件

```go
func main() {
    v := viper.New()
    v.SetConfigFile("./config/config.yaml")
    if err := v.ReadInConfig(); err != nil {
    panic(err)
    }
    fmt.Println(v.Get("name"))
}
```

### 序列化成结构体
```go
type ServerConfig struct {
    Name string `mapstructure:"name"`
    }
    
    func main() {
    var server ServerConfig
    v := viper.New()
    v.SetConfigFile("./config/config.yaml")
    if err := v.ReadInConfig(); err != nil {
    panic(err)
    }
    err := v.Unmarshal(&server)
    if err != nil {
    panic(err)
    }
    
    fmt.Println(server.Name)
}

```