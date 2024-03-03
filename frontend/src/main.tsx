import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import './indec.css';
import { ThemeProvider } from "@/layout/theamproviders/theam.tsx";

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider defaultTheme="light" storageKey="horse-ui-theme">
      <App />
    </ThemeProvider>
  </React.StrictMode>,
)
