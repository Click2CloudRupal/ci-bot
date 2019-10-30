package cibot

import (
	"strings"

	"github.com/antihax/optional"

	"gitee.com/openeuler/go-gitee/gitee"
	"github.com/golang/glog"
)

// Close closes pr or issue
func (s *Server) Close(event *gitee.NoteEvent) error {
	// handle PullRequest
	if *event.NoteableType == "PullRequest" {
		/* when gitee support to close pr by api
		// handle open
		if event.PullRequest.State == "open" {
			// get basic params
			comment := event.Comment.Body
			owner := event.Repository.Namespace
			repo := event.Repository.Name
			prAuthor := event.PullRequest.User.Login
			prNumber := event.PullRequest.Number
			commentAuthor := event.Comment.User.Login
			glog.Infof("close started. comment: %s prAuthor: %s commentAuthor: %s owner: %s repo: %s number: %d",
				comment, prAuthor, commentAuthor, owner, repo, prNumber)

			// check if current author has write permission
			localVarOptionals := &gitee.GetV5ReposOwnerRepoCollaboratorsUsernamePermissionOpts{}
			localVarOptionals.AccessToken = optional.NewString(s.Config.GiteeToken)
			// get permission
			permission, _, err := s.GiteeClient.RepositoriesApi.GetV5ReposOwnerRepoCollaboratorsUsernamePermission(
				s.Context, owner, repo, commentAuthor, localVarOptionals)
			if err != nil {
				glog.Errorf("unable to get comment author permission: %v", err)
				return err
			}
			// permission: admin, write, read, none
			if permission.Permission == "admin" || permission.Permission == "write" || prAuthor == commentAuthor {
				//  pr author or permission: admin, write
				body := gitee.PullRequestUpdateParam{}
				body.AccessToken = s.Config.GiteeToken
				body.State = "closed"
				glog.Infof("invoke api to close: %d", prNumber)

				// patch state
				_, response, err := s.GiteeClient.PullRequestsApi.PatchV5ReposOwnerRepoPullsNumber(s.Context, owner, repo, prNumber, body)
				if err != nil {
					if response.StatusCode == 400 {
						glog.Infof("close successfully with status code %d: %d", response.StatusCode, prNumber)
					} else {
						glog.Errorf("unable to close: %d err: %v", prNumber, err)
						return err
					}
				} else {
					glog.Infof("close successfully: %v", prNumber)
				}
			}
		}*/
	} else if *event.NoteableType == "Issue" {
		// handle open
		if event.Issue.State == "open" {
			// get basic informations
			comment := event.Comment.Body
			owner := event.Repository.Namespace
			repo := event.Repository.Name
			issueNumber := event.Issue.Number
			issueAuthor := event.Issue.User.Login
			commentAuthor := event.Comment.User.Login
			glog.Infof("close started. comment: %s owner: %s repo: %s issueNumber: %s issueAuthor: %s commentAuthor: %s",
				comment, owner, repo, issueNumber, issueAuthor, commentAuthor)

			// check if current author has write permission
			localVarOptionals := &gitee.GetV5ReposOwnerRepoCollaboratorsUsernamePermissionOpts{}
			localVarOptionals.AccessToken = optional.NewString(s.Config.GiteeToken)
			// get permission
			permission, _, err := s.GiteeClient.RepositoriesApi.GetV5ReposOwnerRepoCollaboratorsUsernamePermission(
				s.Context, owner, repo, commentAuthor, localVarOptionals)
			if err != nil {
				glog.Errorf("unable to get comment author permission: %v", err)
				return err
			}
			// permission: admin, write, read, none
			if permission.Permission == "admin" || permission.Permission == "write" || issueAuthor == commentAuthor {
				//  issue author or permission: admin, write
				body := gitee.IssueUpdateParam{}
				body.Repo = repo
				body.AccessToken = s.Config.GiteeToken
				body.State = "closed"
				// build label string
				var strLabel string
				for _, l := range event.Issue.Labels {
					strLabel += l.Name + ","
				}
				strLabel = strings.TrimRight(strLabel, ",")
				if strLabel == "" {
					strLabel = ","
				}
				body.Labels = strLabel
				glog.Infof("invoke api to close: %s", issueNumber)

				// patch state
				_, response, err := s.GiteeClient.IssuesApi.PatchV5ReposOwnerIssuesNumber(s.Context, owner, issueNumber, body)
				if err != nil {
					if response.StatusCode == 400 {
						glog.Infof("close successfully with status code %d: %s", response.StatusCode, issueNumber)
					} else {
						glog.Errorf("unable to close: %s err: %v", issueNumber, err)
						return err
					}
				} else {
					glog.Infof("close successfully: %v", issueNumber)
				}
			}
		}
	}
	return nil
}
