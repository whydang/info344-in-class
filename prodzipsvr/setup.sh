ufw allow OpenSSH
ufw --force enable
#replace your-dockerhug-name with your docker hub name
#and your-image-name with your image name and
#uncomment the line by removing the leading #
#docker run --name zipsvr -d -p 80:80 your-dockerhub-name/your-image-name
