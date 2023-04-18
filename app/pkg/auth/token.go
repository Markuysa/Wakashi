package auth

type Bucket string

const (
	AccessTokens  Bucket = "access_token"
	RequestTokens Bucket = "request_token"
)

type TokenRepos interface {
	Save(chatID int64, token string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}
