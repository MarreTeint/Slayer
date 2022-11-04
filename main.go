package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	Body               image.Image
	assetRow, assetCol int
}

func MoveDown(Character *Character) {
	if Character.Y < 254-16 {
		Character.Y++
	}
}
func MoveUp(Character *Character) {
	if Character.Y > 0 {
		Character.Y--
	}
}
func MoveLeft(Character *Character) {
	if Character.X > 0 {
		Character.X--
	}
}
func MoveRight(Character *Character) {
	if Character.X < 254-16 {
		Character.X++
	}
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

type Game struct {
	// Game's player
	Player Character
	frame  int
	tile   *ebiten.Image
}

func (g *Game) Update() error {
	g.frame++
	var frame = g.frame % 60
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if frame%10 == 0 {
			g.Player.assetRow = (g.Player.assetRow + 1) % 4
		}
		g.Player.assetCol = 0
		MoveDown(&g.Player)
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if frame%10 == 0 {
			g.Player.assetRow = (g.Player.assetRow + 1) % 4
		}
		g.Player.assetCol = 1
		MoveUp(&g.Player)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if frame%10 == 0 {
			g.Player.assetRow = (g.Player.assetRow + 1) % 4
		}
		g.Player.assetCol = 2
		MoveLeft(&g.Player)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if frame%10 == 0 {
			g.Player.assetRow = (g.Player.assetRow + 1) % 4
		}
		g.Player.assetCol = 3
		MoveRight(&g.Player)
	} else {
		g.Player.assetRow = 0
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.Player.assetRow = 4
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw local map
	screen.DrawImage(g.tile, nil)

	//draw shadow of player
	imgShad, err := os.Open("assets/NinjaAdventure/Actor/Characters/Shadow.png")
	if err != nil {
		log.Fatal(err)
	}
	shad, _, err := image.Decode(imgShad)
	if err != nil {
		log.Fatal(err)
	}
	imgShad.Close()
	var shadow = ebiten.NewImageFromImage(shad)
	var opshadow = &ebiten.DrawImageOptions{}
	opshadow.GeoM.Translate(g.Player.X+2, g.Player.Y+12)
	screen.DrawImage(shadow, opshadow)

	// draw the player's body
	var img = ebiten.NewImageFromImage(g.Player.Body)
	var subimg = img.SubImage(image.Rect(g.Player.assetCol*16, g.Player.assetRow*16, g.Player.assetCol*16+16, g.Player.assetRow*16+16)).(*ebiten.Image)
	//player position
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.Player.X, g.Player.Y)
	screen.DrawImage(subimg, op)
	//debug fps
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 254, 254
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Slayer")

	imgFile, err := os.Open("assets/NinjaAdventure/Actor/Characters/Boy/SpriteSheet.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	imgFile.Close()

	var player = Character{
		X:          100,
		Y:          100,
		Name:       "Player",
		Health:     100,
		Attack:     10,
		Defense:    10,
		Speed:      10,
		Level:      1,
		Experience: 0,
		Inventory:  []Item{},
		Equipment:  []Item{},
		Body:       img,
		assetRow:   0,
		assetCol:   0,
	}
	var frame = 0
	var localMap = [16][16]int{
		{0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,4 ,0 ,0 ,0 ,0 ,0 ,0 ,0 },
		{0 ,12,12,12,12,2 ,2 ,2 ,4 ,12,1 ,1 ,1 ,1 ,12,0 },
		{0 ,12,12,12,12,2 ,2 ,2 ,4 ,12,1 ,1 ,1 ,1 ,12,0 },
		{0 ,12,12,12,12,2 ,2 ,2 ,4 ,12,1 ,1 ,1 ,1 ,12,0 },
		{0 ,12,12,12,12,5 ,3 ,3 ,9 ,10,3 ,3 ,3 ,10,3 ,3 },
		{0 ,12,12,12,12,1 ,1 ,1 ,1 ,4 ,12,12,12,4 ,12,0 },
		{0 ,12,12,12,12,1 ,1 ,1 ,1 ,4 ,12,12,12,4 ,12,0 },
		{0 ,12,12,12,12,1 ,1 ,1 ,1 ,4 ,12,12,12,4 ,12,0 },
		{0 ,12,12,12,12,5 ,3 ,3 ,10,11,12,12,12,4 ,12,0 },
		{0 ,12,12,12,12,12,12,12,4 ,1 ,1 ,1 ,1 ,4 ,12,0 },
		{0 ,12,12,12,12,12,12,12,4 ,1 ,1 ,1 ,1 ,4 ,12,0 },
		{0 ,12,12,12,12,12,12,12,4 ,1 ,1 ,1 ,1 ,4 ,12,0 },
		{3 ,3 ,3 ,3 ,3 ,3 ,3 ,3 ,9 ,3 ,3 ,3 ,3 ,9 ,8 ,0 },
		{0 ,12,12,12,12,12,12,12,12,12,12,12,12,12,12,0 },
		{0 ,12,12,12,12,12,12,12,12,12,12,12,12,12,12,0 },
		{0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 },
	}


	var submap = ebiten.NewImage(254,254)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			var op = &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*16), float64(j*16))
			imgMap, err := os.Open("assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png")
			if err != nil {
				log.Fatal(err)
			}
			mapImg, _, err := image.Decode(imgMap)
			if err != nil {
				log.Fatal(err)
			}
			imgMap.Close()
			var tmpmapImg = ebiten.NewImageFromImage(mapImg)
			var submapImg = tmpmapImg.SubImage(image.Rect(0, 5*16, 16, 6*16)).(*ebiten.Image)
			var submaptemp = ebiten.NewImageFromImage(submapImg);
			submap.DrawImage(submaptemp, op)
		}
	}

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			var op = &ebiten.DrawImageOptions{}
			var link string
			var selector image.Rectangle
			op.GeoM.Translate(float64(i*16), float64(j*16))
			switch localMap[j][i] {
			case 0:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetNature.png"
				selector = image.Rect(7*16, 11*16, 8*16, 12*16)
			case 1:
				if localMap[j][i-1] != 1 && localMap[j-1][i] != 1 {
					link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetHouse.png"
					selector = image.Rect(0, 0, 64, 48)
				}else{
					link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
					selector = image.Rect(32, 4*16, 48, 5*16)
				}
			case 2:
				if localMap[j][i-1] != 2 && localMap[j-1][i] != 2 {
					link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetHouse.png"
					selector = image.Rect(4*64, 0, 5*64-16, 48)
				} else {
					link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
					selector = image.Rect(32, 4*16, 48, 5*16)
				}
			case 3:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(16, 3*16, 32, 4*16)
			case 4:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(3*16, 16, 4*16, 32)
			case 5:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(0, 3*16, 16, 4*16)
			case 8:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(32, 3*16, 48, 4*16)
			case 9:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(8*16, 3*16, 9*16, 4*16)
			case 10:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(8*16, 0, 9*16, 16)
			case 11:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(7*16, 3*16, 8*16, 4*16)
			case 12:
				link = "assets/NinjaAdventure/Backgrounds/Tilesets/TilesetFloor.png"
				selector = image.Rect(3*16, 5*16, 4*16, 6*16)
			}
			imgMap, err := os.Open(link)
			if err != nil {
				log.Fatal(err)
			}
			mapImg, _, err := image.Decode(imgMap)
			if err != nil {
				log.Fatal(err)
			}
			imgMap.Close()
			var tmpmapImg = ebiten.NewImageFromImage(mapImg)
			var submapImg = tmpmapImg.SubImage(selector).(*ebiten.Image)
			var submaptemp = ebiten.NewImageFromImage(submapImg);
			submap.DrawImage(submaptemp, op)
		}
	}

	if err := ebiten.RunGame(&Game{player, frame, submap}); err != nil {
		log.Fatal(err)
	}
}
