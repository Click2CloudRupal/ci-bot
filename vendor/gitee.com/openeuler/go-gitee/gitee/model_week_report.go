/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 新建周报
type WeekReport struct {
	Id          int32     `json:"id,omitempty"`
	Content     string    `json:"content,omitempty"`
	ContentHtml string    `json:"content_html,omitempty"`
	Year        string    `json:"year,omitempty"`
	Month       string    `json:"month,omitempty"`
	WeekIndex   string    `json:"week_index,omitempty"`
	WeekBegin   string    `json:"week_begin,omitempty"`
	WeekEnd     string    `json:"week_end,omitempty"`
	CreatedAt   string    `json:"created_at,omitempty"`
	UpdatedAt   string    `json:"updated_at,omitempty"`
	User        *UserMini `json:"user,omitempty"`
}
