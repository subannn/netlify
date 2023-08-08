package minioFunctions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
type pathNames struct {
	path string 
	AnsBucketName string 
	AnsPathName string 
}
func TestParsePath(t *testing.T) {
	tescases := []pathNames {
		pathNames {"/files/htmlcss", "files", "htmlcss/index.html"},
		pathNames {"/files", "files", "index.html"},
	}
	for _, test := range tescases {
		bucketName, pathName := parsePath(test.path)
		assert.Equal(t, bucketName, test.AnsBucketName, "Incorrect bucket name")
		assert.Equal(t, pathName, test.AnsPathName, "Incorrect path name")
	}
	
}