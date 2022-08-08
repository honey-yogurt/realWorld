package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type Profile struct {
	ID        uint
	Username  string
	Bio       string
	Image     string
	Following bool
}

// UserRepo
// @Description: 操作数据库的接口
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	//GetUserByEmail(ctx context.Context, email string) (*User, error)
	//GetUserByUsername(ctx context.Context, username string) (*User, error)
	//GetUserByID(ctx context.Context, id uint) (*User, error)
	//UpdateUser(ctx context.Context, user *User) (*User, error)
}

//type ProfileRepo interface {
//	GetProfile(ctx context.Context, username string) (*Profile, error)
//	FollowUser(ctx context.Context, currentUserID uint, followingID uint) error
//	UnfollowUser(ctx context.Context, currentUserID uint, followingID uint) error
//	GetUserFollowingStatus(ctx context.Context, currentUserID uint, userIDs []uint) (following []bool, err error)
//}

type UserUseCase struct {
	ur UserRepo
	//pr  ProfileRepo
	log *log.Helper
}

//func NewUserUseCase(ur UserRepo, pr ProfileRepo, log *log.Helper) *UserUseCase {
func NewUserUseCase(ur UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		ur:  ur,
		log: log.NewHelper(logger),
	}
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

func (uc *UserUseCase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	// biz --->data 层，调用数据库，将u 入库
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		//Token:    "",
		//Bio:      "",
		//Image:    "",
	}, nil
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}
