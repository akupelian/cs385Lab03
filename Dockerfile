FROM golang:alpine

COPY src/goLangApp.go src/

WORKDIR src/

CMD ["go", "run", "goLangApp.go"]
