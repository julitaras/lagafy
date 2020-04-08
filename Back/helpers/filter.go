package helpers

import "context"

//Filter function
func Filter(ctx context.Context, l int, predicate func(int) bool, appender func(int)) {
	for i := 0; i < l; i++ {
		if predicate(i) {
			appender(i)
		}
	}
}
