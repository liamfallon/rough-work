package types

import (
	"errors"
	"strings"
)

var IssueMap map[string]*GithubRecord

type GithubCommentRecord struct {
	commentUser      string
	commentCreatedAt string
	commentUpdatedAt string
	commentBody      string
}

type GithubRecord struct {
	issueUrl             string
	issueTitle           string
	issueUserHtmlUrl     string
	issueLabels          string
	issueState           string
	issueComments        string
	issueCreatedAt       string
	issueUpdatedAt       string
	issueBody            string
	comments             []GithubCommentRecord
	issueAssigneeHtmlUrl string
}

func NewGithubRecord(csvRecord []string) *GithubRecord {
	ghr := new(GithubRecord)

	ghr.issueUrl = csvRecord[0]
	ghr.issueTitle = csvRecord[9]
	ghr.issueUserHtmlUrl = csvRecord[16]
	ghr.issueLabels = csvRecord[28]
	ghr.issueState = csvRecord[29]
	ghr.issueComments = csvRecord[34]
	ghr.issueCreatedAt = csvRecord[35]
	ghr.issueUpdatedAt = csvRecord[36]
	ghr.issueBody = csvRecord[46]

	ghr.comments = make([]GithubCommentRecord, 1)
	ghr.comments[0].commentUser = csvRecord[60]
	ghr.comments[0].commentCreatedAt = csvRecord[61]
	ghr.comments[0].commentUpdatedAt = csvRecord[62]
	ghr.comments[0].commentBody = csvRecord[63]

	ghr.issueAssigneeHtmlUrl = csvRecord[70]

	return ghr
}

func (ghr *GithubRecord) IssueUrl() string {
	return ghr.issueUrl
}

func (ghr *GithubRecord) IsBlank() bool {
	if len(ghr.issueTitle) == 0 {
		return true
	} else {
		return false
	}
}

func (ghr *GithubRecord) Merge(inGhr *GithubRecord) error {

	if ghr.issueUrl != inGhr.issueUrl {
		err := errors.New("issueUrl not equal")
		return err
	}
	if ghr.issueTitle != inGhr.issueTitle {
		err := errors.New("issueTitle not equal")
		return err
	}
	if ghr.issueUserHtmlUrl != inGhr.issueUserHtmlUrl {
		err := errors.New("issueUserHtmlUrl not equal")
		return err
	}
	if ghr.issueLabels != inGhr.issueLabels {
		err := errors.New("issueLabels not equal")
		return err
	}
	if ghr.issueState != inGhr.issueState {
		err := errors.New("issueState not equal")
		return err
	}
	if ghr.issueComments != inGhr.issueComments {
		err := errors.New("issueComments not equal")
		return err
	}
	if ghr.issueCreatedAt != inGhr.issueCreatedAt {
		err := errors.New("issueCreatedAt not equal")
		return err
	}
	if ghr.issueUpdatedAt != inGhr.issueUpdatedAt {
		err := errors.New("issueUpdatedAt not equal")
		return err
	}
	if ghr.issueBody != inGhr.issueBody {
		err := errors.New("issueBody not equal")
		return err
	}
	if ghr.issueAssigneeHtmlUrl != inGhr.issueAssigneeHtmlUrl {
		err := errors.New("issueAssigneeHtmlUrl not equal")
		return err
	}

	ghr.comments = append(ghr.comments, inGhr.comments...)

	return nil
}

func (ghr *GithubRecord) OutStringArray() []string {
	outSlice := make([]string, 3)

	outSlice[0] = ghr.issueTitle
	outSlice[1] = ghr.GetBody()
	outSlice[2] = ghr.GetLabels()

	return outSlice
}

func (ghr *GithubRecord) GetLabels() string {
	splitLabels := strings.SplitAfter(strings.Trim(ghr.issueLabels, "[]"), "}")

	var labelString = ""

	for _, label := range splitLabels {
		splitLabel := strings.Split(strings.Trim(strings.TrimLeft(label, ",'"), "{}"), ",")

		if len(splitLabel) > 0 {
			for _, labelField := range splitLabel {
				nvp := strings.Split(labelField, ":")
				if nvp[0] == "\"name\"" {
					labelString = labelString + strings.Trim(nvp[1], "\"") + ","
				}
			}
		}
	}

	labelString = strings.TrimRight(labelString, ",")
	return labelString
}

func (ghr *GithubRecord) GetBody() string {
	return "Original issue URL: " + strings.ReplaceAll(strings.ReplaceAll(ghr.issueUrl, "api.github.com", "github.com"), "repos/kptdev", "kptdev") + "\n" +
		"Original issue user: " + ghr.issueUserHtmlUrl + "\n" +
		"Original issue state: " + ghr.issueState + "\n" +
		"Original issue created at: " + ghr.issueCreatedAt + "\n" +
		"Original issue last updated at: " + ghr.issueUpdatedAt + "\n" +
		"Original issue body: " + ghr.issueBody + "\n\n" +
		ghr.GetComments()
}

func (ghr *GithubRecord) GetState() string {
	return ghr.issueState
}

func (ghr *GithubRecord) GetComments() string {
	var commentString = ""

	for _, comment := range ghr.comments {
		if len(comment.commentBody) > 0 {
			commentString = commentString +
				"Comment user: https://github.com/" + comment.commentUser + "\n" +
				"Comment created at: " + comment.commentCreatedAt + "\n" +
				"Comment last updated at: " + comment.commentUpdatedAt + "\n" +
				"Comment body: " + comment.commentBody + "\n\n"
		}
	}

	if len(commentString) == 0 {
		return "Original issue comments: None\n"
	} else {
		return "Original issue comments: \n" + commentString
	}
}
