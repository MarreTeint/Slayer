package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	//"github.com/go-delve/delve/pkg/dwarf/frame"
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
	assetRow, assetCol int	
}
func MoveDown(Character *Character) {
	Character.Y++
}
func MoveUp(Character *Character) {
	Character.Y--
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
	frame int
}

func (g *Game) Update() error {
	g.frame++
	var frame = g.frame%60
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if(frame%15==0){
			g.Player.assetRow = (g.Player.assetRow+1)%4
		}
		g.Player.assetCol = 0
		MoveDown(&g.Player)
	}else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if(frame%15==0){
			g.Player.assetRow = (g.Player.assetRow+1)%4
		}
		g.Player.assetCol = 1
		MoveUp(&g.Player)
	}else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if(frame%15==0){
			g.Player.assetRow = (g.Player.assetRow+1)%4
		}
		g.Player.assetCol = 2
		MoveLeft(&g.Player)
	}else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if(frame%15==0){
			g.Player.assetRow = (g.Player.assetRow+1)%4
		}
		g.Player.assetCol = 3
		MoveRight(&g.Player)
	}else{
		g.Player.assetRow = 0
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.Player.assetRow = 4
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw background
	imgFile,err := os.Open("assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}
	back, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	imgFile.Close()
	var background = ebiten.NewImageFromImage(back)
	var subback = background.SubImage(image.Rect(0*16, 4*16, 1*16, 5*16)).(*ebiten.Image)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			var backop = &ebiten.DrawImageOptions{}
			backop.GeoM.Translate(float64(i*16), float64(j*16))	
			screen.DrawImage(subback, backop)
		}
	}
	
	// draw the player's body
	var img = ebiten.NewImageFromImage(g.Player.Body)
	var subimg = img.SubImage(image.Rect(g.Player.assetCol*16, g.Player.assetRow*16, g.Player.assetCol*16+16, g.Player.assetRow*16+16)).(*ebiten.Image)
	//player position
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(subimg, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 254,254 
}

func main() {
	ebiten.SetWindowSize(1000, 500)
	ebiten.SetWindowTitle("Slayer")

	imgFile,err := os.Open("assets/NinjaAdventure/Actor/Characters/Boy/SpriteSheet.png")
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
		assetRow: 0,
		assetCol: 0,
	}
	var frame = 0;
	if err := ebiten.RunGame(&Game{player, frame}); err != nil {
		log.Fatal(err)
	}
}