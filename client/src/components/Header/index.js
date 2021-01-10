import React from 'react'
import { Image } from "@chakra-ui/react"
import Logo from '../../assets/logo.png'
import { Link } from "react-router-dom"
import { Tabs, TabList, Tab } from "@chakra-ui/react"
import { useLocation } from 'react-router-dom'

function Header() {
  const route = useLocation().pathname
  return (
    <header>
        <Link to="/">
            <Image id="Logo" src={Logo} alt="Logo"/>
        </Link>
        <Tabs>
            <TabList id="Nav">
                <Link to="/">
                    <Tab isSelected={route === "/" ? true : false}>
                        Home
                    </Tab>
                </Link>
                <Link to="/mybusiness">
                    <Tab isSelected={route === "/mybusiness" ? true : false}>
                        My Business
                    </Tab>
                </Link>
                <Link to="/auctions">
                    <Tab isSelected={route === "/auctions" ? true : false}>
                        Auctions
                    </Tab>                    
                </Link>
                <Link to="/wallet">
                    <Tab isSelected={route === "/wallet" ? true : false}>
                        Wallet
                    </Tab>
                </Link>
            </TabList>
        </Tabs>
        <div>
            Credits
            Name
            Avatar
        </div>
    </header>
  )
}

export default Header
