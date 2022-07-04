import "./App.css";
import Login from "./Pages/Login";
import Register from "./Pages/Register";
import { Home } from "./Pages/Home";
import { Routes, Route } from "react-router-dom";
import NavbarApp from "./Components/Navbar";
import Faq from "./Pages/Faq";
import Contct from "./Pages/ContactUs";
import Detailbea from "./Pages/Detailbea";
import Listbea from "./Pages/Listbea";
import { CreateBeasiswa } from "./Pages/CreateBeasiswa";
import Profile from "./Pages/Profile";
import { RequireAuth } from "./Components/ProtectedRoute";
import Admin from "./Pages/Admin";
import Scholarship from "./Pages/Admin/Scholarship";
import Category from "./Pages/Admin/Category";
import Comment from "./Pages/Admin/Comment";
import User from "./Pages/Admin/User";
import Footer from "./Components/Footer";
import Editprof from './Pages/Editprof'
import NotFound from "./Pages/404";

function App() {
  return (
    <div className="App">
      <NavbarApp />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="*" element={<NotFound/>} />
        <Route path="/detailbea" element={<Detailbea />} />
        <Route path="/listbea" element={<Listbea />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/faq" element={<Faq />} />
        <Route path="/contact" element={<Contct />} />
        <Route path="/list-beasiswa" element={<Listbea />} />
        <Route path="/editprofile" element={<Editprof />}/>
        
        <Route
          path="/beasiswa/create"
          element={
            <RequireAuth>
              <CreateBeasiswa />
            </RequireAuth>
          }
        />
        <Route
          path="/profile"
          element={
            <RequireAuth>
              <Profile />
            </RequireAuth>
          }
        />

        <Route
          path="/admin"
          element={
            <RequireAuth>
              <Admin>
                <User />
              </Admin>
            </RequireAuth>
          }
        />
        <Route
          path="/admin/user"
          element={
            <RequireAuth>
              <Admin>
                <User />
              </Admin>
            </RequireAuth>
          }
        />
        <Route
          path="/admin/comment"
          element={
            <RequireAuth>
              <Admin>
                <Comment />
              </Admin>
            </RequireAuth>
          }
        />
        <Route
          path="/admin/scholarship"
          element={
            <RequireAuth>
              <Admin>
                <Scholarship />
              </Admin>
            </RequireAuth>
          }
        />
        <Route
          path="/admin/category"
          element={
            <RequireAuth>
              <Admin>
                <Category />
              </Admin>
            </RequireAuth>
          }
        />
      </Routes>
      <Footer />
    </div>
  );
}

export default App;