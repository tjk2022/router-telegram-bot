# Router Telegram Bot

[![Github Build](https://github.com/oleksiikhr/router-telegram-bot/actions/workflows/build.yml/badge.svg)](https://github.com/oleksiikhr/router-telegram-bot/actions/workflows/build.yml)
[![Support Ukraine](https://img.shields.io/badge/Support-Ukraine-FFD500?style=flat&labelColor=005BBB)](https://savelife.in.ua/en/donate-en/)

<p align="center">
    <img src="https://raw.githubusercontent.com/oleksiikhr/router-telegram-bot/main/docs/image.png?raw=true" alt="Router Telegram Bot" height="400">
</p>

> A very simple script under the telegram bot for the router, which will let you know if the light/internet at home.

## How it works

**No need to re-flash the router, you can use the native firmware.**

*ASUS RT-AC66U B1 was used.*

The cron runs and executes the script every minute. The bot responds only to users who are specified in `CHAT_IDS`.

Since ASUS deletes all files after a restart, you need to add a flash drive to your router, this will restore the task cron.

## Setup

- Create a new bot via [BotFather](https://t.me/BotFather)
- Allow SSH access and connect to the router
- Execute shell:

```shell
wget -O - https://raw.githubusercontent.com/oleksiikhr/router-telegram-bot/main/setup.sh > setup.sh
sh ./setup.sh API_TOKEN CHAT_IDS
```

`API_TOKEN` - token from BotFather

`CHAT_IDS` - chat IDs separated by a comma ([what my ID](https://t.me/userinfobot))

- Start a chat with your bot
- Connect the flash drive to your router

## License

[MIT](https://opensource.org/licenses/MIT)
