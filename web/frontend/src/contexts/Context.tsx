// MyContext.tsx

import React, { createContext, useContext, useState, ReactNode } from 'react';

// Define the context interface
interface MyContextProps {
  days: string;
  setDays: React.Dispatch<React.SetStateAction<string>>;
}

// Create the context
const MyContext = createContext<MyContextProps | undefined>(undefined);

// Create a provider component
export const MyContextProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [days, setDays] = useState<string>('7');

  return (
    <MyContext.Provider value={{ days, setDays }}>
      {children}
    </MyContext.Provider>
  );
};

// Create a custom hook to easily access the context
export const useMyContext = () => {
  const context = useContext(MyContext);

  if (context === undefined) {
    throw new Error('useMyContext must be used within a MyContextProvider');
  }

  return context;
};
