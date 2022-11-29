# gobackend
frontend: pilput-forntend

# to run 


- go build && ./gobackend

- docker build -t gobackend . && docker run -p 8080:8080 -d gobackend
