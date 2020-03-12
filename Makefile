.DEFAULT_GOAL: NaviBot
.PHONY: prep get tidy update test clean

NaviBot :
	go build -o NaviBot main.go

prep: get tidy

get:
	go get "github.com/joho/godotenv"; \
	go get "github.com/bwmarrin/discordgo"

tidy:
	go mod tidy

update:
	go get -u

test: NaviBot
	./NaviBot

clean:
	rm NaviBot
