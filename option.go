package pig

// Option 配置App
type Option func(app *App)

// WithHideBanner 隐藏banner
func WithHideBanner(hideBanner bool) Option {
	return func(app *App) {
		app.hideBanner = hideBanner
	}
}
