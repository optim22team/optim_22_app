package user

import (
  "testing"
//  "github.com/gin-gonic/gin"
//  "optim_22_app/internal/pkg/test"
//  "optim_22_app/pkg/log"
//  "optim_22_app/typefile"
  "github.com/stretchr/testify/assert"
)
  

func TestRegistrationInformationValidate(t *testing.T) {
  tests := []struct {
    name      string
    model     RegistrationInformation
    wantError bool
  }{
    {
      "normal address", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      false,
    },
    {
      "academic address", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@inc.test-ac.jp", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      false,
    },
    {
      "has dot", 
      RegistrationInformation{
        Name:"test", 
        Email:"test.test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      false,
    },
    {
      "invalid name type: empty", 
      RegistrationInformation{
        Name:"", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid hash type: <len", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d65",
      }, 
      true,
    },
    {
      "invalid hash type: >len", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d67722d4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid hash type: !format", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@test.com", 
        Password:" 7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid hash type: !format", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@test.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069 ",
      }, 
      true,
    },
    {
      "invalid hash type: empty", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@test.com", 
        Password:"",
      }, 
      true,
    },
    {
      "invalid email type: no at", 
      RegistrationInformation{
        Name:"test", 
        Email:"testtest.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    /*
    {
      "invalid email type: invalid hyphen", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@test-.com", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    */
    {
      "invalid email type: no TLD", 
      RegistrationInformation{
        Name:"test", 
        Email:"test@testcom", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
    {
      "invalid email type: empty", 
      RegistrationInformation{
        Name:"test", 
        Email:"", 
        Password:"7f83b1657ff1fc53b92dc18148a1d6fffffd4b1fa3d677284addd200126d9069",
      }, 
      true,
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      err := tt.model.Validate()
      assert.Equal(t, tt.wantError, err != nil)
    })
  }
}