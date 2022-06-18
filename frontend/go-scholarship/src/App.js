import './App.css';
import Login from './Pages/Login';
import Register from './Pages/Register';
import { Home } from './Pages/Home';
import { Routes, Route } from 'react-router-dom';
import NavbarApp from './Components/Navbar';
import Faq from './Pages/Faq';
import Contct from './Pages/ContactUs';
function App() {
  return (
    <div className="App">
      <NavbarApp />
      <Routes>
      <Route path="/" element={<Home />}/>
      <Route path="/login" element={<Login />}/>
      <Route path="/register" element={<Register />}/>
      <Route path="/faq" element={<Faq />}/>
      <Route path="/contact" element={<Contct />}/>
    </Routes>
    </div>
  );
}

export default App;
