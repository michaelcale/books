package stackoverflow

import "github.com/essentialbooks/books/pkg/common"

func LoadTopics(path string) ([]Topic, error) {
	var res []Topic
	err := common.JSONDecodeGzipped(path, &res)
	return res, err
}
func LoadExamples(path string) ([]*Example, error) {
	var res []*Example
	err := common.JSONDecodeGzipped(path, &res)
	return res, err
}

func LoadTopicHistories(path string) ([]TopicHistory, error) {
	var res []TopicHistory
	err := common.JSONDecodeGzipped(path, &res)
	return res, err
}

func LoadContibutors(path string) ([]*Contributor, error) {
	var res []*Contributor
	err := common.JSONDecodeGzipped(path, &res)
	return res, err
}

func LoadDocTags(path string) ([]DocTag, error) {
	var res []DocTag
	err := common.JSONDecodeGzipped(path, &res)
	return res, err
}
