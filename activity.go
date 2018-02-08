package fileinput

import (
	"io/ioutil"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-huawei-fileinput")

// FileInputActivity implements a file input activity.
type FileInputActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity.
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &FileInputActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata.
func (a *FileInputActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval.
func (a *FileInputActivity) Eval(context activity.Context) (done bool, err error)  {
	pathname := context.GetInput("pathname").(string)

	if s, err := ioutil.ReadFile(pathname); err != nil {
		log.Errorf("%s: %s", pathname, err.Error())
		return false, err
	} else {
		context.SetOutput("filecontents", s)
		return true, nil
	}
}
