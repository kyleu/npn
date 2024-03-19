package controller

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/controller/cutil"
	"github.com/kyleu/npn/app/request/imprt"
)

func ImportUpload(w http.ResponseWriter, r *http.Request) {
	Act("import.upload", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		err := r.ParseMultipartForm(cutil.MaxBodySize)
		if err != nil {
			return "", err
		}
		mpfrm := r.MultipartForm
		fileHeaders, ok := mpfrm.File["f"]
		if !ok {
			return "", errors.New("no file uploaded")
		}
		if len(fileHeaders) != 1 {
			return "", errors.New("invalid file uploads")
		}

		files := make([]imprt.File, 0)
		for _, fileHeader := range fileHeaders {
			ct, ok := fileHeader.Header["Content-Type"]
			if !ok || len(ct) == 0 {
				ct = []string{"text/plain"}
			}
			files = append(files, imprt.File{
				Filename:    fileHeader.Filename,
				Size:        fileHeader.Size,
				ContentType: ct[0],
			})
		}
		err = as.Services.NPN.Import.Create("import", files)
		if err != nil {
			return "", err
		}
		for _, fileHeader := range fileHeaders {
			file, err := fileHeader.Open()
			if err != nil {
				return "", err
			}
			defer func() { _ = file.Close() }()
			err = as.Services.NPN.Import.WriteFile("import", fileHeader.Filename, file)
			if err != nil {
				return "", errors.Wrapf(err, "unable to write import file [%s]", fileHeader.Filename)
			}
		}
		err = as.Services.NPN.Import.Create("import", files)
		if err != nil {
			return "", err
		}

		msg := fmt.Sprintf("Uploaded [%d] files", len(fileHeaders))
		redir := "/TODO"
		return FlashAndRedir(true, msg, redir, w, ps)
	})
}
