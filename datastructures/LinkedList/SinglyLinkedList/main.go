package main

import "fmt"

/*	https://medium.com/backendarmy/linked-lists-in-go-f7a7b27a03b9
 *	Linked lists are nothing but a bunch of nodes containing the data and a pointer to access the next node.
 *
 *	Singly Linked Lists
 *	-	the most basic type of linked list where a node contains:
 *	-	they can be iterated over in one order, not reverse
 *
 *	Because the pointer to the next node is stored, the nodes do not need to be stored sequentially in memory like an array.
 *
 *	The last element of the singly linked list stores nil as it has no next node to point to.
 */

type song struct {
	name   string
	artist string
	next   *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) error {
	song := &song{
		name:   name,
		artist: artist,
	}
	if p.head == nil {
		p.head = song
		return nil
	}
	// iterate over .next but do not want to change the p.head value.
	currentNode := p.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = song

	return nil
}

func (p *playlist) showAllSongs() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

func (p *playlist) startPlaying() *song {
	p.nowPlaying = p.head
	return p.nowPlaying
}

func (p *playlist) nextSong() *song {
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func main() {
	playlistName := "myplaylist"
	myPlaylist := createPlaylist(playlistName)
	songs := map[string]string{
		"Ophelia":       "The Lumineers",
		"Shape of you":  "Ed Sheeran",
		"Stubborn Love": "The Lumineers",
		"Feels":         "Calvin Harris",
	}
	for name, artist := range songs {
		myPlaylist.addSong(name, artist)
	}
	fmt.Println("Showing all songs in playlist...")
	myPlaylist.showAllSongs()
	myPlaylist.startPlaying()
	for myPlaylist.nowPlaying != nil {
		fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
		myPlaylist.nextSong()
	}
}
