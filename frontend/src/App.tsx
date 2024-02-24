import { Route, Routes } from "react-router-dom";
import Employee from "./pages/employee";
import Home from "./pages/home";
import Ticket from "./pages/ticket";
import Payment from "./pages/payment";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/home" element={<Home/>}/>
        <Route path="/ticket" element={<Ticket/>}/>
        <Route path="/payment" element={<Payment/>}/>
        <Route path="/employee" element={<Employee/>}/>
      </Routes>
    </>
  );
}

export default App;
