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
import Profile from './Pages/Profile'
import editProfile from './Pages/editProfile'

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
      <Route path="/profile" element={<Profile />}/>
      <Route path="/editprofile" element={<editProfile />}/>
    </Routes>
    </div>
  );
}

export default App;
