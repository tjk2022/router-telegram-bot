#!/bin/bash

set -x

if [ $# -eq 0 ]
then
  echo "Pass telegram bot token"
  exit 1
fi

if [ $# -eq 1 ]
then
  echo "Pass telegram chat ids"
  exit 1
fi

# Clear latest code
rm -rf /jffs/telegram

# Download the binary
mkdir -p /jffs/telegram
wget -O /jffs/telegram/telegrambot https://github.com/oleksiikhr/router-telegram-bot/releases/latest/download/telegrambot
chmod +x /jffs/telegram/telegrambot

# Prepare cronjob and setup usbmount
cat > /jffs/telegram/$USER << EOF
* * * * * /jffs/telegram/telegrambot -token $1 -chatIds $2
EOF

cat > /jffs/telegram/init.sh << EOF
cp /jffs/telegram/$USER /var/spool/cron/crontabs
EOF

chmod +x /jffs/telegram/init.sh
nvram set script_usbmount="/jffs/telegram/init.sh"
nvram commit
