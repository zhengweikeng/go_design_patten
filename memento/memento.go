/**
 * @Author: Seedzheng
 * 备忘录模式
 */

package memento

type Content struct {
	Text string
}

type Snapshot struct {
	text string
}

func (s *Snapshot) SetText(t string) {
	s.text = t
}

type CacheManager struct {
	Snapshot
}

func demo() {
	c := Content{Text: "hello"}

	cm := CacheManager{Snapshot{}}
	cm.SetText(c.Text)

	c.Text = "world"
	c.Text = cm.text
}
