import logo from './logo.svg';
import './App.css';
import List from './List';

function App() {
  return (
    <>
    <div className="App">
    If i remember it radio!!
    </div>
    
    <div className="radio">
    <audio controls>
  {/* <source src="horse.ogg" type="audio/ogg"/>
  <source src="horse.mp3" type="audio/mpeg" /> */}
</audio>
<List />
    </div>

    </>
    
  );
}

export default App;
