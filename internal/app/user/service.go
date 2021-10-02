package user

import (
  "regexp"
  "github.com/go-ozzo/ozzo-validation/v4"
//  "github.com/go-ozzo/ozzo-validation/v4/is"
  "optim_22_app/pkg/log"
  "optim_22_app/typefile"
  "context"
)

//#region 登録情報
//`POST /api/user`が要求する情報
type RegistrationInformation struct {
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"password"`
}


func (m RegistrationInformation) Validate() error {
  return validation.ValidateStruct(&m,
    validation.Field(&m.Name, validation.Required, validation.Length(3, 128)),
    //is.Email@ozzo-validation/v4/isはテストケース`success#1`にてエラー
    validation.Field(&m.Email, validation.Required, validation.Match(regexp.MustCompile("[a-zA-Z]+[a-zA-Z0-9\\.]@[a-zA-Z]+((\\.[a-zA-Z0-9\\-])+[a-zA-Z0-9]+)+"))),
    //is SHA256
    validation.Field(&m.Password, validation.Required, validation.Length(64, 64), validation.Match(regexp.MustCompile("[A-Fa-f0-9]{64}$"))),
  )
}
//#endregion

type Service interface {
  Create(ctx context.Context, input RegistrationInformation) (int, error)
  Delete(ctx context.Context, userId int) error
}


type service struct {
  repo   Repository
  logger log.Logger
}

//新たなuser作成サービスを作成
func NewService(repo Repository, logger log.Logger) Service {
  return service{repo, logger}
}


func (s service) Create(ctx context.Context, req RegistrationInformation) (int, error) {
  //リクエストの値を検証
  if err := req.Validate(); err != nil {
    return 0, err
  }
  //クエリの値を定義
  insertValues := typefile.User{
    Name:      req.Name,
    Email:     req.Email,
    Password:  req.Password,
  }
  //INSERTと割り当てられるuserIDを取得
  var userId int
  if err := s.repo.Create(ctx, &insertValues); err != nil {
    return 0, err
  } else {
    userId = insertValues.ID
  }

  return userId, nil
}


func (s service) Delete(ctx context.Context, userId int) error {
  //該当useriDのエントリを削除
  if err := s.repo.Delete(ctx, userId); err != nil {
    return err
  } else {
    return nil
  }
}


func StubNewService(args ...interface{}) Service { return service{nil, nil}}
func StubCreate(args ...interface{}) (string, string, error)  {return "", "", nil}
func StubDelete(args ...interface{}) error {return nil}
func StubLogin(args ...interface{}) (string, string, error) {return "", "", nil}
