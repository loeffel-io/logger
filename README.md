# logger

Private sentry and logrus logger

## Install 

```bash
go get -u github.com/loeffel-io/logger/v2
```

## Example

### Setup
```go 
if err = sentry.Init(sentry.ClientOptions{
    Dsn:       os.Getenv("SENTRY"),
    Transport: sentry.NewHTTPSyncTransport(),
}); err != nil {
    log.Fatal(err)
}

log.SetFormatter(&log.TextFormatter{
    DisableColors: false,
    FullTimestamp: true,
})

logger := &l.Logger{
    Debug:     true,
    SentryHub: sentry.CurrentHub(),
    RWMutex:   new(sync.RWMutex),
}
```

### Logging

```go 
logger.Error(fmt.Errorf("test"))
logger.Log(fmt.Errorf("test"))
logger.Print("test")
```

### Gin Middleware

```go
r := gin.Default()
r.Use(gin.Recovery())
r.Use(gin_middleware.Logger(logger))
```