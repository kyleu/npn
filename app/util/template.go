package util

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	defaultPrefix = "{"
	defaultSuffix = "}"
)

// Merges the provided string with the provided template data using "{" and "}" as affixes
func Merge(content string, args ValueMap) (string, error) {
	return mergeVariables(content, args, defaultPrefix, defaultSuffix, 0)
}

// Retruns true if this string contains "{", the default prefix
func MergeNeeded(key string) bool {
	return strings.Contains(key, defaultPrefix)
}

// Calls `Merge`, logging errors to the provided logger, including the provided key
func MergeLog(key string, content string, args ValueMap, logger Logger) string {
	x, err := Merge(content, args)
	if err != nil {
		logger.Warnf("unable to merge [%v] %+v", key, err)
		return content
	}
	return x
}

func mergeVariables(content string, args ValueMap, start string, end string, depth int) (string, error) {
	if depth > 32 {
		return content, errors.Errorf("template recursion error for [%s]", content)
	}
	sIdx := strings.Index(content, start)
	if sIdx > -1 {
		eIdx := strings.Index(content[sIdx:], end)
		if eIdx > -1 {
			orig := content[sIdx : sIdx+eIdx+len(end)]

			n := orig[len(start) : len(orig)-len(end)]
			d := ""
			dIdx := strings.Index(orig, "|")
			if dIdx > -1 {
				n = orig[len(start):dIdx]
				d = orig[dIdx+1 : len(orig)-len(end)]
			}

			o := args.GetStringOpt(n)
			if len(o) == 0 || o == "<nil>" {
				o = d
			}
			if len(o) == 0 || o == "<nil>" {
				o = n
			}
			return mergeVariables(strings.Replace(content, orig, o, 1), args, start, end, depth+1)
		}
	}

	return content, nil
}
