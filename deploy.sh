GOOS=linux GOARCH=amd64 go build -o build/linux/treetify
scp -i ~/Dropbox/FibonaX/credentials/fibonax_keypair_100.25.146.60.pem -r build/linux/treetify conf/ ubuntu@100.25.146.60:~/treetify/
rm build/linux/treetify
# ssh -i ~/Dropbox/FibonaX/credentials/fibonax_keypair_100.25.146.60.pem ubuntu@100.25.146.60
# sudo docker run -d -p 80:80 \
#     -v /home/ubuntu/caddy/Caddyfile:/etc/caddy/Caddyfile \
#     --name reverse_proxy \
#     caddy
