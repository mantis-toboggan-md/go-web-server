import React from 'react';
import './App.css';
import { Login } from './Pages/Login/Login';
import { HashRouter as Router, Route, Link } from "react-router-dom";
import { Register } from './Pages/Register/Register';
import { Home } from './Pages/Home/Home'
import { privateRoute } from './Components/PrivateRoute'

function App() {
  return (
    <Router basename="/">
        <Route exact path="/" component={privateRoute(Home)} />
        <Route path="/login" component={Login} />
        <Route path="/register" component={Register} />
      </Router>
  );
}



export default App;
