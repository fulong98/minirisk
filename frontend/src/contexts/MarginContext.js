import React, { createContext, useContext, useState, useEffect } from 'react';
import axios from 'axios';

const MarginContext = createContext();

export function MarginProvider({ children }) {
  const [marginData, setMarginData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchMarginData = async (clientId = 1) => {
    try {
      setLoading(true);
      const response = await axios.get(`http://localhost:8080/api/margin/status/${clientId}`);
      setMarginData(response.data);
      setError(null);
    } catch (err) {
      setError(err.message);
      console.error('Error fetching margin data:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchMarginData();
    // Set up polling interval (every 30 seconds)
    const interval = setInterval(() => {
      fetchMarginData();
    }, 30000);

    return () => clearInterval(interval);
  }, []);

  return (
    <MarginContext.Provider
      value={{
        marginData,
        loading,
        error,
        fetchMarginData,
      }}
    >
      {children}
    </MarginContext.Provider>
  );
}

export function useMargin() {
  const context = useContext(MarginContext);
  if (!context) {
    throw new Error('useMargin must be used within a MarginProvider');
  }
  return context;
} 