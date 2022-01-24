/**
 * @Author: Seedzheng
 * 装饰者模式
 */

package decorator

import "fmt"

// MusicPlayer 音乐播放器
type MusicPlayer interface {
	PlayMusic()
}

type Mp3Player struct {
}

func (p Mp3Player) PlayMusic() {
	fmt.Println("播放器播放音乐")
}

// LyricMusicPlayer 携带展示歌词的音乐播放器
type LyricMusicPlayer interface {
	MusicPlayer
	ShowLyric()
}

type Mp4Player struct {
	MusicPlayer
}

func (p Mp4Player) ShowLyric() {
	fmt.Println("播放器展示歌词")
}
