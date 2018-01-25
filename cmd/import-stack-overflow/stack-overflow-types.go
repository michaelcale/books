package main

/*
Dates are in format:
/Date(1447119317900-0500)/
*/

/*
contributordeletionreasons.json
	- Id
	- Name
	- Description

{
	"Id": 1,
	"Name": "Upgrade",
	"Description": "Contribution is ending because the user is now a Major contributor"
}
*/

// ContributroDeleteRasons represents data in contributordeletionreasons.json
type ContributroDeleteRasons struct {
	Id          int
	Name        string
	Description string
}

/*
contributors.json
	- Id
	- DocTopicId
	- DocExampleId
	- UserId
	- DocContributorTypeId
	- CreationDate
	- DeletionDate
	- DocContributorDeletionReasonId

{
	"Id": 1,
	"DocTopicId": 1,
	"UserId": 80572,
	"DocContributorTypeId": 2,
	"CreationDate": "\/Date(1446697142040-0500)\/"
}
*/

// Contributor represents data in contributors.json
type Contributor struct {
	Id                             int
	DocTopicId                     int
	DocExampleId                   int
	UserId                         int
	DocContributorTypeId           int
	CreationDate                   string
	DeletionDate                   string
	DocContributorDeletionReasonId int
}

func loadContibutors(path string) ([]*Contributor, error) {
	var res []*Contributor
	err := jsonDecodeGzipped(path, &res)
	return res, err
}

/*
contributortypes.json
	- Id
	- Name
	- Description
{
	"Id": 1,
	"Name": "Minor",
	"Description": "Has contributed in a small amount."
},
{
	"Id": 2,
	"Name": "Major",
	"Description": "Has contributed a lot."
}
*/

// ContributorTypes represents data in contributortypes.json
type ContributorTypes struct {
	Id          int
	Name        string
	Description string
}

/*
doctags.json
	- Id
	- Tag
	- Title
	- CreationDate
	- HelloWorldDocTopicId
	- TopicCount

{
	"Id": 3,
	"Tag": ".net",
	"Title": ".NET Framework",
	"CreationDate": "\/Date(1447273894907-0500)\/",
	"HelloWorldDocTopicId": 14,
	"TopicCount": 59
}
*/

// DocTag represents data in doctags.json
type DocTag struct {
	Id    int
	Tag   string
	Title string
	//CreationDate         string
	HelloWorldDocTopicId int
	TopicCount           int
	ExampleCount         int
}

func loadDocTags(path string) ([]DocTag, error) {
	var res []DocTag
	err := jsonDecodeGzipped(path, &res)
	return res, err
}

/*
doctagversions.json
	- Id
	- DocTagId
	- Name
	- GroupName
	- CreationDate
	- ReleaseDate
	- LastEditDate
	- LastEditUserId
	- LastEditUserDisplayName: populated if a user has been removed and no longer referenced by user id

  {
    "Id": 206,
    "DocTagId": 13,
    "Name": "1.6.0",
    "CreationDate": "\/Date(1458812411197-0400)\/",
    "LastEditUserId": 17373,
    "ReleaseDate": "\/Date(1455685200000-0500)\/"
  },
*/

// DocTagVersion represents data in doctagversions.json
type DocTagVersion struct {
	Id                      int
	DocTagId                int
	Name                    string
	GroupName               string
	CreationDate            string
	ReleaseDate             string
	LastEditDate            string
	LastEditUserId          int
	LastEditUserDisplayName string
}

/*
examples.json
	- Id
	- DocTopicId
	- Title
	- CreationDate
	- LastEditDate
	- Score
	- ContributorCount
	- BodyHtml
	- IsPinned
	- BodyMarkdown

{
	"Id": 1,
	"DocTopicId": 1,
	"Title": "Basic Usage",
	"CreationDate": "\/Date(1446697142040-0500)\/",
	"LastEditDate": "\/Date(1469351669667-0400)\/",
	"Score": 6,
	"ContributorCount": 2,
	"BodyHtml": "<pre><code>using StackExchange.Redis;\r\n\r\n// ...\r\n\r\n// connect to the server\r\nConnectionMultiplexer connection = ConnectionMultiplexer.Connect(&quot;localhost&$
	"BodyMarkdown": "    using StackExchange.Redis;\n\n    // ...\n\n    // connect to the server\n    ConnectionMultiplexer connection = ConnectionMultiplexer.Connect(\"localhost\");\n$
	"IsPinned": false
},
*/

