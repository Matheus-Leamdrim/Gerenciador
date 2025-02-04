import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/io.dart';

class WebSocketService {
  final WebSocketChannel channel;

  WebSocketService(String url) : channel = IOWebSocketChannel.connect(url);

  Stream get stream => channel.stream;

  void sendMessage(dynamic message) {
    channel.sink.add(message);
  }

  void close() {
    channel.sink.close();
  }
}