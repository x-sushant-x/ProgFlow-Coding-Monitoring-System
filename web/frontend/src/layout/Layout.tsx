import Header from '../components/Header/Header'
import SideBar from '../components/SideBar/Sidebar'
import './Layout.css'
import { Outlet } from 'react-router-dom'
import { useAuth0 } from '@auth0/auth0-react'


import NotLoggedIn from '../components/NotLoggedIn/NotLoggedIn'
import { useState, useEffect } from 'react'

export default function Layout() {
    const { isAuthenticated } = useAuth0()

    const [showMobileWarning, setShowMobileWarning] = useState(false)

    useEffect(() => {
        if (window.innerWidth <= 800)
            setShowMobileWarning(true)
    }, [])

    return (
        <>
            {
                showMobileWarning ? <p>Using Mobile</p> :
                    isAuthenticated ? <div>
                        <div className='body'>
                            <div className='header'>
                                <Header />
                            </div>
                            <div className='sidebar'>
                                <SideBar />
                            </div>
                            <div className='main'>
                                <Outlet />
                            </div>
                        </div>
                    </div> : <NotLoggedIn />
            }
        </>
    )
}