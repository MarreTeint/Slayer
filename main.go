package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Character struct {
	// Character's position
	X, Y float64
	// Character's name
	Name string
	// Character's health
	Health int
	// Character's attack
	Attack int
	// Character's defense
	Defense int
	// Character's speed
	Speed int
	// Character's level
	Level int
	// Character's experience
	Experience int
	// Character's inventory
	Inventory []Item
	// Character's equipment
	Equipment []Item
	// Character's body asset
	Body image.Image	
}
func MoveDown(Character *Character) {
	Character.Y--
}
func MoveUp(Character *Character) {
	Character.Y++
}
func MoveLeft(Character *Character) {
	Character.X--
}
func MoveRight(Character *Character) {
	Character.X++
}

type Item struct {
	// Item's name
	Name string
	// Item's description
	Description string
	// Item's type
	Type string
	// Item's value
	Value int
	// Item's weight
	Weight int
	// Item's asset
	Asset image.Image
}

type Game struct{
	// Game's player
	Player Character
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw the player's body
	var img = ebiten.NewImageFromImage(g.Player.Body)
	var subimg = img.SubImage(image.Rect(0, 0, 64, 64)).(*ebiten.Image)
	//player position
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(subimg, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(1600, 1200)
	ebiten.SetWindowTitle("Slayer")

	imgFile,err := os.Open("assets/character/char_a_p1/char_a_p1_0bas_humn_v01.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	imgFile.Close()

	var player = Character{
		X: 100,
		Y: 100,
		Name: "Player",
		Health: 100,
		Attack: 10,
		Defense: 10,
		Speed: 10,
		Level: 1,
		Experience: 0,
		Inventory: []Item{},
		Equipment: []Item{},
		Body: img,
	}
	
	if err := ebiten.RunGame(&Game{player}); err != nil {
		log.Fatal(err)
	}
}