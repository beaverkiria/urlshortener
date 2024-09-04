package app

import "errors"

var storage = map[string]string{}

func SaveRedirect(fromUrl string, toUrl string) error {
	storage[toUrl] = fromUrl
	return nil
}

func GetRedirect(url string) (string, error) {
	redirect, ok := storage[url]
	if !ok {
		return "", errors.New("redirect not found")
	}
	return redirect, nil
}
