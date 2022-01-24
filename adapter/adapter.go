/**
 * @Author: Seedzheng
 * 适配器模式
 */

package adapter

import "fmt"

type Player interface {
	ShowLyric()
	PlayMusic(isStop bool, prev bool, next bool, addVolume int)
	TurnOff()
}

type NormalPlayer struct {
}

func (n NormalPlayer) ShowLyric() {
	fmt.Println("播放器显示歌词")
}

func (n NormalPlayer) PlayMusic(isStop bool, prev bool, next bool, addVolume int) {
	fmt.Println("播放器播放音乐")
}

func (n NormalPlayer) TurnOff() {
	fmt.Println("关闭播放器")
}

type PlayerOpts struct {
	IsStop    bool
	Prev      bool
	Next      bool
	AddVolume int
}

type PlayerAdapter interface {
	ShowLyric()
	PlayMusic(*PlayerOpts)
	TurnOff()
}

type SmartPlayer struct {
	Player
}

func (s SmartPlayer) PlayMusic(opts *PlayerOpts) {
	fmt.Println("智能播放器播放音乐")
}
