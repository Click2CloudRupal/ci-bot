/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

// 搜索代码片段
type Code struct {
	Url         string `json:"url,omitempty"`
	ForksUrl    string `json:"forks_url,omitempty"`
	CommitsUrl  string `json:"commits_url,omitempty"`
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Public      string `json:"public,omitempty"`
	Owner       string `json:"owner,omitempty"`
	User        string `json:"user,omitempty"`
	Files       string `json:"files,omitempty"`
	Truncated   string `json:"truncated,omitempty"`
	HtmlUrl     string `json:"html_url,omitempty"`
	Comments    string `json:"comments,omitempty"`
	CommentsUrl string `json:"comments_url,omitempty"`
	GitPullUrl  string `json:"git_pull_url,omitempty"`
	GitPushUrl  string `json:"git_push_url,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}
