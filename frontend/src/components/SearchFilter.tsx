import { useState } from 'react';

interface SearchFilterProps {
    onSearch: (query: string) => void;
}

const SearchFilter = ({ onSearch }: SearchFilterProps) => {
    const [query, setQuery] = useState('');

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setQuery(e.target.value);
        onSearch(e.target.value);
    };

    return (
        <input
            type="text"
            value={query}
            onChange={handleChange}
            placeholder="Buscar tarefas..."
            className="w-full px-3 py-2 rounded-md border border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-200 mb-4"
        />
    );
};

export default SearchFilter;