import React, { useEffect, useState } from 'react';
import axiosInstance from "../axiosInstance";

const LifeCirclePage = () => {
  const [lifeCircle, setLifeCircle] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchLifeCircle = async () => {
      try {
        const response = await axiosInstance.get('/api/life-circle', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        setLifeCircle(response.data.life_circle);
      } catch (err) {
        setError('Failed to fetch life circle data');
      } finally {
        setLoading(false);
      }
    };

    fetchLifeCircle();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h1>Your Life Circle</h1>
      {lifeCircle && (
        <div>
          {Object.keys(lifeCircle).map((area, index) => (
            <div key={index}>
              <h3>{area.replace(/([A-Z])/g, ' $1')}</h3>
              <div style={{ width: '100%', background: '#ddd' }}>
                <div
                  style={{
                    width: `${(lifeCircle[area] / 10) * 100}%`,
                    background: '#4caf50',
                    height: '24px',
                  }}
                />
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default LifeCirclePage;