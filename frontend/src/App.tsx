import { Route, Routes } from "react-router-dom";
import {
  Home,
  Login,
  Ticket,
  Payment,
  Employee,
  User
} from "./pages";

function App() {
  return (
    <>
      <Routes>
        <Route path="/login" element={<Login role="employee"/>}/>
        <Route path="/" element={<Home/>}/>
        <Route path="/login/user" element={<User/>}/>
        <Route path="/ticket" element={<Ticket/>}/>
        <Route path="/payment" element={<Payment/>}/>
        <Route path="/employee" element={<Employee/>}/>
      </Routes>
    </>
  );
}

export default App;
