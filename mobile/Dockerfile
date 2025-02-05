FROM ghcr.io/cirruslabs/flutter:3.27.3

WORKDIR /mobile

COPY pubspec.yaml pubspec.lock ./


RUN flutter pub get


COPY . .


EXPOSE 8080

CMD ["flutter", "run", "--dart-define=BACKEND_URL=ws://host.docker.internal:8080"]
