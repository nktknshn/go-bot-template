package bot

import "go.uber.org/zap"

type options struct {
	logger        *zap.Logger
	debugTelegram bool

	adminID         int64
	restrictToAdmin bool

	forwardTo      int64
	useRateLimiter bool
	sessionFile    string

	includeText    bool
	includeURL     bool
	includeBotName bool

	limitPerDay  int
	limitPending int
}

type option func(*options)

func WithLogger(logger *zap.Logger) option {
	return func(opts *options) {
		opts.logger = logger
	}
}

func WithDebugTelegram(debug bool) option {
	return func(opts *options) {
		opts.debugTelegram = debug
	}
}

func WithAdmin(adminID int64, restrict bool) option {
	return func(opts *options) {
		opts.adminID = adminID
		opts.restrictToAdmin = restrict
	}
}

func WithForwardTo(forwardTo int64) option {
	return func(opts *options) {
		opts.forwardTo = forwardTo
	}
}

func WithRateLimiter(useRateLimiter bool) option {
	return func(opts *options) {
		opts.useRateLimiter = useRateLimiter
	}
}

func WithSessionFile(sessionFile string) option {
	return func(opts *options) {
		opts.sessionFile = sessionFile
	}
}

// func WithPostSettings(includeText, includeURL, includeBotName bool) option {
// 	return func(opts *options) {
// 		opts.includeText = includeText
// 		opts.includeURL = includeURL
// 		opts.includeBotName = includeBotName
// 	}
// }
