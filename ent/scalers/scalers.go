package scalers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type JSON map[string]interface{}

func MarshalUrl(u url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(u.String())))
	})
}

func UnmarshalUrl(v interface{}) (url.URL, error) {
	switch v := v.(type) {
	case string:
		u, err := url.Parse(v)
		if err != nil {
			return url.URL{}, fmt.Errorf("parse url: %+v", err)
		}
		return *u, nil
	case url.URL:
		return v, nil
	default:
		return url.URL{}, fmt.Errorf("%T is not a string or url.URL", v)
	}
}

func MarshalId(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(id.String()))
	})
}

func UnmarshalId(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		res, err := uuid.Parse(v)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("error parsing uuid: %+v", err)
		}
		return res, nil
	case uuid.UUID:
		return v, nil
	default:
		return uuid.UUID{}, fmt.Errorf("%T is not a string", v)
	}
}

func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(t.Format(time.RFC3339))))
	})
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		res, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return time.Time{}, fmt.Errorf("parse time: %w", err)
		}
		return res, nil
	case time.Time:
		return v, nil
	default:
		return time.Time{}, fmt.Errorf("%T is not a string", v)
	}
}

func MarshalDuration(d time.Duration) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Quote(d.String())))
	})
}

func UnmarshalDuration(v interface{}) (time.Duration, error) {
	switch v := v.(type) {
	case string:
		d, err := time.ParseDuration(v)
		if err != nil {
			return time.Duration(0), err
		}

		return d, nil
	case time.Duration:
		return v, nil
	default:
		return time.Duration(0), fmt.Errorf("%T is not a string", v)
	}
}

func MarshalJSON(b JSON) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		byteData, _ := json.Marshal(b)
		_, _ = w.Write(byteData)
	})
}

func UnmarshalJSON(v interface{}) (JSON, error) {
	byteData, err := json.Marshal(v)
	if err != nil {
		return JSON{}, fmt.Errorf("unmarshal json error: %w", err)
	}
	tmp := make(map[string]interface{})
	err = json.Unmarshal(byteData, &tmp)
	if err != nil {
		return JSON{}, fmt.Errorf("unmarshal json error: %w", err)
	}
	return tmp, nil
}

func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(fmt.Sprintf("%d", i)))
	})
}

func UnmarshalInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	default:
		return 0, fmt.Errorf("%T is not a string", v)
	}
}
