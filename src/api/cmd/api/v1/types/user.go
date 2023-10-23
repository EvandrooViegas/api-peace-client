package types
import "github.com/golang-jwt/jwt/v4"


type User struct {
	ID         string  `json:"id", bson:"_id"`
	AvatarURL  string  `json:"avatar_url", bson:"avatar_url"`
	Username   string  `json:"username", bson:"username"`
	ProviderID float64 `json:"provider_id" bson:"provider_id"`
	Provider   string  `json:"provider" bson:"provider"`
}

type NewUser struct {
	AvatarURL  string  `json:"avatar_url", bson:"avatar_url"`
	Username   string  `json:"username", bson:"username"`
	ProviderID float64 `json:"provider_id" bson:"provider_id"`
	Provider   string  `json:"provider" bson:"provider"`
}

type UserTokenClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}