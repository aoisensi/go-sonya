package sonya

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"strconv"
)

type WithAvatar []byte

func (w WithAvatar) optModifyCurrentUser(v map[string]any) {
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

func (w WithUsername) optModifyCurrentUser(v map[string]any) {
	v["username"] = string(w)
}

type WithBefore Snowflake

func (w WithBefore) optGetCurrentUserGuilds(v url.Values) {
	v.Set("before", string(w))
}

type WithAfter Snowflake

func (w WithAfter) optGetCurrentUserGuilds(v url.Values) {
	v.Set("after", string(w))
}

type WithLimit int

func (w WithLimit) optGetCurrentUserGuilds(v url.Values) {
	v.Set("limit", strconv.Itoa(int(w)))
}
