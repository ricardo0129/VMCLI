package main

import (
    "fmt"
	"log"
	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func db() {
	// Open the my.db data file in your current directory.
    vm := initVM()
    fmt.Print(vm)
    stopVM(vm.ID)
    listVM() 
}

func sr() {
    // Create an instance of the struct
    original := VM{ID: "1298x", Created: 19}
    bytes := SerializeVM(original)
    // Serialize the struct to a byte array
    fmt.Println("Serialized byte array:", bytes)

    var decode = DeSerializeVM(bytes)
    fmt.Println("Deserialized", decode)

    vm := initVM()
    fmt.Println("inited", vm)
}

func tr() {
    boxStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	toxStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorReset)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(boxStyle)
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	// Event loop
	//ox, oy := -1, -1
    row, col := 0, 0
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
		//	} else if ev.Key() == tcell.KeyCtrlL {
		//		s.Sync()
		//	} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
		//		s.Clear()
			} else if ev.Rune() == 'a' {
		        s.SetContent(col, row, 'a', nil, boxStyle)
            } else if ev.Key() == tcell.KeyRight {
                col++
            } else if ev.Key() == tcell.KeyLeft {
                col--
            } else if ev.Key() == tcell.KeyDown {
                row++
            } else if ev.Key() == tcell.KeyUp {
                row--
            }
		}
        menu := [...]string {"VM Status", "Create VM", "Remove VM"}
        for r, text := range menu {
            for c, ch := range text {
                if r == row {
                    s.SetContent(c, r, ch, nil, toxStyle)
                } else {
                    s.SetContent(c, r, ch, nil, boxStyle)
                }
            }
        }
	}
}

func thr() {
    shellVM("28c9ad09b3fa")
}



func main() {
    // Deserialize the byte array back to a struct
    db()
    //sr()
    //thr()
}
