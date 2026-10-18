package main

import (
	"bufio"
	"errors"
	"math/rand"
	"net"

	"charm.land/log/v2"
	"charm.land/ssh"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	"charm.land/wish/v2/logging"
)

const (
	host = "localhost"
	port = "6767"
)

func main() {
	srv, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		ssh.AllocatePty(),
		wish.WithMiddleware(
			logging.Middleware(),
			activeterm.Middleware(),

			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					state := 1
					//next(sess)
					wish.Printf(sess, "Welcome to a SSH Dating Story Game!\r\n")
					wish.Printf(sess, "Your name is Avery Smith, you're dating your crush named Sarah Anderson.\r\n")
					wish.Printf(sess, "However, you only have one chance, or else your crush will leave you and never come back.\r\n")
					wish.Printf(sess, "Continue? (Press on the keys to continue!)\r\n")
					wish.Printf(sess, "A. Yes\r\n")
					wish.Printf(sess, "B. No\r\n")

					scanner := bufio.NewReader(sess)
					for {
						cho, err := scanner.ReadByte()
						if err != nil {
							return
						}

						switch state {
						case 1:
							if cho == 'a' {
								wish.Printf(sess, "Awesome, you're deciding on a restaurant, you have 3 options, sushi, pizza, and burgers. Which restuarant are you choosing?\r\n")
								wish.Printf(sess, "A. Sushi\r\n")
								wish.Printf(sess, "B. Pizza\r\n")
								wish.Printf(sess, "C. Burgers\r\n")
								state = 2
							} else if cho == 'b' {
								wish.Printf(sess, "Since you're not continuing, that means this ssh session has ended its relationship with you </3, Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Not a valid option, select a or b\r\n")
							}
						case 2:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out the sushi restaurant has 0 stars review and turns out the air was very dusty. Your crush thought you were crazy so she left you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush goes to the sushi restaurant, the server greets you both and you are deciding what to order.\r\n")
									wish.Printf(sess, "Which sushi are you ordering? \r\n")
									wish.Printf(sess, "A. California Roll\r\n")
									wish.Printf(sess, "B. Nessa Roll\r\n")
									wish.Printf(sess, "C. Philly Roll\r\n")
									state = 3
								}
							} else if cho == 'c' { // got these mixed up but lazy to change it soo
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out the burger place has failed the safety inspection and turns out the floor has mold in it. Your crush puked everywhere and she left you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush goes to the burger place, the cashier greets you both and you are deciding what to order.\r\n")
									wish.Printf(sess, "A. Cheeseburger\r\n")
									wish.Printf(sess, "B. Hamburger\r\n")
									wish.Printf(sess, "C. Big Mac\r\n")
									state = 5
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out the pizza palace used cheap ingredients and turns out the air was infected with a virus. You and your crush got sick and your crush is in a coma. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush goes to the pizza palace, the cashier greets you both and you are deciding what to order.\r\n")
									wish.Printf(sess, "A. Cheese Pizza\r\n")
									wish.Printf(sess, "B. Pepperoni Pizza\r\n")
									wish.Printf(sess, "C. Pineapple Pizza\r\n")
									state = 4
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 3:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out your crush's hated sushi was california roll, she instantly breaks up with you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush really enjoyed california roll and left a 5 star review for the restaurant. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out your crush hates vegatables and immediately left you. However, the server gave you some tissues, atleast that's good. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush really enjoyed nessa roll and left a 4 star review for the restaurant. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out your crush is allergic to cream cheese and her skin turned red that she had to go to the hospital. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush really enjoyed philly roll and left a 5 star review for the restaurant. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 4:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out your crush really wanted to eat something that had protein, so she left instantly. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush really enjoyed cheese pizza and left a 5 star review for the palace. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out the pepperoni was expired 67 years ago. You and your crush got sick and your crush is in a coma. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush really enjoyed pepperoni pizza and left a 5 star review for the palace. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else if cho == 'c' {
								wish.Printf(sess, "Who puts fruit on a pizza??? I'm sorry, but I have to personally stop you here. Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 5:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out your crush is lactose intolerent and she thought you didn't know her at all and left you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush enjoyed the cheeseburger and left a 4 star review for the place. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out your crush thought you were being cheap and she didn't like you anymore. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush enjoyed the hamburger and left a 3.5 star review for the place. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out your crush in on a diet and thought you didn't know her at all and left you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush enjoyed the bacon cheeseburger and left a 4.5 star review for the place. Which place do you guys want to go next?\r\n")
									wish.Printf(sess, "A. Casino\r\n")
									wish.Printf(sess, "B. Theater\r\n")
									wish.Printf(sess, "C. Your house\r\n")
									wish.Printf(sess, "D. Gym\r\n")
									state = 6
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 6:
							if cho == 'a' {
								wish.Printf(sess, "You and your crush have arrived at the casino, what would you guys like to play?\r\n")
								wish.Printf(sess, "A. Coinflip\r\n")
								wish.Printf(sess, "B. Slots\r\n")
								wish.Printf(sess, "C. Horserace\r\n")
								state = 7
							} else if cho == 'b' {
								wish.Printf(sess, "You and your crush have arrived at the theater, which movies would you guys like to watch?\r\n")
								wish.Printf(sess, "A. The Emoji Movie\r\n")
								wish.Printf(sess, "B. The new minions movie\r\n")
								wish.Printf(sess, "C. Red Dead Redemption the movie\r\n")
								state = 8
							} else if cho == 'c' {
								wish.Printf(sess, "You and your crush have arrived at your house. What do you guys wanna do?\r\n")
								wish.Printf(sess, "A. Play Chess\r\n")
								wish.Printf(sess, "B. 1v1 each other on fortnite\r\n")
								state = 9
							} else if cho == 'd' {
								wish.Printf(sess, "Your crush thought you were calling her fat so she left you. Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, c, or d.\r\n")
							}
						case 7:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Your crush tells you to gamble everything on heads, however the coin lands on tails, you lost everything and your crush is leaving you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "Your crush tells you to gamble everything on heads, and it successfully landed on tails. You guys bring home the prize money. However it's getting late, where to go now?\r\n")
									wish.Printf(sess, "A. Sleepover at your house\r\n")
									wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
									state = 10
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Your crush tells you to gamble max, however the slots spinned and it didn't give you a single dime and your crush is leaving you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "Your crush tells you to gamble a little bit of your money, however the slots spinned and it gave you a dime, which is okay for you and your crush. However it's getting late, where to go now?\r\n")
									wish.Printf(sess, "A. Sleepover at your house\r\n")
									wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
									state = 10
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Your crush tells you to gamble all on red horse, however the red horse somehow died before the announcer said go? Your crush says goodbye right after. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "Your crush tells you to gamble all on red horse, and the red horse won and you won some money! Hooray! However it's getting late, where to go now?\r\n")
									wish.Printf(sess, "A. Sleepover at your house\r\n")
									wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 8:
							if cho == 'a' {
								wish.Printf(sess, "The Emoji Movie? Seriously? That movie is so bad that it had to be free on youtube movies. Game over.\r\n")
								return
							} else if cho == 'b' {
								wish.Printf(sess, "You and your crush really enjoyed watching the new minions movie. However it's getting late, where to go now?\r\n")
								wish.Printf(sess, "A. Sleepover at your house\r\n")
								wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
								state = 10
							} else if cho == 'c' {
								wish.Printf(sess, "You and your crush really enjoyed watching the new red dead redemption movie since you guys liked to play the game itself. However it's getting late, where to go now?\r\n")
								wish.Printf(sess, "A. Sleepover at your house\r\n")
								wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
								state = 10
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 9:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You have won the chess game and your crush tells you GGs. However it's getting late, where to go now?\r\n")
									wish.Printf(sess, "A. Sleepover at your house\r\n")
									wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
									state = 10
								} else {
									wish.Printf(sess, "Your crush have won the chess game and you tell her GGs. However it's getting late, where to go now?\r\n")
									wish.Printf(sess, "A. Sleepover at your house\r\n")
									wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
									state = 10
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You have won the fortnite 1v1 and your crush tells you GGs. However it's getting late, where to go now?\r\n")
									wish.Printf(sess, "A. Sleepover at your house\r\n")
									wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
									state = 10
								} else {
									wish.Printf(sess, "Your crush have won the fortnite 1v1 and you tell her GGs. However it's getting late, where to go now?\r\n")
									wish.Printf(sess, "A. Sleepover at your house\r\n")
									wish.Printf(sess, "B. Sleepover at your crush's house\r\n")
									state = 10
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 10:
							if cho == 'a' {
								wish.Printf(sess, "You and your crush prepares to sleep at your house. The AC is warm and you guys are sleeping very comfortable. Day 1 Complete!\r\n")
								wish.Printf(sess, "Tomorrow is a new and easy day for you and your crush. She's really engaged with you and thinking about being with you, maybe forever. But hey, who knows?\r\n")
								wish.Printf(sess, "Day 2: You and your crush wakes up and you guys greet each other. However you guys are hungry and really wanna eat breakfast, what breakfast do you want to make?\r\n")
								wish.Printf(sess, "A. Eggs\r\n")
								wish.Printf(sess, "B. Bacon\r\n")
								state = 11
							} else if cho == 'b' {
								wish.Printf(sess, "You and your crush prepares to sleep at her house. The AC is warm and you guys are sleeping very comfortable. Day 1 Complete!\r\n")
								wish.Printf(sess, "Tomorrow is a new and easy day for you and your crush. She's really engaged with you and thinking about being with you, maybe forever. But hey, who knows?\r\n")
								wish.Printf(sess, "Day 2: You and your crush wakes up and you guys greet each other. However you guys are hungry and really wanna eat breakfast, what breakfast do you want to make?\r\n")
								wish.Printf(sess, "A. Eggs\r\n")
								wish.Printf(sess, "B. Bacon\r\n")
								state = 11
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 11:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out you left the stove on after cooking eggs, your crush's house is on fire and got burnt down. Your crush passes out is in a coma, their parents are suing you for $10000000. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You cook eggs and your crush is very impressed with your cooking skills! Where would you guys go to start the day?\r\n")
									wish.Printf(sess, "A. Musuem\r\n")
									wish.Printf(sess, "B. Universal Studios\r\n")
									wish.Printf(sess, "C. Disneyland\r\n")
									state = 12
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out you left the stove on after cooking bacon, your house is on fire and got burnt down. You passed our and now you are in a coma. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You cook bacon and your crush is very impressed with your cooking skills! Where would you guys go to start the day?\r\n")
									wish.Printf(sess, "A. Musuem\r\n")
									wish.Printf(sess, "B. Universal Studios\r\n")
									wish.Printf(sess, "C. Disneyland\r\n")
									state = 12
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 12:
							if cho == 'a' {
								wish.Printf(sess, "Musuems? Are you really that lame? Musuems are boring. Sorry, but I'm pretty sure crushes doesn't like boring boyfriends. Game over.\r\n")
								return
							} else if cho == 'b' {
								wish.Printf(sess, "You and your crush arrives at Universal Studios and really enjoyed the scene. Which rides are you guys gonna go on?\r\n") // i forgot all universal rides :skull:
								wish.Printf(sess, "A. Simpsons Ride\r\n")
								wish.Printf(sess, "B. Kung Fu Panda Ride\r\n")
								wish.Printf(sess, "C. Super Mario Kart Ride\r\n")
								state = 13
							} else if cho == 'c' {
								wish.Printf(sess, "You and your crush arrives at Disneyland and really enjoyed the scene. Which rides are you guys go on?\r\n") // same here xd
								wish.Printf(sess, "A. Mickey Mouse Tea Ride\r\n")
								wish.Printf(sess, "B. Star Wars Ride\r\n")
								wish.Printf(sess, "C. The DEATH Ride\r\n") // not even a ride at disneyland btw just a larp ride :skull:
								state = 14
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 13:
							if cho == 'a' {
								wish.Printf(sess, "You guys rode the simpsons and it was pretty fun and nostalgic. You guys are feeling hungry, which food place would you guys want to go?\r\n")
								wish.Printf(sess, "A. Chinese Express\r\n")
								wish.Printf(sess, "B. Pho House\r\n")
								wish.Printf(sess, "C. Taco Restaurant\r\n")
								state = 15
							} else if cho == 'b' {
								wish.Printf(sess, "You guys rode the kung fu panda ride and it was pretty fun and nostalgic. You guys are feeling hungry, which food place would you guys want to go?\r\n")
								wish.Printf(sess, "A. Chinese Express\r\n")
								wish.Printf(sess, "B. Pho House\r\n")
								wish.Printf(sess, "C. Taco Restaurant\r\n")
								state = 15
							} else if cho == 'c' {
								wish.Printf(sess, "You guys rode the super mario kart ride and it was pretty fun and nostalgic. You guys are feeling hungry, which food place would you guys want to go?\r\n")
								wish.Printf(sess, "A. Chinese Express\r\n")
								wish.Printf(sess, "B. Pho House\r\n")
								wish.Printf(sess, "C. Taco Restaurant\r\n")
								state = 15
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 14:
							if cho == 'a' {
								wish.Printf(sess, "You guys rode the mickey mouse tea ride, it was kinda boring but enjoyable. You guys are feeling hungry, which food place would you guys want to go?\r\n")
								wish.Printf(sess, "A. Chinese Express\r\n")
								wish.Printf(sess, "B. Pho House\r\n")
								wish.Printf(sess, "C. Taco Restaurant\r\n")
								state = 15
							} else if cho == 'b' {
								wish.Printf(sess, "You guys rode the star wars ride and it was pretty scary but fun. You guys are feeling hungry, which food place would you guys want to go?\r\n")
								wish.Printf(sess, "A. Chinese Express\r\n")
								wish.Printf(sess, "B. Pho House\r\n")
								wish.Printf(sess, "C. Taco Restaurant\r\n")
								state = 15
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out, the death ride was indeed a death ride, the ride was so intense that you and your crush passed out and ended in a coma. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "The ride was pretty scary but you guys made it out alive. You guys are feeling hungry, which food place would you guys want to go?\r\n")
									wish.Printf(sess, "A. Chinese Express\r\n")
									wish.Printf(sess, "B. Pho House\r\n")
									wish.Printf(sess, "C. Taco Restaurant\r\n")
									state = 15
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 15:
							if cho == 'a' {
								wish.Printf(sess, "You guys enter the chinese express, the food smells pretty good and the waiter is waiting for you guys to order. Which food would you guys want to order?\r\n")
								wish.Printf(sess, "A. Kung Pao Chicken\r\n")
								wish.Printf(sess, "B. Orange Chicken\r\n")
								wish.Printf(sess, "C. Super Greens\r\n")
								state = 16
							} else if cho == 'b' {
								wish.Printf(sess, "You guys enter the pho house, the pho smells pretty good and the waiter is waiting for you guys to order. Which pho would you guys want to order?\r\n")
								wish.Printf(sess, "A. Diet Pho\r\n")
								wish.Printf(sess, "B. Pho with Meatballs\r\n")
								wish.Printf(sess, "C. Pho with everything in it\r\n")
								state = 17
							} else if cho == 'c' {
								wish.Printf(sess, "You guys enter the taco restaurant, the restaurant smells like geniune mexico restaurant. The waiter is waiting for you guys to order. Which food would you guys want to order?\r\n")
								wish.Printf(sess, "A. Beef Tacos\r\n")
								wish.Printf(sess, "B. Nacho Fries\r\n")
								wish.Printf(sess, "C. Steakhouse Burrito\r\n")
								state = 18
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 16:
							if cho == 'a' {
								wish.Printf(sess, "The kung pao chicken was spicy but you guys really enjoyed the flavor. After eating, where do you guys wanna go next?\r\n")
								wish.Printf(sess, "A. Skatepark\r\n")
								wish.Printf(sess, "B. Trampoline Park\r\n")
								wish.Printf(sess, "C. Regular park\r\n")
								state = 19
							} else if cho == 'b' {
								wish.Printf(sess, "The orange chicken was bit spicy but it was good anyways. After eating, where do you guys wanna go next?\r\n")
								wish.Printf(sess, "A. Skatepark\r\n")
								wish.Printf(sess, "B. Trampoline Park\r\n")
								wish.Printf(sess, "C. Regular park\r\n")
								state = 19
							} else if cho == 'c' {
								wish.Printf(sess, "Your crush is very offended by you since she thought you were calling her fat and needed a diet. She leaves you and throws the food to your face. Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 17:
							if cho == 'a' {
								wish.Printf(sess, "Your crush is very offended by you since she thought you were calling her fat and needed a diet. She leaves you and throws the food to your face. Game over.\r\n")
								return
							} else if cho == 'b' {
								wish.Printf(sess, "The pho was a bit blend but you guys still enjoy the broth and noodles of the pho anyways. After eating, where do you guys wanna go next?\r\n")
								wish.Printf(sess, "A. Skatepark\r\n")
								wish.Printf(sess, "B. Trampoline Park\r\n")
								wish.Printf(sess, "C. Regular park\r\n")
								state = 19
							} else if cho == 'c' {
								wish.Printf(sess, "The pho was pretty amazing and you guys enjoyed every single bite of it. After eating, where do you guys wanna go next?\r\n")
								wish.Printf(sess, "A. Skatepark\r\n")
								wish.Printf(sess, "B. Trampoline Park\r\n")
								wish.Printf(sess, "C. Regular park\r\n")
								state = 19
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 18:
							if cho == 'a' {
								wish.Printf(sess, "The beef tacos was pretty juicy and you guys enjoyed it. After eating, where do you guys wanna go next?\r\n")
								wish.Printf(sess, "A. Skatepark\r\n")
								wish.Printf(sess, "B. Trampoline Park\r\n")
								wish.Printf(sess, "C. Regular park\r\n")
								state = 19
							} else if cho == 'b' {
								wish.Printf(sess, "The fries are crispy and the cheese sauce was pretty cheesy and drippy. After eating, where do you guys wanna go next?\r\n") // I like my cheese drippy bruh
								wish.Printf(sess, "A. Skatepark\r\n")
								wish.Printf(sess, "B. Trampoline Park\r\n")
								wish.Printf(sess, "C. Regular park\r\n")
								state = 19
							} else if cho == 'c' {
								wish.Printf(sess, "Steakhouse burrito was very good and you guys really enjoyed every single bite of it. After eating, where do you guys wanna go next?\r\n")
								wish.Printf(sess, "A. Skatepark\r\n")
								wish.Printf(sess, "B. Trampoline Park\r\n")
								wish.Printf(sess, "C. Regular park\r\n")
								state = 19
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 19:
							if cho == 'a' {
								wish.Printf(sess, "You guys arrive at the skatepark and there's alot of skaters around, however, you don't know how to skate and your girlfriend wants you to do a trick. Should you?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No and tell her you can't skate at all\r\n")
								state = 20
							} else if cho == 'b' {
								wish.Printf(sess, "You guys arrive at the trampoline park and you and your crush jump around for a bit, She wants you to do a backflip and you're good at backflips. Should you?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 21
							} else if cho == 'c' {
								wish.Printf(sess, "You guys arrive at a regular park and sit on a bench. You guys talk about what's happening and who's winning at the national hockey game's kickoff. You guys had alot of fun talking. However it's getting late, she had to go since she had a meeting with her friends. Go back to your home?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 22
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 20:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You try to do a trick, However you slipped on something and you landed on somebody else. Your crush was disappointed and left you. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You try to do a trick, and surprisingly you managed to do one. Your crush is very impressed of you and you feel proud of yourself. However it's getting late, she had to go since she had a meeting with her friends. Go back to your home?\r\n")
									wish.Printf(sess, "A. Yes\r\n")
									wish.Printf(sess, "B. No\r\n")
									state = 22
								}
							} else if cho == 'b' {
								wish.Printf(sess, "Even though you kinda embarrassed yourself, however your crush apperciates that you're a honest guy and not a liar. However it's getting late, she had to go since she had a meeting with her friends. Go back to your home?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 22
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 21:
							if cho == 'a' {
								wish.Printf(sess, "You do a backflip and it was so smooth that your crush is very impressed of you. However it's getting late, she had to go since she had a meeting with her friends. Go back to your home?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 22
							} else if cho == 'b' {
								wish.Printf(sess, "You told you girlfriend that you're not doing one, she thinks that you're a liar now and leaves you. Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 22:
							if cho == 'a' {
								wish.Printf(sess, "You go back to your house, It was a good day and you took a rest to recap on what you guys did today. It was pretty sick and meanwhile, your crush thinks that you're actually really nice with doing things.\r\n")
								wish.Printf(sess, "You sleep in your bed, and day 2 is over, good job! You and your crush are getting along very well!\r\n")
								wish.Printf(sess, "Day 3, In the morning, you wake up and get ready for today. Your crush calls you and asks where would you like to hang out today?\r\n")
								wish.Printf(sess, "A. Go to the arcade\r\n")
								wish.Printf(sess, "B. Play on your nintendo switch\r\n")
								wish.Printf(sess, "C. Netflix and Chill\r\n")
								state = 23
							} else if cho == 'b' {
								wish.Printf(sess, "You sit at a bench of your neighboorhood's park, It was a good day and you took a rest to recap on what you guys did today. It was pretty sick and meanwhile, your crush thinks that you're actually really nice with doing things.\r\n")
								wish.Printf(sess, "However, you were so tired that you slept and day 2 is over, good job! You and your crush are getting along very well!\r\n")
								wish.Printf(sess, "Day 3, In the morning, you wake up and get ready for today. Your crush calls you and asks where would you like to hang out today?\r\n")
								wish.Printf(sess, "A. Go to the arcade\r\n")
								wish.Printf(sess, "B. Play on your nintendo switch\r\n")
								wish.Printf(sess, "C. Netflix and Chill\r\n")
								state = 23
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 23:
							if cho == 'a' {
								wish.Printf(sess, "You guys go to the arcade, However, there are multiple arcade machines to choose from and your crush doesn't know what to play. What do you insist on playing?\r\n")
								wish.Printf(sess, "A. Pac Man\r\n")
								wish.Printf(sess, "B. Subway Surfers\r\n")
								wish.Printf(sess, "C. Crossy Road\r\n")
								state = 24
							} else if cho == 'b' {
								wish.Printf(sess, "You guys set up your nintendo switch, and it seems that you do own alot of games. Your crush doesn't know what to play. What do you insist on playing?\r\n")
								wish.Printf(sess, "A. Super Mario Kart 8 Deluxe\r\n")
								wish.Printf(sess, "B. Super Mario Party\r\n")
								wish.Printf(sess, "C. Super Smash Bros Ultimate\r\n")
								state = 25
							} else if cho == 'c' {
								wish.Printf(sess, "You guys set up netflix, and it seems that netflix added alot of tv shows and movies. Your crush doesn't know what to watch. What do you insist on watching?\r\n")
								wish.Printf(sess, "A. South Park\r\n")
								wish.Printf(sess, "B. Family Guy\r\n")
								wish.Printf(sess, "C. Squid Game\r\n")
								state = 26
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 24:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys 1v1 on pac man and it seems that your crush has won the game! However you guys are feeling hungry, there are new places that just opened up. Which place would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								} else {
									wish.Printf(sess, "You guys 1v1 on pac man and it seems that you have won the game! However you guys are feeling hungry, there are new places that just opened up. Which place would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys played subway surfers to see who can get the most points, and it seems that your crush got the most points! However you guys are feeling hungry, there are new places that just opened up. Which place would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								} else {
									wish.Printf(sess, "You guys played subway surfers to see who can get the most points, and it seems that you got the most points! However you guys are feeling hungry, there are new places that just opened up. Which place would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys played crossy road to see who can get the most points, and it seems that your crush got the most points! However you guys are feeling hungry, there are new places that just opened up. Which place would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								} else {
									wish.Printf(sess, "You guys played crossy road to see who can get the most points, and it seems that you got the most points! However you guys are feeling hungry, there are new places that just opened up. Which place would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 25:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys played super mario kart to see who can win most races, and it seems that your crush did! However you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								} else {
									wish.Printf(sess, "You guys played super mario kart to see who can win most races, and it seems that you did! However you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys played super mario party to see who can wins, and it seems that your crush did! However you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								} else {
									wish.Printf(sess, "You guys played super mario party to see who can wins, and it seems that you did! However you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys played super smash bros ultimate to see who can get most KOs, and it seems that your crush did! However you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								} else {
									wish.Printf(sess, "You guys played super smash bros ultimate to see who can get most KOs, and it seems that you crush did! However you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
									wish.Printf(sess, "A. Hot Chicken\r\n")
									wish.Printf(sess, "B. Olive Pastas\r\n")
									wish.Printf(sess, "C. DJ's\r\n")
									state = 27
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 26:
							if cho == 'a' {
								wish.Printf(sess, "You guys watch Family Guy, and it had some hilarious jokes and not some hilarious jokes... However it was pretty enjoyable and you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
								wish.Printf(sess, "A. Hot Chicken\r\n")
								wish.Printf(sess, "B. Olive Pastas\r\n")
								wish.Printf(sess, "C. DJ's\r\n")
								state = 27
							} else if cho == 'b' {
								wish.Printf(sess, "You guys watch South Park, and the show was pretty weird with offensive jokes and some not good stuff.... However it was pretty enjoyable and you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
								wish.Printf(sess, "A. Hot Chicken\r\n")
								wish.Printf(sess, "B. Olive Pastas\r\n")
								wish.Printf(sess, "C. DJ's\r\n")
								state = 27
							} else if cho == 'c' {
								wish.Printf(sess, "You guys watch the simpsons, it was pretty funny watching bart trolling his father homer. However you guys are feeling hungry, there are new places that just opened up. Which places would you guys like to eat at?\r\n")
								wish.Printf(sess, "A. Hot Chicken\r\n")
								wish.Printf(sess, "B. Olive Pastas\r\n")
								wish.Printf(sess, "C. DJ's\r\n")
								state = 27
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 27:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out the chefs were experimenting some kind of pepper, You and your crush enter but the pepper's smell was so strong that you guys passed out and ended up in hospital. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush enters the place, it smells pretty nice and the cashier greets you guys and asks what would you guys have. Your crush places the order of a normal hot chicken sandwich. Which food are you getting?\r\n")
									wish.Printf(sess, "A. Normal Hot Chicken Sandwich\r\n")
									wish.Printf(sess, "B. Mild Hot Chicken Sandwich\r\n")
									wish.Printf(sess, "C. GHOST REAPER Hot Chicken Sandwich\r\n")
									state = 28
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out the restaurant was getting robbed, you and your crush didn't know and enters. The robber sees you both, your crush escapes but you end up getting shot in the head. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush enters the restaurant, it smells like italy and the waiter greets you guys and asks what would you guys have. Your crush places the order of spaghetti alfredo. Which food are you getting?\r\n")
									wish.Printf(sess, "A. Mariana Spaghetti\r\n")
									wish.Printf(sess, "B. Five Cheese Ziti\r\n")
									wish.Printf(sess, "C. Fettucine Alfredo\r\n")
									state = 29
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out the restaurant was getting robbed, you and your crush didn't know and enters. The robber sees you both, your crush escapes but you end up getting shot in the head. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush enters the restaurant, it smells like italy and the waiter greets you guys and asks what would you guys have. Your crush places the order of steak burger. Which food are you getting?\r\n")
									wish.Printf(sess, "A. Stuffed Chicken\r\n")
									wish.Printf(sess, "B. Bone-In Wings\r\n")
									wish.Printf(sess, "C. Cucumber Egg Rolls\r\n")
									state = 30
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 28:
							if cho == 'a' {
								wish.Printf(sess, "You and your crush really enjoyed the hot chicken sandwiches and decided to leave $50 tip. Your crush really wants to do some activity, what activity do you prefer?\r\n")
								wish.Printf(sess, "A. Hiking\r\n")
								wish.Printf(sess, "B. Photographing\r\n")
								wish.Printf(sess, "C. Darts\r\n")
								state = 31
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out you really couldn't handle spicy foods and had to drink from the soda fountain. However, the cashier thinks that you're stealing and calls the cops on you. Your crush then leaves to avoid getting fines. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush really enjoyed the hot chicken sandwiches and decided to leave $50 tip. Your crush really wants to do some activity, what activity do you prefer?\r\n")
									wish.Printf(sess, "A. Hiking\r\n")
									wish.Printf(sess, "B. Photographing\r\n")
									wish.Printf(sess, "C. Darts\r\n")
									state = 31
								}
							} else if cho == 'c' {
								wish.Printf(sess, "Turns out they weren't kidding when it was called ghost reaper. Your tongue was on fire that you had to lick everything up, however you did get alot of germs in your body and you passed out and got an infection. Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 29:
							if cho == 'a' {
								wish.Printf(sess, "You and your crush really enjoyed the pastas and decided to leave $75 tip. Your crush really wants to do some activity, what activity do you prefer?\r\n")
								wish.Printf(sess, "A. Hiking\r\n")
								wish.Printf(sess, "B. Photographing\r\n")
								wish.Printf(sess, "C. Darts\r\n")
								state = 31
							} else if cho == 'b' {
								wish.Printf(sess, "You and your crush really enjoyed the pastas and decided to leave $75 tip. Your crush really wants to do some activity, what activity do you prefer?\r\n")
								wish.Printf(sess, "A. Hiking\r\n")
								wish.Printf(sess, "B. Photographing\r\n")
								wish.Printf(sess, "C. Darts\r\n")
								state = 31
							} else if cho == 'c' {
								wish.Printf(sess, "You and your crush really enjoyed the pastas and decided to leave $75 tip. Your crush really wants to do some activity, what activity do you prefer?\r\n")
								wish.Printf(sess, "A. Hiking\r\n")
								wish.Printf(sess, "B. Photographing\r\n")
								wish.Printf(sess, "C. Darts\r\n")
								state = 31
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 30:
							if cho == 'a' {
								wish.Printf(sess, "You and your crush really enjoyed the food and decided to leave $25 tip, Your crush really wants to do some activity, what activity do you prefer?\r\n")
								wish.Printf(sess, "A. Hiking\r\n")
								wish.Printf(sess, "B. Photographing\r\n")
								wish.Printf(sess, "C. Darts\r\n")
								state = 31
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "Turns out you choked on a bone and passed out, the bone blocked the whole mouth so you couldn't breath anymore. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You and your crush really enjoyed the food and decided to leave $25 tip, Your crush really wants to do some activity, what activity do you prefer?\r\n")
									wish.Printf(sess, "A. Hiking\r\n")
									wish.Printf(sess, "B. Photographing\r\n")
									wish.Printf(sess, "C. Darts\r\n")
									state = 31
								}
							} else if cho == 'c' {
								wish.Printf(sess, "You and your crush really enjoyed the food and decided to leave $25 tip, Your crush really wants to do some activity, what activity do you prefer?\r\n")
								wish.Printf(sess, "A. Hiking\r\n")
								wish.Printf(sess, "B. Photographing\r\n")
								wish.Printf(sess, "C. Darts\r\n")
								state = 31
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 31:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys are hiking on a cliff, but suddenly a earthquake happen and your hook gets destroyed and you fall down. It was very painful and your back get destroyed by the fall. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You guys are hiking on a cliff for few minutes, then came back down and suddenly it's getting late. Your crush wants to stay at your house and requests you to cook some dinner. What dinner are you gonna make?\r\n")
									wish.Printf(sess, "A. Mac & Cheese\r\n")
									wish.Printf(sess, "B. Pizza\r\n")
									wish.Printf(sess, "C. Curry\r\n")
									state = 32
								}
							} else if cho == 'b' {
								wish.Printf(sess, "You guys are photographing on a scene and it looks very beautiful. However it's getting late. Your crush wants to stay at your house and requests you to cook some dinner. What dinner are you gonna make?\r\n")
								wish.Printf(sess, "A. Mac & Cheese\r\n")
								wish.Printf(sess, "B. Pizza\r\n")
								wish.Printf(sess, "C. Curry\r\n")
								state = 32
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys play darts to see who can hit the target first, it turns out that your crush did! However it's getting late. Your crush wants to stay at your house and requests you to cook some dinner. What dinner are you gonna make?\r\n")
									wish.Printf(sess, "A. Mac & Cheese\r\n")
									wish.Printf(sess, "B. Pizza\r\n")
									wish.Printf(sess, "C. Curry\r\n")
									state = 32
								} else {
									wish.Printf(sess, "You guys play darts to see who can hit the target first, it turns out that you did! However it's getting late. Your crush wants to stay at your house and requests you to cook some dinner. What dinner are you gonna make?\r\n")
									wish.Printf(sess, "A. Mac & Cheese\r\n")
									wish.Printf(sess, "B. Pizza\r\n")
									wish.Printf(sess, "C. Curry\r\n")
									state = 32
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 32:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "As you're cooking the mac & cheese, however you accidently turn the heat so high that you house gets burnt down and you forgot that your crush was still in there. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You're done cooking the mac & cheese, you then serve it to yourself and your crush. Your crush thinks this is very good and cheesy, she goes to bed and then sleeps. What do you wanna do now?\r\n")
									wish.Printf(sess, "A. Go to sleep")
									wish.Printf(sess, "B. Listen to some lofi beats to relax")
									state = 33
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "As you're baking the pizza, however you accidently forgot the owen was on that the pizza caught on fire, then your house gets burnt down and you forgot that your crush was still in there. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You're done baking the pizza, you then serve it to yourself and your crush. Your crush thinks this is very good and cheesy, she goes to bed and then sleeps. What do you wanna do now?\r\n")
									wish.Printf(sess, "A. Go to sleep")
									wish.Printf(sess, "B. Listen to some lofi beats to relax")
									state = 33
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "As you're cooking the steak, however you accidently turn the heat so high that you house gets burnt down and you forgot that your crush was still in there. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You're done cooking the steak, you then serve it to yourself and your crush. Your crush thinks this is very good and juicy, she goes to bed and then sleeps. What do you wanna do now?\r\n")
									wish.Printf(sess, "A. Go to sleep")
									wish.Printf(sess, "B. Listen to some lofi beats to relax")
									state = 33
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 33:
							if cho == 'a' {
								wish.Printf(sess, "You go to sleep and dreamed about your crush, it looks like she's in a white dress, could this foreshadow something?\r\n")
								wish.Printf(sess, "and that's Day 3 completed, great job! You're almost there! Day 4 starts as you and your crush wakes up.\r\n")
								wish.Printf(sess, "Your crush asks for a hot cup of coffee, what coffee do you want to make?\r\n")
								wish.Printf(sess, "A. Regular Coffee\r\n")
								wish.Printf(sess, "B. Espresso Coffee\r\n")
								wish.Printf(sess, "C. Latte Coffee\r\n")
								state = 34
							} else if cho == 'b' {
								wish.Printf(sess, "You listen to lofi's relaxing beats and you fell asleep after 4 mins and dreamed about your crush, it looks like she's in a white dress, could this foreshadow something?\r\n")
								wish.Printf(sess, "and that's Day 3 completed, great job! You're almost there! Day 4 starts as you and your crush wakes up.\r\n")
								wish.Printf(sess, "Your crush asks for a hot cup of coffee, what coffee do you want to make?\r\n")
								wish.Printf(sess, "A. Regular Coffee\r\n")
								wish.Printf(sess, "B. Espresso Coffee\r\n")
								wish.Printf(sess, "C. Latte Coffee\r\n")
								state = 34
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 34:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "While you're making the coffee, the coffee maker shakes and and the liquid speads to the cord, your house gets burnt down and your crush was still in there. Game over.\r\n")
								} else {
									wish.Printf(sess, "You're done making the coffee, you give it to your crush and she said that this is the best coffee she ever drank in a very long time. You feel proud of yourself. However there's heavy rain outside and you guys can't go anywhere. Your crush suggests playing some board games. What board game should you guys play?\r\n")
									wish.Printf(sess, "A. Monopoly\r\n")
									wish.Printf(sess, "B. Connect 4\r\n")
									wish.Printf(sess, "C. Candyland\r\n")
									state = 35
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "While you're making the espresso, the coffee maker shakes and and the liquid speads to the cord, your house gets burnt down and your crush was still in there. Game over.\r\n")
								} else {
									wish.Printf(sess, "You're done making the espresso, you give it to your crush and she said that this is the best espresso coffee she ever drank in a very long time. You feel proud of yourself. However there's heavy rain outside and you guys can't go anywhere. Your crush suggests playing some board games. What board game should you guys play?\r\n")
									wish.Printf(sess, "A. Monopoly\r\n")
									wish.Printf(sess, "B. Connect 4\r\n")
									wish.Printf(sess, "C. Candyland\r\n")
									state = 35
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "While you're making the latte, the coffee maker shakes and and the liquid speads to the cord, your house gets burnt down and your crush was still in there. Game over.\r\n")
								} else {
									wish.Printf(sess, "You're done making the latte, you give it to your crush and she said that this is the best latte coffee she ever drank in a very long time. You feel proud of yourself. However there's heavy rain outside and you guys can't go anywhere. Your crush suggests playing some board games. What board game should you guys play?\r\n")
									wish.Printf(sess, "A. Monopoly\r\n")
									wish.Printf(sess, "B. Connect 4\r\n")
									wish.Printf(sess, "C. Candyland\r\n")
									state = 35
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 35:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys are playing monopoly to see who can get most properties, and it turns out your crush did! However, the rain still continues and you guys are hungry. Since your crush trusts you with making food, she wants you to make lunch. What lunch should you make?\r\n")
									wish.Printf(sess, "A. Bacon Fries\r\n")
									wish.Printf(sess, "B. Popcorn Shrimp\r\n")
									wish.Printf(sess, "C. Buffalo Wings\r\n")
									state = 36
								} else {
									wish.Printf(sess, "You guys are playing monopoly to see who can get most properties, and it turns out you did! However, the rain still continues and you guys are hungry. Since your crush trusts you with making food, she wants you to make lunch. What lunch should you make?\r\n")
									wish.Printf(sess, "A. Bacon Fries\r\n")
									wish.Printf(sess, "B. Popcorn Shrimp\r\n")
									wish.Printf(sess, "C. Buffalo Wings\r\n")
									state = 36
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys are playing connect 4 to see who wins, and it turns out your crush did! However, the rain still continues and you guys are hungry. Since your crush trusts you with making food, she wants you to make lunch. What lunch should you make?\r\n")
									wish.Printf(sess, "A. Bacon Fries\r\n")
									wish.Printf(sess, "B. Popcorn Shrimp\r\n")
									wish.Printf(sess, "C. Buffalo Wings\r\n")
									state = 36
								} else {
									wish.Printf(sess, "You guys are playing connect 4 to see who wins, and it turns out you did! However, the rain still continues and you guys are hungry. Since your crush trusts you with making food, she wants you to make lunch. What lunch should you make?\r\n")
									wish.Printf(sess, "A. Bacon Fries\r\n")
									wish.Printf(sess, "B. Popcorn Shrimp\r\n")
									wish.Printf(sess, "C. Buffalo Wings\r\n")
									state = 36
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You guys are playing candyland to see who gets to the furthest, and it turns out your crush did! However, the rain still continues and you guys are hungry. Since your crush trusts you with making food, she wants you to make lunch. What lunch should you make?\r\n")
									wish.Printf(sess, "A. Bacon Fries\r\n")
									wish.Printf(sess, "B. Popcorn Shrimp\r\n")
									wish.Printf(sess, "C. Buffalo Wings\r\n")
									state = 36
								} else {
									wish.Printf(sess, "You guys are playing candyland to see who gets to the furthest, and it turns out you did! However, the rain still continues and you guys are hungry. Since your crush trusts you with making food, she wants you to make lunch. What lunch should you make?\r\n")
									wish.Printf(sess, "A. Bacon Fries\r\n")
									wish.Printf(sess, "B. Popcorn Shrimp\r\n")
									wish.Printf(sess, "C. Buffalo Wings\r\n")
									state = 36
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 36:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You are frying the fries, however you left the heat up too much that the grease starts to spread everywhere and especially to the outlet. Your house is on fire and your crush was still in there. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You are done frying the fries, you serve it to your crush and yourself. The fries were crispy and your crush says great job. It is still raining outside, you guys wanna do something creative. Your crush suggests to do some art, which art are you going to draw from?\r\n")
									wish.Printf(sess, "A. Mona Lisa\r\n")
									wish.Printf(sess, "B. Virtuvian Man\r\n")
									wish.Printf(sess, "C. St Paul Preaching\r\n")
									state = 37
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You are frying the shrimp, however you left the heat up too much that the grease starts to spread everywhere and especially to the outlet. Your house is on fire and your crush was still in there. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You are done frying the shrimp, you serve it to your crush and yourself. The shrimp was crispy and smooth so your crush says great job. It is still raining outside, you guys wanna do something creative. Your crush suggest to do som art, which art are you going to draw from?\r\n")
									wish.Printf(sess, "A. Mona Lisa\r\n")
									wish.Printf(sess, "B. Virtuvian Man\r\n")
									wish.Printf(sess, "C. St Paul Preaching\r\n")
									state = 37
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You are cooking the wings, however you left the stove on after cooking and the house gets on fire. Your crush was still in there. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You are done cooking the wings, you serve it to your crush and yourself. The wings was juicy and tasty so your crush says great job. It is still raining outside, you guys wanna do something creative. Your crush suggests to do some art, which art are you going to draw fram?\r\n")
									wish.Printf(sess, "A. Mona Lisa\r\n")
									wish.Printf(sess, "B. Virtuvian Man\r\n")
									wish.Printf(sess, "C. St Paul Preaching\r\n")
									state = 37
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 37:
							if cho == 'a' {
								wish.Printf(sess, "Your crush decided to paint Irises, you have decided to draw Mona Lisa. After you guys are done drawing your crush looks at your painting and she says she likes it. You feel impressed. However it's getting late and the rain still isn't gone, she says that she has to go now. What do you wanna do now?\r\n")
								wish.Printf(sess, "A. Clean the dishes\r\n")
								wish.Printf(sess, "B. Watch some news\r\n")
								wish.Printf(sess, "C. Play the Guitar\r\n")
								state = 38
							} else if cho == 'b' {
								wish.Printf(sess, "Your crush decided to paint Irises, you have decided to draw Virtuvian Man. After you guys are done drawing your crush looks at your painting and she says she likes it. You feel impressed. However it's getting late and the rain still isn't gone, she says that she has to go now. What do you wanna do now?\r\n")
								wish.Printf(sess, "A. Clean the dishes\r\n")
								wish.Printf(sess, "B. Watch some news\r\n")
								wish.Printf(sess, "C. Play the Guitar\r\n")
								state = 38
							} else if cho == 'c' {
								wish.Printf(sess, "Your crush decided to paint Irises, you have decided to draw St Paul Preaching. After you guys are done drawing your crush looks at your painting and she says that she like it. You feel impressed. However it's getting late and the rain still isn't gone, she says that she has to go now. What do you wanna do now?\r\n")
								wish.Printf(sess, "A. Clean the dishes\r\n")
								wish.Printf(sess, "B. Watch some news\r\n")
								wish.Printf(sess, "C. Play the Guitar\r\n")
								state = 38
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 38:
							if cho == 'a' {
								wish.Printf(sess, "You have decided the clean the dishes and decided to listen to some chill beats, you felt comfort inside your heart and after cleaning the dishes, should you go to bed or look at the stars?\r\n")
								wish.Printf(sess, "A. Sleep\r\n")
								wish.Printf(sess, "B. Look at the stars\r\n")
								state = 39
							} else if cho == 'b' {
								wish.Printf(sess, "You have decided to watch some news, You turn the tv on. It seems that tomorrow there's a parade celebrating st. patrick's day. You're thinking to take your crush to one of these parades. After watching the news, should you go to bed or look at the stars?\r\n")
								wish.Printf(sess, "A. Sleep\r\n")
								wish.Printf(sess, "B. Look at the stars\r\n")
								state = 39
							} else if cho == 'c' {
								wish.Printf(sess, "You have decided to play the guitar, you develop a catchy tune that you hope to share with your crush. After playing the guitar, should you go to bed or look at the stars?\r\n")
								wish.Printf(sess, "A. Sleep\r\n")
								wish.Printf(sess, "B. Look at the stars\r\n")
								state = 39
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 39:
							if cho == 'a' {
								wish.Printf(sess, "You have decided to sleep, no dreams today, but hopefully tomorrow. And.... that's day 4! Congrats!\r\n")
								wish.Printf(sess, "Day 5 begins, you wake up and your crush texts you, she said that hopefully you can come to her house to make breakfast. Should you?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 40
							} else if cho == 'b' {
								wish.Printf(sess, "You have decided to look at the stars, they look so beautiful and reminds you of you and your crush. You decided to go to sleep. And.... that's day 4! Congrats!\r\n")
								wish.Printf(sess, "Day 5 begins, you wake up and your crush texts you, she said that hopefully you can come to her house to make breakfast. Should you?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 40
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 40:
							if cho == 'a' {
								wish.Printf(sess, "You go to your crush's house, she says that she's excited to eat your handmade breakfast. What breakfast should you make?\r\n")
								wish.Printf(sess, "A. Bacon\r\n")
								wish.Printf(sess, "B. Eggs\r\n")
								wish.Printf(sess, "C. Pancakes\r\n")
								state = 41
							} else if cho == 'b' {
								wish.Printf(sess, "You say no and your crush is very upset at you and she tells you that she doesn't like you anymore. Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b\r\n")
							}
						case 41:
							if cho == 'a' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You accidently turned the heat so high that the bacon catches on fire and so is your crush's house. She passes out and she's in a coma. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You cooked bacon and served it to your crush. She said that it was very crispy and delicious. Your crush says that there are alot of parades happening today. Which parade do you guys want to go?\r\n")
									wish.Printf(sess, "A. St Patrick's Parade\r\n")
									wish.Printf(sess, "B. Macy's Thanksgiving Parade\r\n")
									wish.Printf(sess, "C. Rose Parade\r\n")
									state = 42
								}
							} else if cho == 'b' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You accidently turned the heat so high that the egg catches on fire and so is your crush's house. She passes out and she's in a coma. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You cooked eggs and served it to your crush. She said that it was very soft and delicious. Your crush says that there are alot of parades happening today. Which parade do you guys want to go?\r\n")
									wish.Printf(sess, "A. St Patrick's Parade\r\n")
									wish.Printf(sess, "B. Macy's Thanksgiving Parade\r\n")
									wish.Printf(sess, "C. Rose Parade\r\n")
									state = 42
								}
							} else if cho == 'c' {
								if rand.Float64() < 0.5 {
									wish.Printf(sess, "You accidently turned the heat so high that the pancakes catches on fire and so is your crush's house. She passes out and she's in a coma. Game over.\r\n")
									return
								} else {
									wish.Printf(sess, "You cooked pancakes and topped up with maple syrup and served it to your crush. She said that it was very soft and sweet. Your crush says that there are alot of parades happening today. Which parade do you guys want to go?\r\n")
									wish.Printf(sess, "A. St Patrick's Parade\r\n")
									wish.Printf(sess, "B. Macy's Thanksgiving Parade\r\n")
									wish.Printf(sess, "C. Rose Parade\r\n")
									state = 42
								}
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 42:
							if cho == 'a' {
								wish.Printf(sess, "You guys go to St Patrick's Parade, and it was very exciting to see. You're feeling thirsty and so is your crush, but there are plenty of coffeehouses around you guys. Which coffeehouse would you guys like to go?\r\n")
								wish.Printf(sess, "A. Starbucks\r\n")
								wish.Printf(sess, "B. Teacup\r\n")
								wish.Printf(sess, "C. Bob's Smoothies\r\n")
								state = 43
							} else if cho == 'b' {
								wish.Printf(sess, "You guys go to Macy's Thanksgiving Parade, and it was very epic to see. You're feeling thirsty and so is your crush, but there are plenty of coffeehouses around you guys. Which coffeehouse would you guys like to go?\r\n")
								wish.Printf(sess, "A. Starbucks\r\n")
								wish.Printf(sess, "B. Teacup\r\n")
								wish.Printf(sess, "C. Bob's Smoothies\r\n")
								state = 43
							} else if cho == 'c' {
								wish.Printf(sess, "You guys go to Rose Parade, and it was very beautiful to see. You're feeling thirsty and so is your crush, but there are plenty of coffeehouses around you guys. Which coffeehouse would you guys like to go?\r\n")
								wish.Printf(sess, "A. Starbucks\r\n")
								wish.Printf(sess, "B. Teacup\r\n")
								wish.Printf(sess, "C. Bob's Smoothies\r\n")
								state = 43
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 43:
							if cho == 'a' {
								wish.Printf(sess, "You guys go to starbucks and place a order. Your crush orders a strawberry mango drink, what would you like to have?\r\n")
								wish.Printf(sess, "A. Strawberry Drink\r\n")
								wish.Printf(sess, "B. Dragonfruit Drink\r\n")
								wish.Printf(sess, "C. Pineapple Drink\r\n")
								state = 44
							} else if cho == 'b' {
								wish.Printf(sess, "You guys go to teacup and place a order. Your crush orders a mango boba drink, what would you like to have?\r\n")
								wish.Printf(sess, "A. Regular Milk Tea with Boba\r\n")
								wish.Printf(sess, "B. Coconut Mango Milk Tea with Boba\r\n")
								wish.Printf(sess, "C. Pineapple Strawberry Milk Tea with Boba\r\n")
								state = 45
							} else if cho == 'c' {
								wish.Printf(sess, "You guys go to bob's smoothies and place a order. Your crush orders a mango smoothie, what would you like to have?\r\n")
								wish.Printf(sess, "A. Coconut Smoothie\r\n")
								wish.Printf(sess, "B. Strawberry Smoothie\r\n")
								wish.Printf(sess, "C. Pineapple Smoothie\r\n")
								state = 46
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 44:
							if cho == 'a' {
								wish.Printf(sess, "You guys drinked and the flavors were pretty sweet. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else if cho == 'b' {
								wish.Printf(sess, "You guys drinked and the flavors were pretty sweet. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else if cho == 'c' {
								wish.Printf(sess, "You guys drinked and the flavors were pretty sweet. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 45:
							if cho == 'a' {
								wish.Printf(sess, "You guys drinked and the flavors were pretty sweet and the bobas are pretty soft and juicy. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else if cho == 'b' {
								wish.Printf(sess, "You guys drinked and the flavors were pretty sweet and the bobas are pretty soft and juicy. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else if cho == 'c' {
								wish.Printf(sess, "You guys drinked and the flavors were pretty sweet and the bobas are pretty soft and juicy. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 46:
							if cho == 'a' {
								wish.Printf(sess, "You guys drinked and the smoothie's flavor was pretty sweet and smooth. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else if cho == 'b' {
								wish.Printf(sess, "You guys drinked and the smoothie's flavor was pretty sweet and smooth. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else if cho == 'c' {
								wish.Printf(sess, "You guys drinked and the smoothie's flavor was pretty sweet and smooth. Your crush really enjoyed hers and so do you. One of your crush's idol is having a concert in 2 hours. Do you guys want to go?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 47
							} else {
								wish.Printf(sess, "Invalid Choice, pick a, b, or c\r\n")
							}
						case 47:
							if cho == 'a' {
								wish.Printf(sess, "You guys go to the concert and your crush is very excited, The concert was very cool and you're starting to like it. It is getting late, your crush wants to explore some abandoned buildings, should you?\r\n")
								wish.Printf(sess, "A. Yes\r\n")
								wish.Printf(sess, "B. No\r\n")
								state = 48
							} else if cho == 'b' {
								wish.Printf(sess, "Your crush didn't really like you anymore and thought you were strict, so she left you. Game over.\r\n")
								return
							} else {
								wish.Printf(sess, "Invalid Choice, pick a or b")
							}
						case 48:
							if cho == 'a' {
								wish.Printf(sess, "You guys went into the abandoned building, your crush starts acting strange, You suddenly felt a cold behind you. You secretly escape so your crush doesn't know. What was in there anyways?\r\n")
								wish.Printf(sess, "You're worried about your crush but don't want to go anymore. You went back into your house. What do you wanna do before sleeping?\r\n")
								wish.Printf(sess, "A. Do a research on the building\r\n")
								wish.Printf(sess, "B. Just Sleep\r\n")
								state = 49
							} else if cho == 'b' {
								wish.Printf(sess, "No is not a option, YoU muSt GO theRe.\r\n")
							} else {
								wish.Printf(sess, "Invalid Choice, better pick a thAn B.\r\n")
							}
						case 49: // plot twist inspired by doki doki
							if cho == 'a' {
								wish.Printf(sess, "You do a research on the building, turns out there was serial killer who killed multiple people at the building. Wait, is that your crush?\r\n")
								wish.Printf(sess, "How could she be one of the victims? How? HOW IS THIS POSSIBLE, DO YOU KNOW ANYTHING ABOUT IT????????\r\n")
								wish.Printf(sess, "A. I wish I know something about it....\r\n")
								wish.Printf(sess, "B. I don't really know.... maybe she was a ghost?\r\n")
								state = 50
							} else if cho == 'b' {
								wish.Printf(sess, "You go to sleep, but your crush's ghost appears in front of you. She is trying to talk to you?\r\n")
								wish.Printf(sess, "Sarah) Well well well, you finally figured out huh?\r\n")
								wish.Printf(sess, "A. No I did not\r\n")
								wish.Printf(sess, "B. Yes I did, I recognized that dream really meant something.......\r\n")
								state = 51
							} else {
								wish.Printf(sess, "Come on, you know how to play this game already. Are you going to pick one or are you going to ragebait me?\r\n")
							}
						case 50:
							if cho == 'a' {
								wish.Printf(sess, "Of course you don't know anything about it.... me either, what should I do now?\r\n")
								wish.Printf(sess, "A. Go to sleep\r\n")
								wish.Printf(sess, "B. Buy a cross to protect you\r\n")
								state = 52
							} else if cho == 'b' {
								wish.Printf(sess, "Wait.... that dream... I recognized it now... It was really a foreshadow... But what should I do?\r\n")
								wish.Printf(sess, "A. Go to sleep\r\n")
								wish.Printf(sess, "B. Buy a cross to protect you\r\n")
								state = 52
							}
						case 51:
							if cho == 'a' {
								wish.Printf(sess, "Sarah) It's okay if you don't know anything.... However, today is gonna be your last day...\r\n")
								wish.Printf(sess, "A. What does that suppose to mean?\r\n")
								wish.Printf(sess, "B. No.... please! spear me!\r\n")
								state = 53
							} else if cho == 'b' {
								wish.Printf(sess, "Sarah) Well good, However, today is gonna be your last day....\r\n")
								wish.Printf(sess, "A. What does that suppose to mean?\r\n")
								wish.Printf(sess, "B. No.... please! spear me!\r\n")
								state = 53
							} else {
								wish.Printf(sess, "Pick a or b, your fate depends on it.\r\n")
							}
						case 52:
							if cho == 'a' {
								wish.Printf(sess, "Well well well, Since you know my story now, do you see what he did to me and my friends?\r\n")
								wish.Printf(sess, "A. Why didn't you say anything?\r\n")
								wish.Printf(sess, "B. I saw everything.\r\n")
								state = 54
							} else if cho == 'b' {
								wish.Printf(sess, "Avery buys a cross, he puts it next to his bed and sleeps, his crush's ghost is scared of it and gets evaporated.\r\n")
								wish.Printf(sess, "And... day 6 ends? I don't what you're going to do but, it was a very interesting week.\r\n")
								wish.Printf(sess, "Ending #2: Good Ending, was it worth it?\r\n")
								return
							} else {
								wish.Printf(sess, "No, I NEED TO KNOW WHAT I NEED TO DO! PLEASE! What do I need to do???? I can't take this anymore...\r\n")
							}
						case 53:
							if cho == 'a' {
								wish.Printf(sess, "Sarah) It means that today is your final day... on earth... any last words?\r\n")
								wish.Printf(sess, "A. I would, but this is ssh, so I'm not real and I don't have any value to live anymore...\r\n")
								wish.Printf(sess, "B. No.... please.... I gave you everything I had!\r\n")
								state = 55
							} else if cho == 'b' {
								wish.Printf(sess, "Well too late... this is going to be your last day! Say goodbye.....\r\n")
								wish.Printf(sess, "Ending #1: Bad Ending, it was indeed your last day.\r\n")
								return
							} else {
								wish.Printf(sess, "Since you're not saying anything... Might as well end you :) GOODBYE.\r\n")
								wish.Printf(sess, "Ending #3: Bad Ending, you could've answered something.\r\n")
								return
							}
						case 54:
							if cho == 'a' {
								wish.Printf(sess, "Sarah) I didn't answer anything because... I wanted to let you see the truth by eye.... Not leaving the place! Your time ends NOW.\r\n")
								wish.Printf(sess, "Ending #4, Bad Ending, well, you shouldn't have left to gain mercy.\r\n")
								return
							} else if cho == 'b' {
								wish.Printf(sess, "Sarah) Well you saw everything, but why didn't you stay at the building? I could've spared you. Goodbye.\r\n")
								wish.Printf(sess, "Ending #4, Bad Ending. well, you shouldn't have left to gain mercy.\r\n")
								return
							} else {
								wish.Printf(sess, "Sarah) Well, since you're not answering, might as well end your life now. GOODBYE.\r\n")
								wish.Printf(sess, "Ending #3, Bad Ending, you could've answered something.\r\n")
								return
							}
						case 55:
							if cho == 'a' {
								wish.Printf(sess, "Sarah) Oh really, well that's fair anyways... wait this is ssh?\r\n")
								wish.Printf(sess, "A. Uhhh yeah..... Didn't you know?\r\n")
								wish.Printf(sess, "B. No, you're kinda dumb\r\n")
								state = 56
							} else if cho == 'b' {
								wish.Printf(sess, "Sarah) Well.... Think twice before someone is a ghost. GOODBYE.\r\n")
								wish.Printf(sess, "Ending #5: Bad Ending, think twice before dating someone that is a ghost.\r\n")
								return
							} else {
								wish.Printf(sess, "Sarah) Well, since you're not answering, might as well end your life now. GOODBYE.\r\n")
								wish.Printf(sess, "Ending #3, Bad Ending, you could've answered something.\r\n")
								return
							}
						case 56:
							if cho == 'a' {
								wish.Printf(sess, "Sarah) Then why am I a ghost??? Is the author trying to make the story entertaining???? Wait, since you're the main character, why am I a ghost?\r\n")
								wish.Printf(sess, "A. Well uhhh, do you want me to make it so that the ghost stuff never happened?\r\n")
								wish.Printf(sess, "B. Well uhh, I don't really know I'm kinda dumb at making stories lmao\r\n")
								state = 57
							} else if cho == 'b' {
								wish.Printf(sess, "Sarah) Well, that wasn't very nice, rude to end your life NOW. GOODBYE.\r\n")
								wish.Printf(sess, "Ending #6, Bad Ending, Seriously??? She's a ghost and you're risking your life to make fun of a ghost? You kinda deserved this ending tbh\r\n")
								return
							} else {
								wish.Printf(sess, "Sarah) Answer me NOW!\r\n")
							}
						case 57:
							if cho == 'a' {
								wish.Printf(sess, "Sarah) Yes please! I just want to go back to my date and pretend this never happened... even though this wasn't my fault\r\n")
								wish.Printf(sess, "A. Okay then\r\n")
								state = 58
							} else if cho == 'b' {
								wish.Printf(sess, "Sarah) Wouldn't be surprised anyways... do you want to play a arcade game with me? Avery?\r\n")
								wish.Printf(sess, "A. Yes!\r\n")
								wish.Printf(sess, "B. no you don't have hands (roasted)\r\n")
								state = 59
							} else {
								wish.Printf(sess, "Sarah) HEY! Don't just ignore me!\r\n")
							}
						case 58:
							if cho == 'a' {
								wish.Printf(sess, "The story continues. (pretend the ghost thing never happened) You and your crush slept at your house and thinking about stuff you did over this week. You are nervous to ask your crush, but had to do it. You asked, and your crush said yes! CONGRATS!!!\r\n")
								wish.Printf(sess, "Ending #7, True Good Ending, Thanks for playing! :) (Made for Hack club's wrong project)\r\n")
								return
							} else {
								wish.Printf(sess, "Sarah) Uhh I'm pretty sure there's no other options but a\r\n")
							}
						case 59:
							if cho == 'a' {
								wish.Printf(sess, "The story continues as they played arcade games together, however atleast your crush isn't going to hunt you anymore and instead be forever with you.\r\n")
								wish.Printf(sess, "Ending #8, Maybe True Good Ending, Thanks for playing! :) (Made for Hack club's wrong project)\r\n")
							} else if cho == 'b' {
								wish.Printf(sess, "Sarah) wow.... since you're not being nice. Rude doesn't get you anywhere, especially tomorrow. GOODBYE.\r\n")
								wish.Printf(sess, "Ending #9, Bad Ending, Aw come on! Did you really have to roast a ghost to get a bad ending?\r\n")
								return
							} else {
								wish.Printf(sess, "Sarah) Waiting for your choice... Just don't try to roast me or whatever.\r\n")
							}
						default:
							wish.Printf(sess, "I really don't know how you got here. You shouldn't be in here, this is so embarrasing for me...\r\n")
							return
						}
					}
				}
			},
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}
	log.Info("Starting SSH server", "host", host, "port", port)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not start server", "errpr", err)
	}
}
