package persegi

import "errors"

var ErrSisiTidakBolehMinus = errors.New("Something errors")

func HitungLuasPersegi(sisi int) (int, error) {
	if sisi <= 0 {
		return 0, ErrSisiTidakBolehMinus
	}
	return sisi * sisi, nil
}
