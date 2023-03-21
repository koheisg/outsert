package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	expected := `INSERT INTO table_name (id,title) VALUES
('5','マルチサイトのテーマの切り替え'),
('11','iA writer に感動しました。'),
('12','DHHの記事のメモ'),
('13','組織をスケールさせること、させないこと。プロダクトをスケールさせること、させないこと。'),
('14','config.hostsを設定した'),
('16','コピペするときに書式をペーストされるのを防ぐ'),
('17','今さら、キングダムにどっぷりハマった'),
('21','Apple TVのYoutubeアプリが使いやすくなってた！'),
('22','diffまとめ'),
('29','コダテルというコワーキングスペースでリモートワークをしてきました');
`

	input := ` id |                                         title
----+----------------------------------------------------------------------------------------
  5 | マルチサイトのテーマの切り替え
 11 | iA writer に感動しました。
 12 | DHHの記事のメモ
 13 | 組織をスケールさせること、させないこと。プロダクトをスケールさせること、させないこと。
 14 | config.hostsを設定した
 16 | コピペするときに書式をペーストされるのを防ぐ
 17 | 今さら、キングダムにどっぷりハマった
 21 | Apple TVのYoutubeアプリが使いやすくなってた！
 22 | diffまとめ
 29 | コダテルというコワーキングスペースでリモートワークをしてきました
(10 rows)`

	r := strings.NewReader(input)
	b := &bytes.Buffer{}
	if err := Run(r, b, "table_name", "id,title"); err != nil {
		t.Fatal(err)
	}

	if b.String() != expected {
		t.Errorf("unexpected output: got %q, want %q", b.String(), expected)
	}
}
