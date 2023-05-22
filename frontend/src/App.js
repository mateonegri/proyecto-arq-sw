import logo from './logo.svg';
import './App.css';
import { BrowserRouter } from 'react-router-dom'
import Navbar from './componentes/Navbar';
import {Home} from './paginas/Home';
function App() {

  return (
    <div className="App">
        <Navbar />
        <Home />
    </div>
  );
}

export default App;
