package model

import (
	"context"
	"time"
)

type UserService interface {
	GetById(ctx context.Context, id uint) (*User, error)
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, user *User) error
	GetAllUsers(ctx context.Context) ([]User, error)
	AddRoleToUser(ctx context.Context, user_id uint, role_id uint) error
	UpdateUserArea(ctx context.Context, user_id uint, area_id uint) error
}

type UserRepository interface {
	FindById(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	AddRoleToUser(ctx context.Context, user_id uint, role_id uint) error
	UpdateUserArea(ctx context.Context, user_id uint, area_id uint) error
}

type TokenService interface {
	GetNewPairFromUser(ctx context.Context, user *User, prevTokenId string) (*TokenPair, error)
	ValidateIdToken(tokenString string) (*User, error)
	ValidateRefreshToken(refresh_token string) (*RefreshToken, error)
	SignOut(ctx context.Context, user_id uint) error
}

type TokenRepository interface {
	SetNewRefreshToken(ctx context.Context, userId string, tokenId string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userId string, prevTokenId string) error
	DeleteUserRefreshTokens(ctx context.Context, userId string) error
}

type RoleService interface {
	GetRoles(ctx context.Context) ([]Role, error)
	GetRoleByName(ctx context.Context, role_name string) (*Role, error)
}

type RoleRepository interface {
	GetRoles(ctx context.Context) ([]Role, error)
	GetRoleByName(ctx context.Context, role_name string) (*Role, error)
}

type AreaService interface {
	GetAreas(ctx context.Context) ([]Area, error)
	AddArea(ctx context.Context, area *Area) (*Area, error)
}

type AreaRepository interface {
	GetAreas(ctx context.Context) ([]Area, error)
	CreateArea(ctx context.Context, area *Area) (*Area, error)
}

type QuestionService interface {
	AddQuestion(ctx context.Context, question *Question) (*Question, error)
	Clasify(ctx context.Context, question_id uint, area_id uint) error
	TakeQuestion(ctx context.Context, user *User, question_id uint) error
	ResponseQuestion(ctx context.Context, user *User, question_id uint, response string) error
	UpLevel(ctx context.Context, user *User, question_id uint) error
	UpToAdmin(ctx context.Context, user *User, question_id uint) error
	GetMyQuestions(ctx context.Context, user_id uint) ([]Question, error)
	GetUnClasifiedQuestions(ctx context.Context) ([]Question, error)
	GetQuestionsByStatus(ctx context.Context, status uint, area_id *uint) ([]Question, error)
	GetTakenQuestions(ctx context.Context, user_id uint) ([]Question, error)
	GetTakenQuestionById(ctx context.Context, question_id uint, user_id uint) (*Question, error)
	GetQuestionById(ctx context.Context, question_id uint) (*Question, error)
}

type QuestionRepository interface {
	CreateQuestion(ctx context.Context, question *Question) (*Question, error)
	Clasify(ctx context.Context, question_id uint, area_id uint) error
	GetById(ctx context.Context, question_id uint) (*Question, error)
	TakeQuestion(ctx context.Context, user_id uint, question_id uint) error
	ResponseQuestion(ctx context.Context, question_id uint, response string) error
	UpLevel(ctx context.Context, question_id uint) error
	UpToAdmin(ctx context.Context, question_id uint) error
	GetMyQuestions(ctx context.Context, user_id uint) ([]Question, error)
	GetUnClasifiedQuestions(ctx context.Context) ([]Question, error)
	GetQuestionsByStatus(ctx context.Context, status uint, area_id *uint) ([]Question, error)
	GetTakenQuestions(ctx context.Context, user_id uint) ([]Question, error)
	GetTakenQuestionById(ctx context.Context, question_id uint, user_id uint) (*Question, error)
	GetQuestionById(ctx context.Context, question_id uint) (*Question, error)
}

type ChatService interface {
	SendMessage(ctx context.Context, question_id uint, user_id uint, text string) (*MessageChat, error)
	GetMessages(ctx context.Context, question_id uint, user_id uint) ([]MessageChat, error)
	CantUnreadedMessage(ctx context.Context, question_id uint, user_id uint) (int, error)
}

type ChatRepository interface {
	SendMessage(ctx context.Context, question_id uint, user_id uint, text string) (*MessageChat, error)
	GetMessages(ctx context.Context, question_id uint) ([]MessageChat, error)
	ReadMessages(ctx context.Context, question_id uint, user_id uint) error
	CantUnreadedMessage(ctx context.Context, question_id uint, user_id uint) (int, error)
}
