package foundation

const (
	// ApplicationPort Application Server Default port
	ApplicationPort = "8080"
	ErrorMessage    = "coming soon"
	ConferenceURL   = "https://conference2019.laravel.jp"
	HallAddress     = "東京都港区芝浦３丁目４−１　田町グランパーク　プラザ棟３階・４階"
	HallURL         = "http://granpark-c.com/"
	HallNote        = "JR田町駅（徒歩5分）\n地下鉄都営浅草線・三田線　三田駅（徒歩7分）"
	HallName        = "グランパークカンファレンス"
	HallLatitude    = "35.64450085807801"
	HallLongitude   = "139.74727770000004"
	ConferenceDate  = "2019 / 2 / 16"
	ConferenceName  = "Laravel JP Conference"
)

// FindHallNote カンファレンスホール
func FindHallNote() *string {
	s := HallNote
	return &s
}
