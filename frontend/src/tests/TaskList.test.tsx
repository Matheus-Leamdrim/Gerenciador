import { render, screen, } from '@testing-library/react';
import TaskList from '../components/TaskList';
import { Task } from '../types/Task';

const mockTasks: Task[] = [
    {
        id: 1,
        title: 'Tarefa 1',
        completed: false,
    },
];

test('exibe a lista de tarefas', () => {
    render(<TaskList initialTasks={mockTasks} />);
    expect(screen.getByText('Tarefa 1')).toBeInTheDocument();
});