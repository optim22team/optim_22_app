// type.goに利用する構造体をまとめる。
// 構造体の属性は必ず最初の文字を大文字にする。
// Winnerが必要になるかもしれない
package typefile

import (
	"time"
)

type User struct{
	ID             int            `gorm:"primaryKey"`
	Name           string         `gorm:"not null"`
}

type Client struct{
	User           User           `gorm:"embedded"`
}

type Engineer struct{
	User           User           `gorm:"embedded"`
	Requests       []Request      `gorm:"many2many:engineer_requests;"`
}

type Winner struct{
	EngineerID     int            `gorm:"not null"`
	Engineer       Engineer
	RequestID      int            `gorm:"unique;not null"`
}

type Request struct{
	ID             int            `gorm:"primaryKey"`
	ClientID       int            `gorm:"not null"`
	Client         Client
	RequestName    string         `gorm:"not null"`
	Content        string         `gorm:"not null"`
	Winner         Winner
	Finish         bool           `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Submission struct{
	ID             int            `gorm:"primaryKey"`
	RequestID      int            `gorm:"not null"`
	Request        Request
	EngineerID     int            `gorm:"not null"`
	Engineer       Engineer
	Content        string         `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}