package memory

import "math/rand"

func (r *repository) Create(URL string) (string, error) {
	ru := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, shortUrlLen)

	for i := range shortUrlLen {
		str[i] = ru[rand.Intn(runeLen)]
	}

	r.data[string(str)] = URL

	return string(str), nil
}
