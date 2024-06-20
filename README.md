# treetify

## Setup
- Webhook
```
https://treetify.fibonax.com/webhook
```

## Deploy steps

GOOS=linux GOARCH=amd64 go build -o build/linux

ssh -i ~/Dropbox/FibonaX/credentials/fibonax_keypair_100.25.146.60.pem ubuntu@100.25.146.60

tmux attach -t treetify

scp -i ~/Dropbox/FibonaX/credentials/fibonax_keypair_100.25.146.60.pem build/linux/treetify ubuntu@100.25.146.60:/home/ubuntu/treetify/treetify

scp -i ~/Dropbox/FibonaX/credentials/fibonax_keypair_100.25.146.60.pem conf/prod/config.yml ubuntu@100.25.146.60:/home/ubuntu/treetify/conf/config.yml

scp -i ~/Dropbox/FibonaX/credentials/fibonax_keypair_100.25.146.60.pem conf/language.yml ubuntu@100.25.146.60:/home/ubuntu/treetify/conf/language.yml