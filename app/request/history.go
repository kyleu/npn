package request

import (
	"github.com/kyleu/npn/app/lib/filesystem"
	"github.com/kyleu/npn/app/util"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

const shouldSaveHistory = true

func (s *Service) saveHistory(userID *uuid.UUID, coll string, req *Request, p string, msg []byte) error {
	if shouldSaveHistory {
		hp := s.historyPath(userID, coll, req.Key)
		now := time.Now()
		hfn := path.Join(hp, util.TimeToFull(&now)+".json")
		hd := filepath.Dir(hfn)
		err := s.files.CreateDirectory(hd)
		if err != nil {
			return errors.Wrap(err, "unable to create request history directory ["+hd+"]")
		}

		x, _ := os.Stat(p)
		if x == nil {
			err = s.files.WriteFile(hfn, msg, filesystem.DefaultMode, true)
		} else {
			err = s.files.CopyFile(p, hfn)
		}
		if err != nil {
			return errors.Wrap(err, "unable to create request history ["+hp+"]")
		}
	}

	err := s.files.WriteFile(p, msg, filesystem.DefaultMode, true)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}

	return nil
}

func (s *Service) historyPath(userID *uuid.UUID, coll string, key string) string {
	return path.Join(strings.TrimSuffix(s.dirFor(userID, coll), "requests"), "history", "requests", key)
}
