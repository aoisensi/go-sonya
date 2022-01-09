package sonya

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type WithAvatar []byte

func (w WithAvatar) optModifyCurrentUser(v map[string]interface{}) {
	buf := bytes.NewBuffer(w)
	_, format, err := image.DecodeConfig(buf)
	if err != nil {
		panic(err)
	}
	if format != "jpeg" && format != "png" && format != "gif" {
		panic(errors.New("unsupported image format"))
	}
	v["avatar"] = "data:image/" + format + ";base64," +
		base64.StdEncoding.EncodeToString(w)
}

type WithUsername string

func (w WithUsername) optModifyCurrentUser(v map[string]interface{}) {
	v["username"] = string(w)
}
