package date

import "time"

func MustParse(layout, value string) Type {
	v, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}

	return New(v)
}
