package session

import (
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/npn/app/lib/filesystem"
	"github.com/kyleu/npn/app/util"
)

type Service struct {
	multiuser bool
	files     filesystem.FileLoader
	logger    util.Logger
}

func NewService(multiuser bool, f filesystem.FileLoader, logger util.Logger) *Service {
	return &Service{
		multiuser: multiuser,
		files:     f,
		logger:    logger,
	}
}

func (s *Service) List(userID *uuid.UUID) (Sessions, error) {
	jsons := s.files.ListJSON(s.dirFor(userID), nil, false, s.logger)

	if len(jsons) == 0 {
		return Sessions{defaultSession}, nil
	}

	ret := make(Sessions, 0, len(jsons))
	hasDefault := false
	for _, json := range jsons {
		c, err := s.Load(userID, json)
		if err != nil {
			return nil, err
		}
		if c == nil {
			c = &Session{Key: json}
		}
		if json == "_" {
			hasDefault = true
		}
		ret = append(ret, c)
	}
	if !hasDefault {
		ret = append(Sessions{defaultSession}, ret...)
	}
	return ret, nil
}

func (s *Service) Load(userID *uuid.UUID, key string) (*Session, error) {
	if key == "" {
		key = "_"
	}
	p := path.Join(s.dirFor(userID), key+".json")
	exists := s.files.Exists(p)
	if !exists {
		if key == "_" {
			return defaultSession, nil
		}
		return nil, nil
	}
	ret := &Session{}
	filePath := p
	fileExists := s.files.Exists(filePath)
	if fileExists {
		content, err := s.files.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		err = util.FromJSON(content, ret)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read session from ["+filePath+"]")
		}
	}

	return ret.Normalize(key), nil
}

func (s *Service) Counts(userID *uuid.UUID) ([]*Summary, error) {
	l, err := s.List(userID)
	if err != nil {
		return nil, err
	}

	ret := make(Summaries, 0, len(l))
	for _, coll := range l {
		ret = append(ret, coll.ToSummary())
	}
	return ret, nil
}

func (s *Service) Save(userID *uuid.UUID, originalKey string, sess *Session) error {
	originalKey = util.Slugify(originalKey)
	if sess.Key == "" {
		sess.Key = "new"
	}
	slug := util.Slugify(sess.Key)
	if slug != sess.Key {
		s.logger.Debug(fmt.Sprintf("renaming session key from [%v] to [%v]", sess.Key, slug))
		if sess.Title == "" {
			sess.Title = sess.Key
		}
		sess.Key = slug
	}

	shouldDelete := len(originalKey) > 0 && sess.Key != originalKey

	if shouldDelete {
		orig, err := s.Load(userID, sess.Key)
		if err == nil && orig != nil {
			return errors.New("session file already exists with key [" + sess.Key + "]")
		}
	}

	p := s.sessionPath(userID, sess.Key)

	if shouldDelete {
		o := path.Join(s.files.Root(), s.sessionPath(userID, originalKey))
		n := path.Join(s.files.Root(), p)
		err := os.Rename(o, n)
		if err != nil {
			return errors.Wrap(err, "unable to rename original request ["+originalKey+"] in path ["+o+"]")
		}
	}

	msg := util.ToJSONBytes(sess, true)
	err := s.files.WriteFile(p, msg, filesystem.DefaultMode, true)
	if err != nil {
		return errors.Wrap(err, "unable to write file")
	}

	return nil
}

func (s *Service) dirFor(userID *uuid.UUID) string {
	if (!s.multiuser) || userID == nil || *userID == util.SystemUserID {
		return "sessions"
	}
	return path.Join("users", userID.String(), "sessions")
}

func (s *Service) sessionPath(userID *uuid.UUID, key string) string {
	return path.Join(s.dirFor(userID), key+".json")
}

func (s *Service) Delete(userID *uuid.UUID, key string) error {
	p := path.Join(s.dirFor(userID), key+".json")
	return s.files.Remove(p, s.logger)
}
