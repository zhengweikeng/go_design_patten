/**
 * @Author: Seedzheng
 */

package decorator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("decorator: ", playMp3Music)
}

func playMp3Music(t *testing.T) {
	mp3 := Mp3Player{}
	mp3.PlayMusic()
	fmt.Println("==============")

	mp4 := Mp4Player{mp3}
	mp4.PlayMusic()
	mp4.ShowLyric()
}
