import { useState } from 'react';
import { Task } from '../types/Task';

interface TaskFormProps {
    task?: Task;
    onSubmit: (task: Task) => void;
}

const TaskForm = ({ task, onSubmit }: TaskFormProps) => {
    const [title, setTitle] = useState(task?.title || '');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        const newTask: Task = {
            ...task,
            id: task?.id || 0,
            title,
            completed: task?.completed || false,
        };
        onSubmit(newTask);
    };

    return (
        <form onSubmit={handleSubmit} className="mb-4">
            <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Título da tarefa"
                className="w-full px-3 py-2 rounded-md border border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-200"
                required
            />
            <button
                type="submit"
                className="bg-black hover:bg-slate-800 text-white font-medium py-2 px-4 rounded-md mt-2"
            >
                {task ? 'Atualizar' : 'Criar'}
            </button>
        </form>
    );
};

export default TaskForm;