package authentication

import (
  "net/http"
  "time"
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
//  "fmt"
)

//type ginHandler func(*gin.Context)
//
//



func (auth *Authorizer) RefreshTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.AbortWithStatus(http.StatusBadRequest)
    } else {

      token, _ := jwt.Parse(refreshToken, auth.refreshTokenSecretSender)
      _, ok := token.Claims.(jwt.MapClaims)
  
      if ok {
        if token.Valid {
          auth.authorizationService.Endpoint(c)
        } else {
          c.AbortWithStatus(http.StatusUnauthorized)
        }
      } else {
        c.AbortWithStatus(http.StatusBadRequest)
      }
    }
  }
}

//リフレッシュトークンが有効期限内のとき任意のアクセストークンの有効期限を延長
func (auth *Authorizer) AccessTokenRefreshHandler() gin.HandlerFunc {
  return func(c *gin.Context) {

    refreshToken, err := c.Cookie("refresh_token")

    if err != nil {
      c.AbortWithStatus(http.StatusBadRequest)
    } else {

      token, _ := jwt.Parse(refreshToken, auth.refreshTokenSecretSender)
      _, ok := token.Claims.(jwt.MapClaims)

      if ok {
        //リフレッシュトークンが期限内
        if token.Valid {

          //#region アクセストークンのclaimsを取り出し有効期限を更新し符号化後、署名

          //Authorizationヘッダーからstring型のトークンを取得
          tokenString := c.GetHeader("Authorization")
          //トークンの改竄と期限を検証
          token, _ := jwt.Parse(tokenString, auth.accessTokenSecretSender)
          //claimsを辞書型として取得
          claims, ok := token.Claims.(jwt.MapClaims)
      
          if ok {
            expiration := time.Now()
            expiration = expiration.Add(auth.validityPeriod)

            claims["exp"] = expiration.Unix()
            newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)              
            // Sign and get the complete encoded token as a string using the secret
            newTokenString, _ := newToken.SignedString([]byte(auth.accessTokenSecret))
            c.Header("Authorization", newTokenString)

            c.Status(http.StatusOK)
          } else {
            c.AbortWithStatus(http.StatusBadRequest)
          }

          //#endregion
        } else {
          
          c.AbortWithStatus(http.StatusUnauthorized)
        }
      } else {
        c.AbortWithStatus(http.StatusBadRequest)
      }
    }
  }
}

//認証情報を空文字列で上書き
func (auth *Authorizer) RevokeHandler() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.SetCookie("refresh_token", "", 0, "/", "", false, true)
    c.Header("Authorization", "")
  }
}

//パース関数にリフレッシュトークン用秘密鍵を渡すコールバック
func (auth *Authorizer) refreshTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return auth.refreshTokenSecret, nil
}

//パース関数にアクセストークン用秘密鍵を渡すコールバック
func (auth *Authorizer) accessTokenSecretSender(token *jwt.Token) (interface{}, error) {
  return auth.accessTokenSecret, nil
}