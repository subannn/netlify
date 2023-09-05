package logic

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type pathNames struct {
	path        string
	AnsPathName string
}

func TestParsePath(t *testing.T) {
	tescases := []pathNames{
		{"/files/htmlcss.com", "files/htmlcss.com"},
		{"/files/", "files/index.html"},
		{"/", "index.html"},
		{"/files/aboba.html", "files/aboba.html"},
		{"/qwerty", "qwerty/index.html"},
		{"", "index.html"},
		{"/htmlcss/files/index.html", "htmlcss/files/index.html"},
		{"/files/../../../../../../etc/shadow", "etc/shadow/index.html"},
	}
	for _, test := range tescases {
		pathName := ParsePath(test.path)
		assert.Equal(t, pathName, test.AnsPathName, pathName)
	}
}
