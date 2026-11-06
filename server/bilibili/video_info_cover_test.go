package bilibili

import (
	"encoding/json"
	"testing"
)

// TestVideoInfoUgcSeasonEpisodeArcPic 验证合集内每个稿件的 arc.pic 能被正确解析。
// 这正是修复“批量下载时所有分集共用入口视频封面”的关键字段。
func TestVideoInfoUgcSeasonEpisodeArcPic(t *testing.T) {
	const payload = `{
		"bvid": "BV1faxq6KE83",
		"title": "合集入口视频",
		"pic": "http://example.com/entry.jpg",
		"ugc_season": {
			"title": "测试合集",
			"sections": [{
				"title": "正片",
				"episodes": [
					{
						"title": "第5集",
						"bvid": "BV1kR4q65EhY",
						"pages": [{"cid": 43377951560, "page": 1, "part": "1", "duration": 1789}],
						"arc": {"pic": "http://example.com/episode-5.jpg"}
					}
				]
			}]
		}
	}`

	var info VideoInfo
	if err := json.Unmarshal([]byte(payload), &info); err != nil {
		t.Fatalf("unmarshal VideoInfo: %v", err)
	}

	if info.Pic != "http://example.com/entry.jpg" {
		t.Fatalf("entry pic = %q, want http://example.com/entry.jpg", info.Pic)
	}
	if len(info.UgcSeason.Sections) != 1 || len(info.UgcSeason.Sections[0].Episodes) != 1 {
		t.Fatalf("unexpected season structure: %+v", info.UgcSeason)
	}
	ep := info.UgcSeason.Sections[0].Episodes[0]
	if ep.Bvid != "BV1kR4q65EhY" {
		t.Fatalf("episode bvid = %q", ep.Bvid)
	}
	if ep.Arc.Pic != "http://example.com/episode-5.jpg" {
		t.Fatalf("episode arc.pic = %q, want http://example.com/episode-5.jpg", ep.Arc.Pic)
	}
}
