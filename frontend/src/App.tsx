import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home  from "./page/home";
import Login  from "./page/login";
import User  from "./page/user";
import Employee  from "./page/employee";
import Ticket  from "./page/ticket";
import Payment from "./page/payment";
import Navbar from "./layout/narBar/navbar";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login role="user"/>}/>
        <Route path="/login" element={<Login role="employee"/>}/>

        <Route path="/" element={<Home/>}/>
        <Route path="/user" element={<User/>}/>
        <Route path="/ticket" element={<Ticket/>}/>
        <Route path="/payment" element={<Payment/>}/>
        <Route path="/employee" element={<Employee/>}/>
        <Route path="/navber" element={<Navbar/>}/>
      </Routes>
    </Router>
  );
}

export default App;
