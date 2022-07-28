docker build -t myjwt:v1 .
docker run -d -p 9090:9090 --name myjwt_1 myjwt:v1
