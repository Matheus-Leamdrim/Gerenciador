import type { AppProps } from 'next/app';
import '../styles/globals.css';

function MyApp({ Component, pageProps }: AppProps) {
    return (
        <div className="dark">
            <Component {...pageProps} />
        </div>
    );
}

export default MyApp;