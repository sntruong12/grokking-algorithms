package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	lookForKey(mainbox)

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}

// problem, you are looking for a key in a box, grammy said the key is in a box that may or may not contain more boxes
type item struct {
	name   string
	isABox bool
	items  []item
}

var mainbox = []item{
	{
		name:   "teddy bear",
		isABox: false,
		items:  make([]item, 0),
	},
	{
		name:   "photo",
		isABox: false,
		items:  make([]item, 0),
	},
	{
		name:   "box 1",
		isABox: true,
		items: []item{
			{
				name:   "photo b",
				isABox: false,
				items:  make([]item, 0),
			},
			{
				name:   "pencil",
				isABox: false,
				items:  make([]item, 0),
			},
		},
	},
	{
		name:   "mug",
		isABox: false,
		items:  make([]item, 0),
	},
	{
		name:   "box 2",
		isABox: true,
		items: []item{
			{
				name:   "photo c",
				isABox: false,
				items:  make([]item, 0),
			},
			{
				name:   "box 3",
				isABox: true,
				items: []item{
					{
						name:   "photo d",
						isABox: false,
						items:  make([]item, 0),
					},
					{
						name:   "headphones",
						isABox: false,
						items:  make([]item, 0),
					},
				},
			},
			{
				name:   "crayons",
				isABox: false,
				items:  make([]item, 0),
			},
		},
	},
	{
		name:   "box 4",
		isABox: true,
		items: []item{
			{
				name:   "toothpick",
				isABox: false,
				items:  make([]item, 0),
			},
			{
				name:   "box 5",
				isABox: true,
				items: []item{
					{
						name:   "box 6",
						isABox: true,
						items: []item{
							{
								name:   "usb cord",
								isABox: false,
								items:  make([]item, 0),
							},
						},
					},
					{
						name:   "fake key",
						isABox: false,
						items:  make([]item, 0),
					},
				},
			},
			{
				name:   "key",
				isABox: false,
				items:  make([]item, 0),
			},
			{
				name:   "cardstock paper",
				isABox: false,
				items:  make([]item, 0),
			},
		},
	},
	{
		name:   "key",
		isABox: false,
		items:  make([]item, 0),
	},
}

// Recursive solution
func lookForKey(items []item) {
	for _, item := range items {
		if item.isABox {
			lookForKey(item.items)
		} else if item.name == "usb cord" {
			fmt.Println("found usb cord")
			fmt.Println("All items in box below")
			fmt.Println(items)
		}
	}
}
