import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:mobile/blocs/task_bloc.dart';
import 'package:mobile/blocs/task_event.dart';
import 'package:mobile/blocs/task_state.dart';

class HomePage extends StatelessWidget {
  final TextEditingController _controller = TextEditingController();

  HomePage({super.key});

  void showRenameDialog(BuildContext context, int index, String currentName) {
    TextEditingController renameController = TextEditingController();
    renameController.text = currentName;

    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          backgroundColor: Colors.grey[900],
          title: const Text(
            'Renomear Tarefa',
            style: TextStyle(color: Colors.white),
          ),
          content: TextField(
            controller: renameController,
            style: const TextStyle(color: Colors.white),
            decoration: const InputDecoration(
              hintText: 'Novo nome da tarefa',
              hintStyle: TextStyle(color: Colors.grey),
              enabledBorder: UnderlineInputBorder(
                borderSide: BorderSide(color: Colors.white),
              ),
              focusedBorder: UnderlineInputBorder(
                borderSide: BorderSide(color: Colors.blue),
              ),
            ),
          ),
          actions: [
            TextButton(
              onPressed: () => Navigator.pop(context),
              child: const Text(
                'Cancelar',
                style: TextStyle(color: Colors.white),
              ),
            ),
            TextButton(
              onPressed: () {
                if (renameController.text.isNotEmpty) {
                  context.read<TaskBloc>().add(RenameTask(index, renameController.text));
                  Navigator.pop(context);
                }
              },
              child: const Text(
                'Salvar',
                style: TextStyle(color: Colors.blue),
              ),
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[900],
      appBar: AppBar(
        title: const Text(
          'Gerenciador de Tarefas',
          style: TextStyle(color: Colors.white),
        ),
        backgroundColor: Colors.grey[850],
        iconTheme: const IconThemeData(color: Colors.white),
      ),
      body: BlocBuilder<TaskBloc, TaskState>(
        builder: (context, state) {
          if (state is TaskInitialState) {
            return const Center(
              child: CircularProgressIndicator(color: Colors.white),
            );
          } else if (state is TaskErrorState) {
            return Center(
              child: Text(
                'Erro: ${state.message}',
                style: const TextStyle(color: Colors.white),
              ),
            );
          } else if (state is TaskLoadedState) {
            final tasks = state.tasks;
            return ListView.builder(
              itemCount: tasks.length,
              itemBuilder: (BuildContext context, index) {
                return Card(
                  margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                  color: Colors.grey[850],
                  child: ListTile(
                    title: Text(
                      tasks[index][0],
                      style: TextStyle(
                        decoration: tasks[index][1] ? TextDecoration.lineThrough : null,
                        color: Colors.white,
                      ),
                    ),
                    leading: Checkbox(
                      value: tasks[index][1],
                      onChanged: (value) {
                        context.read<TaskBloc>().add(ToggleTask(index));
                      },
                      activeColor: Colors.blue,
                    ),
                    trailing: IconButton(
                      icon: const Icon(Icons.delete, color: Colors.red),
                      onPressed: () {
                        context.read<TaskBloc>().add(DeleteTask(index));
                      },
                    ),
                    onTap: () => showRenameDialog(context, index, tasks[index][0]),
                  ),
                );
              },
            );
          } else {
            return const Center(
              child: Text(
                'Estado desconhecido',
                style: TextStyle(color: Colors.white),
              ),
            );
          }
        },
      ),
      floatingActionButton: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 20),
        child: Row(
          children: [
            Expanded(
              child: Container(
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(15),
                ),
                child: TextField(
                  controller: _controller,
                  decoration: const InputDecoration(
                    hintText: 'Adicione novas tarefas',
                    border: InputBorder.none,
                    contentPadding: EdgeInsets.symmetric(horizontal: 15),
                  ),
                ),
              ),
            ),
            const SizedBox(width: 10),
            FloatingActionButton(
              onPressed: () {
                if (_controller.text.isNotEmpty) {
                  context.read<TaskBloc>().add(AddTask(_controller.text));
                  _controller.clear();
                }
              },
              backgroundColor: Colors.blue, 
              child: const Icon(Icons.add, color: Colors.white),
            ),
          ],
        ),
      ),
    );
  }
}