import React, { useState, useEffect } from 'react';
import Login from './components/Login';
import VideoShare from './components/VideoShare';
import SharedVideoList from './components/SharedVideoList';
import Notification from './components/Notification';
import './App.css';

function App() {
  const [user, setUser] = useState(null);
  const [refreshTrigger, setRefreshTrigger] = useState(0);
  const [notification, setNotification] = useState(null);


  useEffect(() => {
    const token = localStorage.getItem('authToken');
    if (token) {
      setUser({ token });
    }

      // Initialize WebSocket connection
    const socket = new WebSocket('ws://localhost:3001/broadcast-ws');

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.email) {
      setNotification(`${data.email} has just shared a new video`);
      }
    };

    socket.onclose = () => {
      console.log('WebSocket connection closed');
    };

    return () => {
      socket.close();
    };
  }, []);

  const handleLogin = (userData) => {
    localStorage.setItem('authToken', userData.token);
    setUser(userData);
  };

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    setUser(null);
  };

  const refreshSharedVideos = () => {
    setRefreshTrigger(prev => prev + 1);
  };

  return (
    <div className="App">
      <h1>Video Sharing App</h1>
      {user ? (
        <>
          <button onClick={handleLogout}>Logout</button>
          <VideoShare user={user} onVideoShared={refreshSharedVideos} />
        </>
      ) : (
        <Login onLogin={handleLogin} />
      )}
      <SharedVideoList refreshTrigger={refreshTrigger} />
      {notification && (
        <Notification
          message={notification}
          onClose={() => setNotification(null)}
        />
      )}
    </div>
  );
}

export default App;