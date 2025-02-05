import { GetServerSideProps } from 'next';
import { Task } from '../types/Task';
import TaskList from '../components/TaskList';

interface HomeProps {
    tasks: Task[];
}

export const getServerSideProps: GetServerSideProps<HomeProps> = async () => {
    const tasks: Task[] = [
        {
            id: 1,
            title: 'Estudar CSS',
            completed: false,
        },
        {
            id: 2,
            title: 'Beber Agua',
            completed: true,

        }, 
        {
            id: 3,
            title: 'Fazer Exercicios',
            completed: false,

        },
        {
            id: 4,
            title: 'Estudar Next,js',
            completed: true,

        },
    ];

    return {
        props: {
            tasks,
        },
    };
};

const Home = ({ tasks }: HomeProps) => {
    return <TaskList initialTasks={tasks} />;
};

export default Home;