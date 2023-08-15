package minioFunctions

import (
	"testing"
	// "fmt"

	"github.com/stretchr/testify/assert"
)
type pathNames struct {
	path string 
	AnsPathName string 
}
func TestParsePath(t *testing.T) {
	tescases := []pathNames {
		pathNames {"/files/htmlcss.com", "files/htmlcss.com"},
		pathNames {"/files/", "files/index.html"},
		pathNames {"/", "index.html"},
		pathNames {"/files/aboba.html", "files/aboba.html"},
		pathNames {"/qwerty", "qwerty/index.html"},
		pathNames {"", "index.html"},
		pathNames {"/htmlcss/files/index.html", "htmlcss/files/index.html"},
		pathNames {"/files/../../../../../../etc/shadow", "htmlcss/files/index.html"},
	}
	for _, test := range tescases {
		pathName := ParsePath(test.path)
		assert.Equal(t, pathName, test.AnsPathName, pathName)
	}
}