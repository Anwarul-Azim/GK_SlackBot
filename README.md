# GK SlackBot

**GK SlackBot** is a simple slack bot written in Go that can answer any question on general knowledge written in English. It uses [Wit](https://wit.ai/) as the chatbot framework and Wolfram Alpha [API](https://developer.wolframalpha.com/) for NLP  model.

  
## Installation
In order to use this chat bot, you need,

 - A slack workspace
 - A slackbot installed in that workspace with necessary (message read/write) permissions
 - An **wit** app with "wolfram_search_query" selected as `intent`
 - Access to wolfram api (app created in developers [portal](https://developer.wolframalpha.com/portal)

After completing  these prerequisites, get  four tokens/ids and set those in the `.env` file.  These are,  `slack app token`,  `slack bot token`,  `wit ai token` and  `wolfram app id`.  Also, install all the third party libraries in your directory (check Acknowledgement).

Now the program is ready to fetch question from the bot, prepare the answer and send that again to the bot.

## Acknowledgement

Instructions from this **freeCodeCamp** [tutorial](https://www.youtube.com/watch?v=jFfo23yIWac) was followed for designing this bot.  

Used third party libraries are -
 - [go-wolfram](github.com/krognol/go-wolfram)
 - [slacker](github.com/shomali11/slacker)
 - [gjson](github.com/tidwal/gjson)
 - [godotenv](github.com/joho/godotenv)
 - [wit-go](github.com/wit-ai/wit-go)

The README  is created using [stackedit](https://stackedit.io).
