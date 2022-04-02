FROM golang:1.17

WORKDIR $home\source\repos\configRoot

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get github.com/Jacobbrewer1/configRoot

RUN go build -a -v -work -o /configrootexe

CMD [ "/configrootexe" ]