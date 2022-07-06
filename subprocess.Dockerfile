FROM docker.io/library/bash
WORKDIR /
COPY go-mem-allocate ./
COPY run.sh ./
RUN chmod 777 ./run.sh
RUN chmod 777 ./go-mem-allocate

ENTRYPOINT [ "./run.sh" ]

