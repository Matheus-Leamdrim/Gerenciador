import 'package:equatable/equatable.dart';

abstract class TaskEvent extends Equatable {
  @override
  List<Object> get props => [];
}

class AddTask extends TaskEvent {
  final String taskName;

  AddTask(this.taskName);

  @override
  List<Object> get props => [taskName];
}

class DeleteTask extends TaskEvent {
  final int index;

  DeleteTask(this.index);

  @override
  List<Object> get props => [index];
}

class ToggleTask extends TaskEvent {
  final int index;

  ToggleTask(this.index);

  @override
  List<Object> get props => [index];
}

class RenameTask extends TaskEvent {
  final int index;
  final String newName;

  RenameTask(this.index, this.newName);

  @override
  List<Object> get props => [index, newName];
}
