/**
 * @Author: Seedzheng
 * 观察者模式
 */

package observer

import "fmt"

type PlayObserver interface {
	PlayMusicSuccess(userID, songName string)
}

type CollectUserLike struct {
}

func (c CollectUserLike) PlayMusicSuccess(userID, songName string) {
	fmt.Println(fmt.Sprintf("根据歌曲:%s，收集用户:%s 的喜好", songName, userID))
}

type CacheMusic struct {
}

func (c CacheMusic) PlayMusicSuccess(userID, songName string) {
	fmt.Println(fmt.Sprintf("为用户:%s 的播放器缓存歌曲:%s", userID, songName))
}

// SmartPlayer 智能播放器
type SmartPlayer struct {
	UserID        string
	playObservers []PlayObserver
}

func (p *SmartPlayer) SetPlayObservers(observers ...PlayObserver) {
	if len(observers) == 0 {
		return
	}
	p.playObservers = append(p.playObservers, observers...)
}

func (p *SmartPlayer) PlayMusic(songName string) {
	fmt.Println(fmt.Sprintf("播放器播放名字为%s的音乐", songName))
	for _, playObserver := range p.playObservers {
		playObserver.PlayMusicSuccess(p.UserID, songName)
	}
}
