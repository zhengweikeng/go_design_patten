/**
 * @Author: Seedzheng
 */

package adapter

import "testing"

func Test(t *testing.T) {
    PlayMusic(t)
}

func PlayMusic(t *testing.T) {
    player := SmartPlayer{NormalPlayer{}}

    player.ShowLyric()
    player.PlayMusic(&PlayerOpts{
        IsStop:    false,
        Prev:      false,
        Next:      true,
        AddVolume: 10,
    })
    player.TurnOff()
}
