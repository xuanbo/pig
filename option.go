package pig

// Option 配置App
type Option func(app *App)

// WithHideBanner 隐藏banner
func WithHideBanner(hideBanner bool) Option {
	return func(app *App) {
		app.hideBanner = hideBanner
	}
}

// WithBanner 设置banner
func WithBanner(banner string) Option {
	return func(app *App) {
		app.banner = banner
	}
}
