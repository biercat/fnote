// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

type TodayTrafficStatsVO struct {
	ViewCount     int64 `json:"view_count"`
	UserViewCount int64 `json:"user_view_count"`
	CommentCount  int64 `json:"comment_count"`
	LikeCount     int64 `json:"like_count"`
}

type TrafficStatsVO struct {
	ViewCount    int64 `json:"view_count"`
	CommentCount int64 `json:"comment_count"`
	LikeCount    int64 `json:"like_count"`
}

type ContentStatsVO struct {
	PostCount     int64 `json:"post_count"`
	CategoryCount int64 `json:"category_count"`
	TagCount      int64 `json:"tag_count"`
}

type TendencyDataVO struct {
	PV []TendencyData `json:"pv"`
	UV []TendencyData `json:"uv"`
}

type TendencyData struct {
	Timestamp int64 `json:"timestamp"`
	ViewCount int64 `json:"view_count"`
}

type UserDistributionVO struct {
	UserCount int64  `json:"user_count"`
	Location  string `json:"location"`
}
