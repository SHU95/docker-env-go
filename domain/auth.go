//ここにjwtの構造体を書くのもあり
package domain

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json;message`
}
