package helpers

import (
	"fmt"
	"time"

	"github.com/tgwaffles/gladis/helpers/time_type"
)

type TimestampOptions struct {
	Format time_type.TimeType
}

func ToTimestamp(time time.Time, options TimestampOptions) string {
	return fmt.Sprintf("<t:%d:%s>", time.Unix(), options.Format)
}
