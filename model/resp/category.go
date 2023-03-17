// @User CPR
package resp

import "time"

// 前台
type CategoryVo struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	AnimeCount int       `json:"anime_count"` // 动漫数量
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
