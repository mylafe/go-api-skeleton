// model
package user

import (
    "go_api_skeleton/app/models"
)

// User 用户模型
type User struct {
    models.BaseModel

    Name     string `json:"name,omitempty"`
    Email    string `json:"-"`
    Phone    string `json:"-"`
    Password string `json:"-"`

    models.CommonTimestampsField
}
