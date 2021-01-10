import React from 'react'
import { Image } from "@chakra-ui/react"
import Logo from '../../assets/logo.png'
import { Link } from "react-router-dom"
import { Tabs, TabList, Tab } from "@chakra-ui/react"

function Header() {
  return (
    <header>
        <Link to="/">
            <Image id="Logo" src={Logo} alt="Logo"/>
        </Link>
        <Tabs>
            <TabList id="Nav">
                <Link to="/">
                    <Tab>
                        Home
                    </Tab>
                </Link>
                <Link to="/mybusiness">
                    <Tab>
                        My Business
                    </Tab>
                </Link>
                <Link to="/auctions">
                    <Tab>
                        Auctions
                    </Tab>                    
                </Link>
                <Link to="/wallet">
                    <Tab>
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
