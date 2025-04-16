import React, { useState, useEffect } from 'react';
import axios from '../axiosConfig';
import './SharedVideoList.css';

function SharedVideoList({ refreshTrigger }) {
  const [videos, setVideos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchVideos = async () => {
      try {
        const response = await axios.get('/api/public/v1/videos');
        setVideos(response.data.videos);
        setLoading(false);
      } catch (err) {
        console.error('Error fetching shared videos:', err);
        setError('Failed to fetch shared videos. Please try again later.');
        setLoading(false);
      }
    };

    fetchVideos();
  }, [refreshTrigger]);

  const getYouTubeVideoId = (url) => {
    const regExp = /^.*(youtu.be\/|v\/|u\/\w\/|embed\/|watch\?v=|\&v=)([^#\&\?]*).*/;
    const match = url.match(regExp);
    return (match && match[2].length === 11) ? match[2] : null;
  };

  if (loading) return <div>Loading shared videos...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="shared-video-list">
      <h2>Shared Videos</h2>
      {videos.length === 0 ? (
        <p>No videos have been shared yet.</p>
      ) : (
        <ul>
          {videos.map((video, index) => {
            const videoId = getYouTubeVideoId(video.VideoUrl);
            const embedUrl = videoId ? `https://www.youtube.com/embed/${videoId}` : null;

            return (
              <li key={index} className="video-item">
                {embedUrl && (
                  <div className="video-player">
                    <iframe
                      width="280"
                      height="157"
                      src={embedUrl}
                      frameBorder="0"
                      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                      allowFullScreen
                      title={video.VideoTitle}
                    ></iframe>
                  </div>
                )}
                <div className="video-info">
                  <h3>{video.VideoTitle}</h3>
                  <p>Shared by: {video.SharerEmail}</p>
                </div>
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
}

export default SharedVideoList;