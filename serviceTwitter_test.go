package main

import "testing"

func TestGetTweetText(t *testing.T) {
	tags := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight",
		"nine", "ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "twenty"}
	tweetMap := map[string]string{
		"This is a purely english tweet text": "This is a purely english tweet text #one #two #three #four #five #six #seven #eight #nine #ten #eleven #twelve #thirteen #fourteen #fifteen #sixteen #seventeen #eighteen #twenty",
		"挿絵 絵 イラスト 일러스트 일러스트레이션 アニメ アート マンガ イラストレーション イラストレーター ドローイング スケッチ": "挿絵 絵 イラスト 일러스트 일러스트레이션 アニメ アート マンガ イラストレーション イラストレーター ドローイング スケッチ #one #two #three #four #five #six #seven #eight #nine #ten #eleven #twelve #thirteen #fourteen #fifteen"}

	for text, expected := range tweetMap {
		actual := addTags(text, tags)
		if actual != expected {
			t.Errorf(`Expected "%v", but got "%v"`, expected, actual)
		}
	}
}
