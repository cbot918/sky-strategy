import './App.css'

// import Map from './components/map/Map'
import Map from './components/map/Map'
import { BrowserRouter as Router, Route, Link, Routes } from 'react-router-dom';
import Tavern from './components/tavern/Tavern';
import General from './components/general/General';
import Bag from './components/bag/Bag';
import City from './components/city/City';
const Navbar = () =>{

  // const linkStyle = {
  //   textDecoration: 'none', // Remove underline
  //   color: 'black', // Set the link color to black
  //   cursor: 'pointer',

  // };

  return (
    <div className="navbar">
      <Link to="/" className="btn" >酒</Link>
      <Link to="/general" className="btn" >將</Link>
      <Link to="/map" className="btn" >野</Link>
      <Link to="/city" className="btn" >城</Link>
      <Link to="/bag" className="btn" >包</Link>
    </div>
  )
}

const Routing = () => {
  return (
    <Routes>
      <Route path="/" element={<Tavern/>} />
      <Route path="/general" element={<General />} />
      <Route path="/map" element={<Map />} />
      <Route path="/city" element={<City />} />
      <Route path="/bag" element={<Bag />} />
      {/* Define routes for other paths as needed */}
    </Routes>
  )
}

function App() {
  return (
    <Router>
      <div className="app">
        <div className="title">
          天 策
        </div>
        <div className="main"> <Routing/> </div>
        <div className="bottom"><Navbar/></div>
        {/* <Grid /> */}
      </div>   
    </Router>
  )
}

export default App
