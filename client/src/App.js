import React from 'react'
import './App.css'
import Header from './components/Header'
import Home from './components/Home'
import MyBusiness from './components/MyBusiness'
import Auctions from './components/Auctions'
import Wallet from './components/Wallet'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

function App() {
  return (
    <Router>
      <Header/>
      <Switch>
        <Route exact path="/">
          <Home/>
        </Route>
        <Route exact path="/mybusiness">
          <MyBusiness/>
        </Route>
        <Route exact path="/auctions">
          <Auctions/>
        </Route>
        <Route exact path="/wallet">
          <Wallet/>
        </Route>
      </Switch>
    </Router>
  )
}

export default App
