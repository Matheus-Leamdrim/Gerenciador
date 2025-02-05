import { Task } from '../types/Task';

interface TaskItemProps {
    task: Task;
    onDelete: () => void;
    onToggleCompletion: (id: number) => void;
}

const TaskItem = ({ task, onDelete, onToggleCompletion }: TaskItemProps) => {
    return (
        <div className="flex items-center justify-between w-full">
            {/* Checkbox e Título */}
            <div className="flex items-center flex-1">
                <input
                    type="checkbox"
                    checked={task.completed}
                    onChange={() => onToggleCompletion(task.id)}
                    className="mr-2"
                />
                <span
                    className={`flex-1 text-gray-800 dark:text-gray-200 ${
                        task.completed ? 'line-through text-gray-500 dark:text-gray-400' : ''
                    }`}
                >
                    {task.title}
                </span>
            </div>

            {/* Botão de Deletar */}
            <button
                onClick={onDelete}
                className="bg-red-500 hover:bg-red-600 text-white font-medium py-1 px-2 rounded-md"
            >
                Excluir
            </button>
        </div>
    );
};

export default TaskItem;