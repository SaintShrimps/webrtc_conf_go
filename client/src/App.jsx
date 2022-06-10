import React, { useState } from 'react';
import { BrowserRouter, Routes, Route  } from "react-router-dom";

//In react-router-dom v6, "Switch" is replaced by routes "Routes"

import CreateRoom from './components/CreateRoom';
import Login from './components/Login';
import Room from './components/Room';


function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Login />} />
          <Route path="/user" element={<CreateRoom />} />
          <Route path='room/' element={<Room />} >
            <Route path=':roomID' element={<Room />} />
          </Route>
        
        </Routes>
      </BrowserRouter> 
    </div>
  );
}

export default App;

//Route exact path="/"