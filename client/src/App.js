import React, { useState, useEffect } from 'react';
import './App.css';
import List from './List';
import axios from 'axios';

function App() {
  const [playingId, setPlayingId] = useState(null);

  // useEffect(() => {
  //   // For example: fetch first song id to play automatically
  //   const fetchSongs = async () => {
  //     try {
  //       const { data } = await axios.get('http://localhost:8081/songs');
  //       if (data.length > 0) {
  //         setPlayingId(data[0].id); // play first song by default
  //       }
  //     } catch (error) {
  //       console.error(error);
  //     }
  //   };
  //   fetchSongs();
  // }, []);

  useEffect(() => {
    const fetchSongs = async () => {
      try {
        const { data } = await axios.get('http://localhost:8081/songs');
        console.log('Fetched songs:', data);
        if (data.length > 0) {
          setPlayingId(data[data.length].song_id);
        } else {
          console.warn('No songs found');
        }
      } catch (error) {
        console.error('Error fetching songs:', error);
      }
    };
    fetchSongs();
  }, []);
  

  return (
    <>
      <div>
        {playingId ? (
          <audio controls autoPlay>
            <source src={`http://localhost:8081/stream/id/${playingId}`} type="audio/mpeg" />
            Your browser does not support the audio element.
          </audio>
        ) : (
          <p>Loading audio...</p>
        )}
      </div>

      <List />
    </>
  );
}

export default App;