// Example represents data in examples.json
type Example struct {
	Id         int
	DocTopicId int
	Title      string
	//CreationDate     string
	//LastEditDate     string
	Score int
	//ContributorCount int
	BodyHtml     string
	IsPinned     bool
	BodyMarkdown string
}

func loadExamples(path string) ([]*Example, error) {
	var res []*Example
	err := jsonDecodeGzipped(path, &res)
	return res, err
}

/*
topichistories.json
	- Id
	- DocTopicHistoryTypeId
	- DocTagId
	- DocTopicId
	- DocExampleId
	- CreationDate
	- RevisionNumber
	- CreationUserId
	- CreationUserDisplayName
	- Comment
	- Text

	{
	"Id": 1,
	"DocTopicHistoryTypeId": 1,
	"DocTagId": 1,
	"DocTopicId": 1,
	"CreationDate": "\/Date(1446695335230-0500)\/",
	"RevisionNumber": 1,
	"CreationUserId": -1,
	"Text": "Hello World"
}
*/

// TopicHistory represents data in topichistories.json
type TopicHistory struct {
	Id                    int
	DocTopicHistoryTypeId int
	DocTagId              int
	DocTopicId            int
	DocExampleId          int
	//CreationDate            string
	RevisionNumber          int
	CreationUserId          int
	CreationUserDisplayName string
	Comment                 string
	Text                    string
}

func loadTopicHistories(path string) ([]TopicHistory, error) {
	var res []TopicHistory
	err := jsonDecodeGzipped(path, &res)
	return res, err
}

/*
topichistorytypes.json
	- Id
	- Name
	- Description

{
	"Id": 1,
	"Name": "InitialTitle",
	"Description": "The first title a topic was posted with."
}
*/

// TopicHistoryTypes represents data in topichistorytypes.json
type TopicHistoryTypes struct {
	Id          int
	Name        string
	Description string
}

/*
topics.json
	- Id
	- DocTagId
	- IsHelloWorldTopic
	- Title
	- CreationDate
	- ViewCount
	- LastEditDate
	- ContributorCount
	- IntroductionHtml
	- SyntaxHtml
	- ParametersHtml
	- RemarksHtml
	- HelloWorldVersionsHtml
	- VersionsJson
	- ExampleCount
	- ExampleScore
	- LastEditUserId
	- LastEditUserDisplayName: populated if a user has been removed and no longer referenced by user id
	- IntroductionMarkdown
	- SyntaxMarkdown
	- ParametersMarkdown
	- RemarksMarkdown

{
	"Id": 1,
	"DocTagId": 1,
	"IsHelloWorldTopic": true,
	"Title": "Getting started with StackExchange.Redis",
	"CreationDate": "\/Date(1446695335230-0500)\/",
	"ViewCount": 266,
	"LastEditDate": "\/Date(1492733341057-0400)\/",
	"LastEditUserId": 3331861,
	"ContributorCount": 5,
	"ExampleCount": 3,
	"ExampleScore": 6,
	"SyntaxHtml": "",
	"ParametersHtml": "",
	"RemarksHtml": "<a class=\"remarks-subsection-anchor\" name=\"remarks-installing-0\"></a>\r\n<h3>Installing</h3>\r\n<p>Binaries for StackExchange.Redis are <a href=\"https://www.nug$
	"IntroductionMarkdown": "",
	"SyntaxMarkdown": "",
	"ParametersMarkdown": "",
	"RemarksMarkdown": "### Installing\n\nBinaries for StackExchange.Redis are [available on Nuget][1], and the source is [available on Github][2].\n\n  [1]: https://www.nuget.org/packa$
	"HelloWorldVersionsHtml": "<table><thead><tr><th>Version</th><th>Release Date</th></tr></thead><tbody><tr><td>1.0.187</td><td>2014-03-18</td></tr></tbody></table>\r\n"
}
*/

// Topic represents data in topics.json
type Topic struct {
	Id                int
	DocTagId          int
	IsHelloWorldTopic bool
	Title             string
	//CreationDate      string
	//ViewCount int
	//LastEditDate      string
	//ContributorCount int
	IntroductionHtml       string
	SyntaxHtml             string
	ParametersHtml         string
	RemarksHtml            string
	HelloWorldVersionsHtml string
	VersionsJson           string
	ExampleCount           int
	ExampleScore           int
	//LastEditUserId          int
	//LastEditUserDisplayName string
	IntroductionMarkdown string
	SyntaxMarkdown       string
	ParametersMarkdown   string
	RemarksMarkdown      string
}

func loadTopics(path string) ([]Topic, error) {
	var res []Topic
	err := jsonDecodeGzipped(path, &res)
	return res, err
}
