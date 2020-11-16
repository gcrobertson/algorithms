package main

import "fmt"

/*
 *	The only difference between a singly and doubly linked list is that each node in a Doubly Linked List stores
 *  also stores a pointer to the previous node.
 *
 *  There is no way to traverse backward in a Singly Linked List.
 *
 *	To look at a previous song a Doubly Linked List must be implemented.
 *
 *	In the Singly Linked List, addSong() traverses through every node costing O(n) time.
 *	-	By storing the tail node, we can link new node to tail directly without iterating through the list.
 *	-	The addSong() method using the tail node changes the insert operation from O(n) to constant time, O(1).
 */

type song struct {
	name     string
	artist   string
	previous *song
	next     *song
}

type playlist struct {
	name       string
	head       *song
	tail       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) error {
	s := &song{
		name:   name,
		artist: artist,
	}
	if p.head == nil {
		p.head = s
	} else {
		pt := p.tail
		pt.next = s
		s.previous = pt
	}
	p.tail = s
	return nil
}

func (p *playlist) showAllSongs() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("empty playlist.")
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

func (p *playlist) previousSong() *song {
	p.nowPlaying = p.nowPlaying.previous
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
	myPlaylist.showAllSongs()
	myPlaylist.startPlaying()
	for myPlaylist.nowPlaying != nil {
		fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
		myPlaylist.nextSong()
	}
	myPlaylist.nowPlaying = myPlaylist.tail
	for myPlaylist.nowPlaying != nil {
		fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
		myPlaylist.previousSong()
	}
}
