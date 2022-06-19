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

type WithBefore Snowflake

func (w WithBefore) optGetCurrentUserGuilds(v url.Values) {
	v.Set("before", Snowflake(w).String())
}

type WithAfter Snowflake

func (w WithAfter) optGetCurrentUserGuilds(v url.Values) {
	v.Set("after", Snowflake(w).String())
}

type WithLimit int

func (w WithLimit) optGetCurrentUserGuilds(v url.Values) {
	v.Set("limit", strconv.Itoa(int(w)))
}
