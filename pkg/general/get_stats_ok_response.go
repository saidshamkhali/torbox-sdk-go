package general

type GetStatsOkResponse struct {
	Data   *GetStatsOkResponseData `json:"data,omitempty"`
	Detail *string                 `json:"detail,omitempty"`
}

func (g *GetStatsOkResponse) SetData(data GetStatsOkResponseData) {
	g.Data = &data
}

func (g *GetStatsOkResponse) GetData() *GetStatsOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetStatsOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetStatsOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

type GetStatsOkResponseData struct {
	ActiveTorrents        *float64 `json:"active_torrents,omitempty"`
	ActiveUsenetDownloads *float64 `json:"active_usenet_downloads,omitempty"`
	ActiveWebDownloads    *float64 `json:"active_web_downloads,omitempty"`
	TotalBytesDownloaded  *float64 `json:"total_bytes_downloaded,omitempty"`
	TotalBytesUploaded    *float64 `json:"total_bytes_uploaded,omitempty"`
	TotalDownloads        *float64 `json:"total_downloads,omitempty"`
	TotalServers          *float64 `json:"total_servers,omitempty"`
	TotalTorrentDownloads *float64 `json:"total_torrent_downloads,omitempty"`
	TotalUsenetDownloads  *float64 `json:"total_usenet_downloads,omitempty"`
	TotalUsers            *float64 `json:"total_users,omitempty"`
	TotalWebDownloads     *float64 `json:"total_web_downloads,omitempty"`
}

func (g *GetStatsOkResponseData) SetActiveTorrents(activeTorrents float64) {
	g.ActiveTorrents = &activeTorrents
}

func (g *GetStatsOkResponseData) GetActiveTorrents() *float64 {
	if g == nil {
		return nil
	}
	return g.ActiveTorrents
}

func (g *GetStatsOkResponseData) SetActiveUsenetDownloads(activeUsenetDownloads float64) {
	g.ActiveUsenetDownloads = &activeUsenetDownloads
}

func (g *GetStatsOkResponseData) GetActiveUsenetDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.ActiveUsenetDownloads
}

func (g *GetStatsOkResponseData) SetActiveWebDownloads(activeWebDownloads float64) {
	g.ActiveWebDownloads = &activeWebDownloads
}

func (g *GetStatsOkResponseData) GetActiveWebDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.ActiveWebDownloads
}

func (g *GetStatsOkResponseData) SetTotalBytesDownloaded(totalBytesDownloaded float64) {
	g.TotalBytesDownloaded = &totalBytesDownloaded
}

func (g *GetStatsOkResponseData) GetTotalBytesDownloaded() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalBytesDownloaded
}

func (g *GetStatsOkResponseData) SetTotalBytesUploaded(totalBytesUploaded float64) {
	g.TotalBytesUploaded = &totalBytesUploaded
}

func (g *GetStatsOkResponseData) GetTotalBytesUploaded() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalBytesUploaded
}

func (g *GetStatsOkResponseData) SetTotalDownloads(totalDownloads float64) {
	g.TotalDownloads = &totalDownloads
}

func (g *GetStatsOkResponseData) GetTotalDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalDownloads
}

func (g *GetStatsOkResponseData) SetTotalServers(totalServers float64) {
	g.TotalServers = &totalServers
}

func (g *GetStatsOkResponseData) GetTotalServers() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalServers
}

func (g *GetStatsOkResponseData) SetTotalTorrentDownloads(totalTorrentDownloads float64) {
	g.TotalTorrentDownloads = &totalTorrentDownloads
}

func (g *GetStatsOkResponseData) GetTotalTorrentDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalTorrentDownloads
}

func (g *GetStatsOkResponseData) SetTotalUsenetDownloads(totalUsenetDownloads float64) {
	g.TotalUsenetDownloads = &totalUsenetDownloads
}

func (g *GetStatsOkResponseData) GetTotalUsenetDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalUsenetDownloads
}

func (g *GetStatsOkResponseData) SetTotalUsers(totalUsers float64) {
	g.TotalUsers = &totalUsers
}

func (g *GetStatsOkResponseData) GetTotalUsers() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalUsers
}

func (g *GetStatsOkResponseData) SetTotalWebDownloads(totalWebDownloads float64) {
	g.TotalWebDownloads = &totalWebDownloads
}

func (g *GetStatsOkResponseData) GetTotalWebDownloads() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalWebDownloads
}