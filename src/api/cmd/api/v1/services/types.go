package services


type User struct {
    ID int `json:"id", bson:"id"`
    AvatarURL string `json:"avatar_url", bson:"avatar_url"`
    Username string `json:"username", bson:"username"`
    ProviderID float64 `json:"provider_id" bson:"provider_id"`
    Provider string `json:"provider" bson:"provider"`
}

type NewUser struct {
    AvatarURL string `json:"avatar_url", bson:"avatar_url"`
    Username string `json:"username", bson:"username"`
    ProviderID float64 `json:"provider_id" bson:"provider_id"`
    Provider string `json:"provider" bson:"provider"`
}

type Arc struct {
    ID           int    `json:"id"`
    Title        string `json:"title"`
    Description  string `json:"description"`
    Episodes     interface{}    `json:"episodes"`
    Image        string `json:"image"`
    Time         interface{} `json:"time"`
    CreatedDate  string `json:"created_date"`
    FinishedDate string `json:"finished_date"`
    Rate         interface{} `json:"rate"`
}