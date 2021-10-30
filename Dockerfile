FROM alpine:3.13.5
COPY . /app
WORKDIR /app
RUN chmod +x run.sh
CMD ./run.sh
