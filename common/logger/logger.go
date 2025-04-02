package logger

import (
    "os"
    "net/http"
    "strings"
    "bytes"
    "fmt"
    "encoding/json"
    _ "time"

	"github.com/sirupsen/logrus"
    "github.com/go-chi/chi/v5/middleware"
)

type LoggingResponseWriter struct {
    middleware.WrapResponseWriter
    body bytes.Buffer
}

type customTextFormatter struct {}

func NewLoggingResponseWriter(w http.ResponseWriter, protoMajor int) *LoggingResponseWriter {
    return &LoggingResponseWriter {
        WrapResponseWriter: middleware.NewWrapResponseWriter(w, protoMajor),
    }
}

func (w *LoggingResponseWriter) Write (data []byte) (int, error) {
    w.body.Write(data)

    return w.WrapResponseWriter.Write(data)
}

func (w *LoggingResponseWriter) GetBody() string {
    return w.body.String()
}

func (f *customTextFormatter) Format(entry *logrus.Entry) ([]byte, error){
    var b bytes.Buffer

    b.WriteString(fmt.Sprintf("[%s] [%s] [%s:%d] %s", 
        entry.Time.Format("2006-01-02 15:04:05"),
        entry.Level.String(),
        entry.Caller.Func.Name(),
        entry.Caller.Line,
        entry.Message,
    ))

    if len(entry.Data) > 0 {
        for key, value := range entry.Data {
            b.WriteString(fmt.Sprintf(" %s=%v", key, value))
        }
    } else {
        b.WriteString(" <no data>")
    }
    b.WriteString("\n")   
    return b.Bytes(), nil
}

var logger *logrus.Logger

func init() {
    logger = logrus.New() 
    logger.SetOutput(os.Stdout)

    logger.SetReportCaller(true)
    logger.SetFormatter(&customTextFormatter{})
    logger.Info("Logger has been init")

}

func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        ww := NewLoggingResponseWriter(w, r.ProtoMajor)
        
        var request []byte
        json.NewDecoder(r.Body).Decode(&request)

        logger.WithFields(logrus.Fields{
            "method": r.Method,
            "path": r.URL.Path,
            "request": string(request)}).Info("Request was sended")
        

        next.ServeHTTP(ww, r)
       
        raw := json.RawMessage(ww.body.Bytes())
        
        defer logger.WithFields(logrus.Fields{
            "method": r.Method,
            "path": r.URL.Path,
            "status": ww.Status(),
            "response": strings.ReplaceAll(string(raw), "\n", "")}).Info("Request was done")
    })
}
