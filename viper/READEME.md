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