import 'dart:async';
import 'dart:convert';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:connectivity_plus/connectivity_plus.dart';
import 'task_event.dart';
import 'task_state.dart';

class TaskBloc extends Bloc<TaskEvent, TaskState> {
  final WebSocketChannel channel;
  final Connectivity connectivity = Connectivity();

  TaskBloc(this.channel) : super(TaskInitialState()) {

    connectivity.onConnectivityChanged.listen((result) {
      if (result == ConnectivityResult.none) {
        emit(TaskErrorState('Sem conexão com a internet'));
      }
    });


    channel.stream.listen(
      (data) {
        print('Dados recebidos do WebSocket: $data');
        try {
          final tasks = parseTasksFromServer(data);
          emit(TaskLoadedState(tasks));
        } catch (e) {
          print('Erro ao processar dados do WebSocket: $e');
          emit(TaskErrorState('Erro ao processar dados do WebSocket'));
        }
      },
      onError: (error) {
        print('Erro na conexão WebSocket: $error');
        emit(TaskErrorState('Erro na conexão WebSocket: $error'));
      },
      onDone: () {
        print('Conexão WebSocket fechada');
        emit(TaskErrorState('Conexão WebSocket fechada'));
      },
    );

    on<AddTask>((event, emit) async {
      final isConnected = await checkInternetConnection();
      if (!isConnected) {
        emit(TaskErrorState('Sem conexão com a internet'));
        return;
      }

      channel.sink.add(jsonEncode({
        'action': 'add',
        'task': {'name': event.taskName, 'completed': false},
      }));
    });

    on<DeleteTask>((event, emit) async {
      final isConnected = await checkInternetConnection();
      if (!isConnected) {
        emit(TaskErrorState('Sem conexão com a internet'));
        return;
      }

      channel.sink.add(jsonEncode({
        'action': 'delete',
        'index': event.index,
      }));
    });

    on<ToggleTask>((event, emit) async {
      final isConnected = await checkInternetConnection();
      if (!isConnected) {
        emit(TaskErrorState('Sem conexão com a internet'));
        return;
      }

      channel.sink.add(jsonEncode({
        'action': 'toggle',
        'index': event.index,
      }));
    });

    on<RenameTask>((event, emit) async {
      final isConnected = await checkInternetConnection();
      if (!isConnected) {
        emit(TaskErrorState('Sem conexão com a internet'));
        return;
      }

      channel.sink.add(jsonEncode({
        'action': 'rename',
        'index': event.index,
        'newName': event.newName,
      }));
    });
  }

  Future<bool> checkInternetConnection() async {
    var connectivityResult = await Connectivity().checkConnectivity();
    return connectivityResult != ConnectivityResult.none;
  }

  List<List<dynamic>> parseTasksFromServer(String data) {
    try {
      final jsonData = jsonDecode(data) as List;
      return jsonData.map((task) => [task['name'], task['completed']]).toList();
    } catch (e) {
      print('Erro ao parsear dados do WebSocket: $e');
      return [];
    }
  }

  @override
  Future<void> close() {
    channel.sink.close();
    return super.close();
  }
}