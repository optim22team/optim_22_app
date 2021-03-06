package comment

import (
  "net/http"
  "strconv"
  "encoding/json"
  "github.com/gin-gonic/gin"
  "optim_22_app/pkg/log"
)

//コメント操作の依存関係
type resource struct {
  service Service
  logger  log.Logger
}

//コメント操作についてエンドポイントを登録
func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {

  rc := resource{service, logger}

  //ディスカッション ID(:requestID) のコメント一覧を取得
  r.GET("/:requestID", rc.get())
  //ディスカッション ID(:requestID) にコメントを投稿
  r.POST("/:requestID", rc.post())

}


func (rc resource) get() gin.HandlerFunc {
  return func(c *gin.Context) {
    requestID := c.Param("requestID")

    if !isIntegerString(requestID) {
      c.Status(http.StatusBadRequest)
      return
    }

    comments, err := rc.service.Get(c.Request.Context(), requestID)
    //rc.logger.Debug(comments)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusNotFound)
      return 
    } else {
      if commentsText, err := json.Marshal(comments); err != nil {
        rc.logger.Error(err)
        c.Status(http.StatusInternalServerError)
        return
      } else {
        //rc.logger.Debug(reflect.TypeOf(commentsText))
        c.Header("Content-Type", "application/json")
        c.String(http.StatusOK, Presenter(commentsText))
        return
      }
    }
  }
}


func (rc resource) post() gin.HandlerFunc {
  return func(c *gin.Context) {
    var input comment

    requestID := c.Param("requestID")
  
    //BodyからJSONをパースして読み取る
    if err := c.ShouldBindJSON(&input); err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    }
    
    //コメントを登録
    err := rc.service.Post(c.Request.Context(), input, requestID)
    if err != nil {
      rc.logger.Error(err)
      c.Status(http.StatusBadRequest)
      return 
    } else {
      c.Status(http.StatusCreated)
      return 
    }
  }
}

func isIntegerString(query string) bool {
  _, err := strconv.Atoi(query)
  return err == nil
}