import 'dart:io';
import 'dart:convert';

void main() async {

  List<Map<String, dynamic>> tasks = [
    {'name': 'Estudar Flutter', 'completed': false},
    {'name': 'Fazer Exercícios', 'completed': true},
    {'name': 'Beber Água', 'completed': false},
  ];

  final server = await HttpServer.bind(InternetAddress.loopbackIPv4, 8080);
  print('Servidor WebSocket rodando em ws://localhost:8080');

  await for (HttpRequest request in server) {
    if (WebSocketTransformer.isUpgradeRequest(request)) {
 
      final websocket = await WebSocketTransformer.upgrade(request);
      print('Cliente conectado!');


      websocket.add(jsonEncode(tasks));


      websocket.listen((message) {
        print('Mensagem recebida do cliente: $message');


        final decodedMessage = jsonDecode(message) as Map<String, dynamic>;
        final action = decodedMessage['action'];

        switch (action) {
          case 'add':

            final task = decodedMessage['task'];
            tasks.add(task);
            print('Nova tarefa adicionada: ${task['name']}');
            break;

          case 'delete':

            final index = decodedMessage['index'];
            if (index >= 0 && index < tasks.length) {
              final deletedTask = tasks.removeAt(index);
              print('Tarefa deletada: ${deletedTask['name']}');
            } else {
              print('Índice inválido para deletar tarefa: $index');
            }
            break;

          case 'toggle':

            final index = decodedMessage['index'];
            if (index >= 0 && index < tasks.length) {
              tasks[index]['completed'] = !tasks[index]['completed'];
              print(
                  'Tarefa alternada: ${tasks[index]['name']} -> ${tasks[index]['completed']}');
            } else {
              print('Índice inválido para alternar tarefa: $index');
            }
            break;

          case 'rename':

            final index = decodedMessage['index'];
            final newName = decodedMessage['newName'];
            if (index >= 0 && index < tasks.length) {
              tasks[index]['name'] = newName;
              print('Tarefa renomeada: $newName');
            } else {
              print('Índice inválido para renomear tarefa: $index');
            }
            break;

          default:
            print('Ação desconhecida: $action');
        }

        websocket.add(jsonEncode(tasks));
      }, onDone: () {
        print('Cliente desconectado');
      });
    } else {
      request.response.statusCode = HttpStatus.forbidden;
      request.response.close();
    }
  }
}