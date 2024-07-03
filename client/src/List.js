import React, { useState, useEffect } from 'react';
import axios from 'axios';

function List() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const {data : response } = await axios.get('//localhost:8081/songs')
        setData(response);
      } catch (error) {
        console.error(error.message)
      } 
    }
    fetchData();
  }, []);

  return (
    <div>
      {console.log('data',data)}
      {data.map((song, index) => (
                <div key={index}>
                    <span>{song.title}</span>
                    &nbsp;
                    <span>{song.artist}</span>
                </div>
            ))}
    </div>
  );
}

export default List;