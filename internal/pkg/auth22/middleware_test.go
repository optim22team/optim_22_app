package auth22

import (
  "net/http"
  "testing"
  "github.com/gin-gonic/gin"
  "optim_22_app/internal/pkg/test"
  "optim_22_app/internal/pkg/config"
  "optim_22_app/pkg/log"

)

const (
  accessToken2021 = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzM0NDQ0MTEsInVzZXJJRCI6Mn0.Vr3t57_ty9jBtNpLfYzupz59stbEHciPaPbZEFT2J88`
  refreshToken2022 = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTk5MzA4MTEsInVzZXJJRCI6Mn0.qbU04Ub-s5O7dy5NAvb4modDqQ4_iSqgmjH-sCQRaZw`
  /*
   *
   *  jwtパッケージにより生成
   *
   */
  refreshToken2010 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjEyODM2NTg3Mjh9.krKE34GBpQBMwSMFHf8iMpM36fxycGLvUf9Mi70--cM"
  refreshToken2020 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjE1NzgxOTMyNzl9.QrcRvgE6PbiqpAI9eLM9TeQWe6iRt0tEb-rQvnp7U_E"
  refreshToken2030 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOiIwMDEiLCJleHAiOjE5MTQ4MTA3Mjh9.QHGNRk1KMQyx8rLscYdkKxQ7nBp7ZmcLDsF8fsk40dA"
  refreshToken2000 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjk2ODI0MTk5OX0.xHLUtwyp-R_rKUJaW0HtZlk6DStLBLyKG-NZie3YvVQ"
  refreshToken2100 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.VJbsifEaA5uaGmJdH__e270WJ20hxrlGQF79jc789vw"
  noExpRefreshToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEifQ.1LotTqA4yjwjMk9SLHPKJ3ggH2Z0j1ADVyFZqDNkZbM"
  noUserIdRefreshToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjk2ODA0NzQxN30.-pwuyxI6oFh7nkbKzdRU-3u-F6baQAMtKjwKcNRWMVo"
  accessToken2000 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjk2ODA0NzQxN30.x6_2QJHDmemdSz7ev6By6iyAtpWibjZLbBWZZCd3Q-U"
  accessToken2100 = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEiLCJleHAiOjQxMjM3MjEwMTd9.JezlU23njkKGldV4ZH1QI37O1yCd0Y-mWmnIu-7aKEo"
  noExpAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIwMDEifQ.oRHMenxs4DlJy79Has9ASiu0qD0MFh9vmYevOOksizE"
  noUserIdAccessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjQxMjM3MjEwMTd9.pLQbblBZu_rNBqB97E6U_rFUess4Utz2tPA3aAd8mcY"
  /*
  * pythonパッケージPyJWTをによりテストトークンを生成
  *
  * #python3 -m pip install PyJWT==2.1.0
  *
  * #2010: 1283658728
  * #2015: 1420426879
  * #2020: 1578193279
  * #2025: 1736046079
  * #2030: 1914810728
  * 
  * import jwt 
  * 
  * payload_data = {
  *     "userID": "001",
  *     "exp": 4123721017 #2100年
  * }
  * 
  * token = jwt.encode(
  *     algorithm="HS256",
  *     payload=payload_data,
  *     key='secret_key', 'secret_key_for_refresh'
  * )
  * 
  * print(token)
  *
  */
)


func TestAccessTokenAuthentication(t *testing.T) {

  router := gin.New()
  logger := log.New()
  cfg, _ := config.Load("test/app.yaml", logger)
  logger.Debug(cfg.RefreshTokenSecret)
  auth := New(NewService(cfg, nil, logger), "localhost")
  
  router.Use(auth.ValidateAccessToken(GetRuleForTest(), true))


  router.POST("/test", func(c *gin.Context) {
    c.String(http.StatusOK, "")
  })

  tests := []test.APITestCase{
    {
      "authentication success", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader("", accessToken2100), 
      http.StatusOK, 
      "",
    },
    {
      "authentication failed: expired", 
      "POST", 
      "/test", 
      "", 
      MakeAuthorizationHeader("", accessToken2000), 
      http.StatusForbidden, 
      "",
    },
  }
  for _, tc := range tests {
    test.Endpoint(t, router, tc)
  }
}


func GetRuleForTest() Rule {
  return Rule{
    "GET": map[string]bool{
      "*": true,
    },
  }
}