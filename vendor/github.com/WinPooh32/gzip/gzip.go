package gzip

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	BestCompression    = gzip.BestCompression
	BestSpeed          = gzip.BestSpeed
	DefaultCompression = gzip.DefaultCompression
	NoCompression      = gzip.NoCompression
)

func Gzip(level int) gin.HandlerFunc {
	var gzPool sync.Pool
	gzPool.New = func() interface{} {
		gz, err := gzip.NewWriterLevel(ioutil.Discard, level)
		if err != nil {
			panic(err)
		}
		return gz
	}
	return func(c *gin.Context) {
		if !shouldCompress(c.Request) {
			return
		}

		gz := gzPool.Get().(*gzip.Writer)
		defer gzPool.Put(gz)
		defer gz.Reset(ioutil.Discard)
		gz.Reset(c.Writer)

		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding")
		c.Writer = &GzipWriter{c.Writer, gz, false}
		defer func() {
			gz.Close()
			c.Header("Content-Length", fmt.Sprint(c.Writer.Size()))
		}()
		c.Next()
	}
}

type GzipWriter struct {
	gin.ResponseWriter
	writer          *gzip.Writer
	skipCompression bool
}

func (g *GzipWriter) WriteString(s string) (int, error) {
	if g.skipCompression {
		return g.ResponseWriter.Write([]byte(s))
	} else {
		return g.writer.Write([]byte(s))
	}
}

func (g *GzipWriter) Write(data []byte) (int, error) {
	if g.skipCompression {
		return g.ResponseWriter.Write(data)
	} else {
		return g.writer.Write(data)
	}
}

// Fix: https://github.com/mholt/caddy/issues/38
func (g *GzipWriter) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}

func (g *GzipWriter) SetSkipCompression(c *gin.Context) {
	g.skipCompression = true
	c.Header("Content-Encoding", "")
	c.Header("Vary", "")
}

func shouldCompress(req *http.Request) bool {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") ||
		strings.Contains(req.Header.Get("Connection"), "Upgrade") ||
		strings.Contains(req.Header.Get("Content-Type"), "text/event-stream") {

		return false
	}

	extension := filepath.Ext(req.URL.Path)
	if len(extension) < 4 { // fast path
		return true
	}

	switch extension {
	case ".png", ".gif", ".jpeg", ".jpg":
		return false
	default:
		return true
	}
}
