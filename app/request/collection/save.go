package collection

import (
	"github.com/google/uuid"
	"github.com/kyleu/npn/app/lib/filesystem"
	"github.com/kyleu/npn/app/util"
	"github.com/pkg/errors"
	"os"
	"path"
)

func (s *Service) Save(userID *uuid.UUID, originalKey string, newKey string, title string, description string) error {
	originalKey = util.Slugify(originalKey)
	if newKey == "" {
		newKey = "new"
	}
	newKey = util.Slugify(newKey)

	var orig *Collection
	var err error

	if len(originalKey) > 0 {
		orig, err = s.Load(userID, originalKey)
		if err != nil {
			return errors.Wrap(err, "unable to load original collection ["+originalKey+"]")
		}
		if orig != nil && originalKey != newKey {
			o := path.Join(s.files.Root(), s.dirFor(userID), originalKey)
			n := path.Join(s.files.Root(), s.dirFor(userID), newKey)
			err = os.Rename(o, n)
			if err != nil {
				return errors.Wrap(err, "unable to rename original collection ["+originalKey+"] in path ["+o+"]")
			}
		}
	}

	n := &Collection{
		Key:         newKey,
		Title:       title,
		Description: description,
	}

	if orig == nil {
		n.Owner = "system"
	} else {
		n.Owner = orig.Owner
		n.RequestOrder = orig.RequestOrder
	}
	n.Path = newKey

	p := path.Join(s.dirFor(userID), newKey, "collection.json")
	content := util.ToJSONBytes(n, true)
	err = s.files.WriteFile(p, content, filesystem.DefaultMode, true)
	if err != nil {
		return errors.Wrap(err, "unable to save collection ["+newKey+"]")
	}

	return nil
}
