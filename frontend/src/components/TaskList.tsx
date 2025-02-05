import { useState } from 'react';
import { Task } from '../types/Task';
import TaskItem from '../components/TaskItem';
import SearchFilter from './SearchFilter';
import TaskForm from './TaskForm';

interface TaskListProps {
    initialTasks: Task[];
}

const TaskList = ({ initialTasks }: TaskListProps) => {
    const [tasks, setTasks] = useState<Task[]>(initialTasks);
    const [filteredTasks, setFilteredTasks] = useState<Task[]>(initialTasks);
    const [editingTask, setEditingTask] = useState<Task | null>(null);

    const handleSearch = (query: string) => {
        const filtered = tasks.filter((task) =>
            task.title.toLowerCase().includes(query.toLowerCase())
        );
        setFilteredTasks(filtered);
    };

    const handleDelete = (id: number) => {
        const updatedTasks = tasks.filter((task) => task.id !== id);
        setTasks(updatedTasks);
        setFilteredTasks(updatedTasks);
    };

    const handleToggleCompletion = (id: number) => {
        const updatedTasks = tasks.map((task) =>
            task.id === id ? { ...task, completed: !task.completed } : task
        );
        setTasks(updatedTasks);
        setFilteredTasks(updatedTasks);
    };

    const handleCreateOrUpdateTask = (task: Task) => {
        if (task.id) {
            const updatedTasks = tasks.map((t) => (t.id === task.id ? task : t));
            setTasks(updatedTasks);
            setFilteredTasks(updatedTasks);
        } else {
            const newTask = { ...task, id: tasks.length + 1 };
            const updatedTasks = [...tasks, newTask];
            setTasks(updatedTasks);
            setFilteredTasks(updatedTasks);
        }
        setEditingTask(null);
    };

    return (
        <div className="flex flex-col items-center justify-center h-screen bg-gray-100 dark:bg-gray-900">
            <div className="w-full max-w-md bg-white dark:bg-gray-800 shadow-lg rounded-lg p-6">
                <h1 className="text-2xl font-bold mb-4 text-gray-800 dark:text-gray-200">
                    Gerenciador de Tarefas
                </h1>
                <SearchFilter onSearch={handleSearch} />
                {editingTask ? (
                    <TaskForm
                        task={editingTask}
                        onSubmit={handleCreateOrUpdateTask}
                    />
                ) : (
                    <button
                        onClick={() => setEditingTask({ id: 0, title: '', completed: false })}
                        className="bg-black hover:bg-slate-800 text-white font-medium py-2 px-4 rounded-md mb-4"
                    >
                        Criar Nova Tarefa
                    </button>
                )}
                <div className="space-y-2">
                    {filteredTasks.map((task) => (
                        <div
                            key={task.id}
                            className="flex items-center justify-between bg-gray-100 dark:bg-gray-700 rounded-md px-4 py-2"
                        >
                            <TaskItem
                                task={task}
                                onDelete={() => handleDelete(task.id)}
                                onToggleCompletion={handleToggleCompletion}
                            />
                            <button
                                onClick={() => setEditingTask(task)}
                                className="bg-gray-300 hover:bg-gray-400 text-gray-800 dark:bg-gray-600 dark:hover:bg-gray-500 dark:text-gray-200 font-medium py-1 px-2 rounded-md"
                            >
                                Editar
                            </button>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default TaskList;