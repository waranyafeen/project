import { Route, Routes } from "react-router-dom";
import Employee from "./pages/employee";
import Home from "./pages/home";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/home" element={<Home/>}/>
      </Routes>
    </>
  );
}

export default App;
