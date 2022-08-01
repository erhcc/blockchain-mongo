package main

import (
	"time"
	"go.uber.org/zap"
)

func main()  {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url:="testurl"
	sugar.Infow("failed to fetch URL",
	// Structured context as loosely typed key-value pairs.
	"url", url,
	"attempt", 3,
	"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	loggerZap()
}

/*
In the rare contexts where every microsecond and every allocation matter, use the Logger. 
It's even faster than the SugaredLogger and allocates far less, but it only supports strongly-typed, structured logging.
*/
func loggerZap()  {
	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
	  zap.String("url", "http://example.com"),
	  zap.Int("attempt", 3),
	  zap.Duration("backoff", time.Second),
	)
}