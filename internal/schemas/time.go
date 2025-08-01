package schemas

import (
	"fmt"
	"strings"
	"time"
)

const (
	layoutDate = "02-01-2006"
	layoutTime = "15:04"
)

type CustomDate time.Time
type CustomTime time.Time

func ToCustomDate(s string) (CustomDate, error) {
	t, err := time.Parse(layoutDate, s)
	if err != nil {
		return CustomDate{}, fmt.Errorf("%v (esperado formato dd-mm-yyyy)", err)
	}
	return CustomDate(t), nil
}

func ToCustomTime(s string) (CustomTime, error) {
	t, err := time.Parse(layoutTime, s)
	if err != nil {
		return CustomTime{}, fmt.Errorf("%v (esperado formato hh:mm)", err)
	}
	return CustomTime(t), nil
}


func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(layoutDate, s)
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(cd).Format(layoutDate) + `"`), nil
}

func (cd CustomDate) String() string {
	return time.Time(cd).Format(layoutDate)
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(layoutTime, s)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(ct).Format(layoutTime) + `"`), nil
}

func (ct CustomTime) String() string {
	return time.Time(ct).Format(layoutTime)
}

func (cd CustomDate) ToTime() time.Time {
	return time.Time(cd)
}

func (ct CustomTime) ToTime() time.Time {
	return time.Time(ct)
}



