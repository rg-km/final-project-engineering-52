import './App.css';
import Login from './Pages/Login';
import Register from './Pages/Register';
import { Home } from './Pages/Home';
import { Routes, Route } from 'react-router-dom';
import NavbarApp from './Components/Navbar';
import Faq from './Pages/Faq';
import Contct from './Pages/ContactUs';
import Detailbea from './Pages/Detailbea'
import Listbea from './Pages/Listbea'
import { CreateBeasiswa } from './Pages/CreateBeasiswa';
import Profile from './Pages/Profile'

function App() {
  return (
    <div className="App">
      <NavbarApp />
      <Routes>
      <Route path="/" element={<Home />}/>
      <Route path='/detailbea' element={<Detailbea />}/>
      <Route path='/listbea' element={<Listbea/>}/>
      <Route path="/login" element={<Login />}/>
      <Route path="/register" element={<Register />}/>
      <Route path="/faq" element={<Faq />}/>
      <Route path="/contact" element={<Contct />}/>
      <Route path="/list-beasiswa" element={<Listbea />}/>
      <Route path="/beasiswa/create" element={<CreateBeasiswa />}/>
      <Route path="/profile" element={<Profile />}/>
    </Routes>
    </div>
  );
}

export default App;
